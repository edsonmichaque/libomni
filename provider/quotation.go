package provider

type Quotation struct {
	Amount   string      `json:"amount"`
	Currency string      `json:"currency"`
	Count    int         `json:"count"`
	Size     MessageSize `json:"size"`
}
