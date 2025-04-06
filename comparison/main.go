package main

import "time"

type getBillOutput struct {
	Bills []retrieveBillOutput
}

// RetrieveBillOutput the bill object of GetBillOutput
type retrieveBillOutput struct {
	AccountNumber string
	PdfURL        string
	Amount        int64
	Date          time.Time
	Period        time.Time
	DueDate       time.Time
}

func main() {
	newBills := getBillOutput{Bills: []retrieveBillOutput{
		{AccountNumber:"008952239609", PdfURL: "https://b2c.api.spdigital.sg/njord/v1/bills/download?sig=aW6ZQ-r8gV4Ntgm2kfSMKkkDIOzPztyEi6ojBcBaznPdSBCeu8dItaHn8mZJw4EcEihIHl0oUAYCz1lWmbB8n_LzRC52fMDuWkW_DviiMA0n4VYzXpBUzfX1p45CBzA6m-zaAkzy3zR7lC_YvWyu9Sow8deGS483y20MMaV5FhdC-07-JL2v6Qx046qqS-CXhi7bKvWupraZspyikPhlUg==",
			Amount:10028,
			Date: 2024-06-12 16:00:00 +0000 UTC,
			Period: "2024-05-31 16:00:00 +0000 UTC",
			DueDate: "2024-06-26 16:00:00 +0000 UTC",
		},
		{},
	}}
}
