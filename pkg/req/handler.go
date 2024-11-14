package req

import (
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		return nil, err
	}
	err = IsValid[T](body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
