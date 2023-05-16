package clients

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ViaCepClient struct {
	httpClient *http.Client
}

func NewViaCepClient() *ViaCepClient {
	return &ViaCepClient{
		httpClient: http.DefaultClient,
	}
}

type ViaCepOutPut struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (c ViaCepClient) FindByZipCode(zipCode string) (*ViaCepOutPut, error) {
	req, err := http.NewRequest("GET", "https://viacep.com.br/ws/"+zipCode+"/json/", nil)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var viaCepOutput ViaCepOutPut
	err = json.Unmarshal(body, &viaCepOutput)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &viaCepOutput, nil
}
