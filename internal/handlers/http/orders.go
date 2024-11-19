package handlers

import (
	"net/http"
)

type OrdersHttpHandler struct {
}

func NewOrdersHttpHandler(router *http.ServeMux) {
	handler := OrdersHttpHandler{}
	router.HandleFunc("POST /orders", handler.Create())
	router.HandleFunc("PUT /orders/{id}", handler.UpdateById())

	router.HandleFunc("GET /orders/{id}", handler.GetById())
	router.HandleFunc("GET /orders/user/{id}", handler.GetByUserId())
}

func (handler *OrdersHttpHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *OrdersHttpHandler) GetByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *OrdersHttpHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *OrdersHttpHandler) UpdateById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
