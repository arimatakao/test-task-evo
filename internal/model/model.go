package model

type Info struct {
	TransactionId      string `json:"TransactionId"`
	RequestId          string `json:"RequestId "`
	TerminalId         string `json:"TerminalId"`
	PartnerObjectId    string `json:"PartnerObjectId"`
	AmountTotal        string `json:"AmountTotal"`
	AmountOriginal     string `json:"AmountOriginal"`
	CommissionPS       string `json:"CommissionPS"`
	CommissionClient   string `json:"CommissionClient"`
	CommissionProvider string `json:"CommissionProvider"`
	DateInput          string `json:"DateInput"`
	DatePost           string `json:"DatePost"`
	Status             string `json:"Status"`
	PaymentType        string `json:"PaymentType"`
	PaymentNumber      string `json:"PaymentNumber"`
	ServiceId          string `json:"ServiceId"`
	Service            string `json:"Service"`
	PayeeId            string `json:"PayeeId"`
	PayeeName          string `json:"PayeeName"`
	PayeeBankMfo       string `json:"PayeeBankMfo"`
	PayeeBankAccount   string `json:"PayeeBankAccount"`
	PaymentNarrative   string `json:"PaymentNarrative"`
}
