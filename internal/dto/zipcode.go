package dto

type Result struct {
	ZipCodeDetails *ZipCodeDetails `json:"zipCodeDetails,omitempty"`
	Provider       string          `json:"provider,omitempty"`
	Error          string          `json:"error,omitempty"`
}

type ZipCodeDetails struct {
	ZipCode  string `json:"zipCode"`
	Address  string `json:"address"`
	District string `json:"district"`
	State    string `json:"state"`
	City     string `json:"city"`
}
