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
	id := chi.URLParam(r, "zipCode")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	zipDetails, err := z.ZipCodeService.FindByZipCode(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zipDetails)
}
