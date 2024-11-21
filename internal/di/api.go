package di

import (
	"net/http"
	"time"
)

type ApiService interface {
	AddCookie(w http.ResponseWriter, name, value string, expirationTime time.Time)
}
