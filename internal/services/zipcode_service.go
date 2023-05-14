package services

import (
	"github.com/mmunareto/zipcode-api/internal/clients"
	"github.com/mmunareto/zipcode-api/internal/dto"
	"net/http"
)

type ZipCodeService struct {
	apiCepClient *clients.ApiCepClient
	viaCepClient *clients.ViaCepClient
}

func NewZipCodeService(httpClient *http.Client) *ZipCodeService {
	return &ZipCodeService{
		apiCepClient: clients.NewApiCepClient(httpClient),
		viaCepClient: clients.NewViaCepClient(httpClient),
	}
}

func (z *ZipCodeService) FindByZipCode(zipCode string) (*dto.ZipCodeOutput, error) {
	channelApiCep := make(chan *dto.ZipCodeOutput)
	channelViaCep := make(chan *dto.ZipCodeOutput)

	go func() {
		apiCepResponse, _ := z.apiCepClient.FindByZipCode(zipCode)
		apiCepOutPut := &dto.ZipCodeOutput{
			Localidade: apiCepResponse.City,
		}
		channelApiCep <- apiCepOutPut
	}()

	go func() {
		viaCepResponse, _ := z.viaCepClient.FindByZipCode(zipCode)
		viaCepOutPut := &dto.ZipCodeOutput{
			Cep: viaCepResponse.Cep,
		}
		channelViaCep <- viaCepOutPut
	}()

	select {
	case msg := <-channelApiCep:
		return msg, nil
	case msg := <-channelViaCep:
		return msg, nil
	}
}
