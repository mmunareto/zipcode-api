package clients

import (
	"encoding/json"
	"io"
	"net/http"
)

type ApiCepClient struct {
	httpClient *http.Client
}

func NewApiCepClient(client *http.Client) *ApiCepClient {
	return &ApiCepClient{
		httpClient: client,
	}
}

type ApiCepOutPut struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func (a ApiCepClient) FindByZipCode(zipCode string) (*ApiCepOutPut, error) {
	req, err := http.NewRequest("GET", "https://cdn.apicep.com/file/apicep/"+zipCode+".json", nil)

	res, err := a.httpClient.Do(req)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	var apiCepOutPut ApiCepOutPut
	err = json.Unmarshal(body, &apiCepOutPut)

	return &apiCepOutPut, err
}
