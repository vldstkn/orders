package payload

type AccountRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type AccountRegisterResponse struct {
	Email string `json:"email"`
}
