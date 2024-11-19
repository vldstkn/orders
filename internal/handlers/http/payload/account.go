package payload

import "orders/internal/domain"

// Register

type AccountRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type AccountRegisterResponse struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

// Login

type AccountLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountLoginResponse struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

// GetNewTokens

type AccountGetNewTokensResponse struct {
	AccessToken string `json:"access_token"`
}

// GetPublicProfile

type AccountGetPublicProfileResponse struct {
	Id                    int               `json:"id"`
	Email                 string            `json:"email"`
	Name                  string            `json:"name"`
	Role                  string            `json:"role"`
	Rating                float32           `json:"rating"`
	NumberCompletedOrders int               `json:"number_completed_orders"`
	Offerings             []domain.Offering `json:"offerings"`
}

// UpdateById

type AccountUpdateByIdResponse struct {
	IsSuccess bool `json:"is_success"`
}

// ChangeRoleById

type AccountChangeRoleByIdRequest struct {
	Role string `json:"role"`
}

type AccountChangeRoleByIdResponse struct {
	IsSuccess bool `json:"is_success"`
}
