package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/mmunareto/zipcode-api/internal/services"
	"net/http"
)

type ZipCodeHandler struct {
	ZipCodeService services.ZipCodeServiceInterface
}

func NewZipCodeHandler(service services.ZipCodeServiceInterface) *ZipCodeHandler {
	return &ZipCodeHandler{
		ZipCodeService: service,
	}
}

func (z *ZipCodeHandler) GetZipCodeDetails(w http.ResponseWriter, r *http.Request) {
	zipCode := chi.URLParam(r, "zipCode")
	if zipCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := z.ZipCodeService.FindByZipCode(zipCode)
	if result.Error != "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
