package http

import (
	"fmt"
	"net/http"
)

type OrdersHttpHandler struct {
}

func NewOrdersHttpHandler(router *http.ServeMux) {
	handler := OrdersHttpHandler{}
	router.HandleFunc("GET /orders/{id}", handler.GetById())
	router.HandleFunc("POST /orders", handler.Create())
	router.HandleFunc("PUT /orders/{id}", handler.UpdateById())
}

func (handler *OrdersHttpHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetById")
	}
}

func (handler *OrdersHttpHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create")
	}
}

func (handler *OrdersHttpHandler) UpdateById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("UpdateById")
	}
}
