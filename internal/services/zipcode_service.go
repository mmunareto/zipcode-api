package services

import (
	"errors"
	"github.com/mmunareto/zipcode-api/internal/clients"
	"github.com/mmunareto/zipcode-api/internal/dto"
	"time"
)

type ZipCodeService struct {
	apiCepClient *clients.ApiCepClient
	viaCepClient *clients.ViaCepClient
}

func NewZipCodeService() *ZipCodeService {
	return &ZipCodeService{
		apiCepClient: clients.NewApiCepClient(),
		viaCepClient: clients.NewViaCepClient(),
	}
}

func (z *ZipCodeService) FindByZipCode(zipCode string) *dto.Result {
	channelApiCep := make(chan *dto.Result)
	channelViaCep := make(chan *dto.Result)

	go func() {
		provider := "ApiCep"
		apiCepResponse, err := z.apiCepClient.FindByZipCode(zipCode)
		if err != nil {
			result := &dto.Result{Provider: provider, Error: err}
			channelApiCep <- result
			return
		}
		apiCepOutPut := &dto.ZipCodeDetails{
			ZipCode:  apiCepResponse.Code,
			Address:  apiCepResponse.Address,
			District: apiCepResponse.District,
			City:     apiCepResponse.City,
			State:    apiCepResponse.State,
		}

		result := &dto.Result{
			Provider:       provider,
			ZipCodeDetails: apiCepOutPut,
		}
		channelApiCep <- result
	}()

	go func() {
		provider := "ViaCep"
		viaCepResponse, err := z.viaCepClient.FindByZipCode(zipCode)
		if err != nil {
			result := &dto.Result{Provider: provider, Error: err}
			channelViaCep <- result
			return
		}

		viaCepOutPut := &dto.ZipCodeDetails{
			ZipCode:  viaCepResponse.Cep,
			Address:  viaCepResponse.Logradouro,
			District: viaCepResponse.Bairro,
			City:     viaCepResponse.Localidade,
			State:    viaCepResponse.Uf,
		}

		result := &dto.Result{
			Provider:       provider,
			ZipCodeDetails: viaCepOutPut,
		}
		channelViaCep <- result
	}()

	select {
	case msg := <-channelApiCep:
		return msg
	case msg := <-channelViaCep:
		return msg
	case <-time.After(1 * time.Second):
		return &dto.Result{Error: errors.New("request timeout")}
	}
}
