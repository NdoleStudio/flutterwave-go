package stubs

// BillsCreateDStvPaymentResponse is a dummy JSOn response for checking a dstv user
func BillsCreateDStvPaymentResponse() string {
	return `
	{
		"status": "success",
		"message": "Bill payment successful",
		"data": {
			"phone_number": "+23490803840303",
			"amount": 500,
			"network": "9MOBILE",
			"flw_ref": "CF-FLYAPI-20200311081921359990",
			"tx_ref": "BPUSSD1583957963415840"
		}
	}
`
}

// BillsValidateDstvResponse is a dummy response for validating a DStv payment
func BillsValidateDstvResponse() string {
	return `
	{
		"status": "success",
		"message": "Item validated successfully",
		"data": {
			"response_code": "00",
			"address": null,
			"response_message": "Successful",
			"name": "MTN",
			"biller_code": "BIL099",
			"customer": "08038291822",
			"product_code": "AT099",
			"email": null,
			"fee": 100,
			"maximum": 0,
			"minimum": 0
		}
	}
`
}
