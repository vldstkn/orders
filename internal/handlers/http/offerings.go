package handlers

import (
	"net/http"
)

type OfferingsHttpHandler struct {
}

func NewOfferingsHttpHandler(router *http.ServeMux) {
	handler := OfferingsHttpHandler{}
	router.HandleFunc("POST /offerings", handler.Create())
	router.HandleFunc("DELETE /offerings/{id}", handler.DeleteById())
	router.HandleFunc("PUT /offerings/{id}", handler.UpdateById())

	router.HandleFunc("GET /offerings/{id}", handler.GetById())
	router.HandleFunc("GET /offerings/user/{id}", handler.GetByUserId())
	router.HandleFunc("GET /offerings", handler.GetByTitle())
}

func (handler *OfferingsHttpHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (handler *OfferingsHttpHandler) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (handler *OfferingsHttpHandler) UpdateById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *OfferingsHttpHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *OfferingsHttpHandler) GetByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *OfferingsHttpHandler) GetByTitle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
