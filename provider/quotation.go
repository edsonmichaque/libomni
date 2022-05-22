package provider

type Quotation struct {
	Amount   string       `json:"amount"`
	Currency string       `json:"currency"`
	Count    int          `json:"count"`
	Size     MessageStats `json:"size"`
}

type MessageStats struct {
	SizeBytes      int `json:"size_bytes"`
	SizeCharacters int `json:"size_characters"`
}
