package domain

// invoice exported private methods for testing
var ExportedInvoice_calculateFeeAmount = (*Invoice).calculateFeeAmount
var ExportedInvoice_calculateTaxAmount = (*Invoice).calculateTaxAmount
var ExportedInvoice_calculateTotalAmount = (*Invoice).calculateTotalAmount
