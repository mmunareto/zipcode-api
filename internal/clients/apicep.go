package clients

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ApiCepClient struct {
	httpClient *http.Client
}

func NewApiCepClient() *ApiCepClient {
	return &ApiCepClient{
		httpClient: http.DefaultClient,
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
	if err != nil {
		panic(err)
	}
	res, err := a.httpClient.Do(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var apiCepOutPut ApiCepOutPut
	err = json.Unmarshal(body, &apiCepOutPut)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &apiCepOutPut, nil
}
