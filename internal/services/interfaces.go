package services

import (
	"github.com/mmunareto/zipcode-api/internal/dto"
)

type ZipCodeServiceInterface interface {
	FindByZipCode(zipCode string) (*dto.ZipCodeOutput, error)
}
