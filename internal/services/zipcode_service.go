package services

import (
	"github.com/mmunareto/zipcode-api/internal/clients"
	"github.com/mmunareto/zipcode-api/internal/dto"
	"strings"
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

	zipCodeNormalized := normalizeZipCode(zipCode)

	go func() {
		provider := "ApiCep"
		apiCepResponse, err := z.apiCepClient.FindByZipCode(zipCodeNormalized)
		if err == nil {
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
		}
	}()

	go func() {
		provider := "ViaCep"
		viaCepResponse, err := z.viaCepClient.FindByZipCode(zipCodeNormalized)
		if err == nil {
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
		}
	}()

	select {
	case result := <-channelApiCep:
		return result
	case result := <-channelViaCep:
		return result
	case <-time.After(1 * time.Second):
		return &dto.Result{Error: "request timeout"}
	}
}

func normalizeZipCode(zipCode string) string {
	if !strings.Contains(zipCode, "-") {
		zipCode = zipCode[0:5] + "-" + zipCode[5:8]
	}
	return zipCode
}
