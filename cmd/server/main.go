// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/moov-io/base/admin"
	"github.com/moov-io/base/http/bind"
	"github.com/moov-io/wire"

	"github.com/go-kit/kit/log"
)

var (
	httpAddr  = flag.String("http.addr", bind.HTTP("wire"), "HTTP listen address")
	adminAddr = flag.String("admin.addr", bind.Admin("wire"), "Admin HTTP listen address")

	flagLogFormat = flag.String("log.format", "", "Format for log lines (Options: json, plain")
)

func main() {
	flag.Parse()

	var logger log.Logger
	if strings.ToLower(*flagLogFormat) == "json" {
		logger = log.NewJSONLogger(os.Stderr)
	} else {
		logger = log.NewLogfmtLogger(os.Stderr)
	}
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	logger.Log("startup", fmt.Sprintf("Starting wire server version %s", wire.Version))

	// Channel for errors
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Start Admin server (with Prometheus metrics)
	adminServer := admin.NewServer(*adminAddr)
	go func() {
		logger.Log("admin", fmt.Sprintf("listening on %s", adminServer.BindAddr()))
		if err := adminServer.Listen(); err != nil {
			err = fmt.Errorf("problem starting admin http: %v", err)
			logger.Log("admin", err)
			errs <- err
		}
	}()
	defer adminServer.Shutdown()

	var wireFileTTL time.Duration
	if v := os.Getenv("WIRE_FILE_TTL"); v != "" {
		dur, err := time.ParseDuration(v)
		if err == nil {
			wireFileTTL = dur
			logger.Log("main", fmt.Sprintf("Using %v as wire.File TTL", wireFileTTL))
		}
	}

	// Setup business HTTP routes
	repo := NewRepositoryInMemory(wireFileTTL, logger)
	svc := NewService(repo)

	handler := MakeHTTPHandler(svc, repo, logger)

	// Start business HTTP server
	readTimeout, _ := time.ParseDuration("30s")
	writTimeout, _ := time.ParseDuration("30s")
	idleTimeout, _ := time.ParseDuration("60s")

	serve := &http.Server{
		Addr:    *httpAddr,
		Handler: handler,
		TLSConfig: &tls.Config{
			InsecureSkipVerify:       false,
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
		},
		ReadTimeout:  readTimeout,
		WriteTimeout: writTimeout,
		IdleTimeout:  idleTimeout,
	}
	shutdownServer := func() {
		if err := serve.Shutdown(context.TODO()); err != nil {
			logger.Log("shutdown", err)
		}
	}

	// Start business logic HTTP server
	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- serve.ListenAndServe()
		// TODO(adam): support TLS
		// func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
	}()

	// Block/Wait for an error
	if err := <-errs; err != nil {
		shutdownServer()
		logger.Log("exit", err)
	}
}