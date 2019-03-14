/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// FEDWireMessage
type FedWireMessage struct {
	MessageDisposition MessageDisposition `json:"messageDisposition,omitempty"`
	ReceiptTimeStamp ReceiptTimeStamp `json:"receiptTimeStamp,omitempty"`
	OutputMessageAccountabilityData OutputMessageAccountabilityData `json:"outputMessageAccountabilityData,omitempty"`
	Error FedWireError `json:"error,omitempty"`
	SenderSupplied SenderSupplied `json:"senderSupplied"`
	TypeSubType TypeSubType `json:"typeSubType"`
	InputMessageAccountabilityData InputMessageAccountabilityData `json:"inputMessageAccountabilityData"`
	Amount Amount `json:"amount"`
	SenderDepositoryInstitution SenderDepositoryInstitution `json:"senderDepositoryInstitution"`
	ReceiverDepositoryInstitution ReceiverDepositoryInstitution `json:"receiverDepositoryInstitution"`
	BusinessFunctionCode BusinessFunctionCode `json:"businessFunctionCode"`
	// SenderReference
	SenderReference string `json:"senderReference,omitempty"`
	// PreviousMessageIdentifier
	PreviousMessageIdentifier string `json:"previousMessageIdentifier,omitempty"`
	LocalInstrument LocalInstrument `json:"localInstrument,omitempty"`
	PaymentNotification PaymentNotification `json:"paymentNotification,omitempty"`
	Charges Charges `json:"charges,omitempty"`
	InstructedAmount InstructedAmount `json:"instructedAmount,omitempty"`
	// ExchangeRate  Must contain at least one numeric character and only one decimal comma marker (e.g., an exchange rate of 1.2345 should be entered as 1,2345). 
	ExchangeRate string `json:"exchangeRate,omitempty"`
	BeneficiaryIntermediaryFI FinancialInstitution `json:"beneficiaryIntermediaryFI,omitempty"`
	BeneficiaryFI FinancialInstitution `json:"beneficiaryFI,omitempty"`
	Beneficiary Personal `json:"beneficiary,omitempty"`
	// ReferenceForBeneficiary {4320}
	ReferenceForBeneficiary string `json:"referenceForBeneficiary,omitempty"`
	AccountDebitedDrawdown AccountDebitedDrawdown `json:"accountDebitedDrawdown,omitempty"`
	Originator Personal `json:"originator,omitempty"`
	// OriginatorOptionF {5010}
	OriginatorOptionF map[string]interface{} `json:"originatorOptionF,omitempty"`
	OriginatorFI FinancialInstitution `json:"originatorFI,omitempty"`
	InstructingFI FinancialInstitution `json:"instructingFI,omitempty"`
	AccountCreditedDrawdown AccountCreditedDrawdown `json:"accountCreditedDrawdown,omitempty"`
	OriginatorToBeneficiary OriginatorToBeneficiary `json:"originatorToBeneficiary,omitempty"`
	ReceiverFI FiToFi `json:"receiverFI,omitempty"`
	DrawdownDebitAccountAdvice Advice `json:"drawdownDebitAccountAdvice,omitempty"`
	IntermediaryFI FiToFi `json:"intermediaryFI,omitempty"`
	IntermediaryFinacialInstitutionAdvice Advice `json:"intermediaryFinacialInstitutionAdvice,omitempty"`
	OriginatorBeneficiaryFinancialInstitution FiToFi `json:"originatorBeneficiaryFinancialInstitution,omitempty"`
	OriginatorBeneficiaryFinancialInstitutionAdvice Advice `json:"originatorBeneficiaryFinancialInstitutionAdvice,omitempty"`
	OriginatorBeneficiary FiToFi `json:"originatorBeneficiary,omitempty"`
	BeneficiaryAdvice Advice `json:"beneficiaryAdvice,omitempty"`
	PaymentMethodToBeneficiary PaymentMethodToBeneficiary `json:"paymentMethodToBeneficiary,omitempty"`
	AdditionalFIToFI AdditionalFiToFi `json:"additionalFIToFI,omitempty"`
	CurrencyInstructedAmount CoverPayment `json:"currencyInstructedAmount,omitempty"`
	OrderingCustomer CoverPayment `json:"orderingCustomer,omitempty"`
	OrderingInstitution CoverPayment `json:"orderingInstitution,omitempty"`
	IntermediaryInstitution CoverPayment `json:"intermediaryInstitution,omitempty"`
	InstitutionAccount CoverPayment `json:"institutionAccount,omitempty"`
	BeneficiaryCustomer CoverPayment `json:"beneficiaryCustomer,omitempty"`
	Remittance CoverPayment `json:"remittance,omitempty"`
	SenderToReceiver CoverPayment `json:"senderToReceiver,omitempty"`
	UnstructuredAddenda UnstructuredAddenda `json:"unstructuredAddenda,omitempty"`
	RelatedRemittance RelatedRemittance `json:"relatedRemittance,omitempty"`
	RemittanceOriginator RemittanceOriginator `json:"remittanceOriginator,omitempty"`
	RemittanceBeneficiary RemittanceBeneficiary `json:"remittanceBeneficiary,omitempty"`
	PrimaryRemittanceDocument PrimaryRemittanceDocument `json:"primaryRemittanceDocument,omitempty"`
	ActualAmountPaid RemittanceAmount `json:"actualAmountPaid,omitempty"`
	GrossAmountRemittanceDocument RemittanceAmount `json:"grossAmountRemittanceDocument,omitempty"`
	AmountNegotiatedDiscount RemittanceAmount `json:"amountNegotiatedDiscount,omitempty"`
	Adjustment AdjustmentEnum `json:"adjustment,omitempty"`
	DateRemittanceDocument DateRemittanceDocument `json:"dateRemittanceDocument,omitempty"`
	SecondaryRemittanceDocument SecondaryRemittanceDocumentEnum `json:"secondaryRemittanceDocument,omitempty"`
	RemittanceFreeText RemittanceFreeText `json:"remittanceFreeText,omitempty"`
	ServiceMessage ServiceMessage `json:"serviceMessage,omitempty"`
}
