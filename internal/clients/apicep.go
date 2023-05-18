package clients

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
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
	start := time.Now()
	req, err := http.NewRequest("GET", "https://cdn.apicep.com/file/apicep/"+zipCode+".json", nil)
	log.Printf("ApiCepClient - FindByZipCode - params: %s", zipCode)

	if err != nil {
		panic(err)
	}
	res, err := a.httpClient.Do(req)
	if err != nil {
		log.Printf("ApiCepClient - Request error: %s", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("ApiCepClient - Read response error: %s", err)
		return nil, err
	}

	var apiCepOutPut ApiCepOutPut
	err = json.Unmarshal(body, &apiCepOutPut)
	if err != nil {
		log.Printf("ApiCepClient - Unmarshal error: %s", err)
		return nil, err
	}

	if apiCepOutPut.Status != http.StatusOK {
		log.Printf("ApiCepClient - Response error - Status: %d", apiCepOutPut.Status)
		return nil, errors.New(apiCepOutPut.StatusText)
	}

	elapsed := time.Since(start).Milliseconds()
	log.Printf("ApiCepClient - FindByZipCode - found: %v, elapsed: %vms", apiCepOutPut, elapsed)
	return &apiCepOutPut, nil
}
