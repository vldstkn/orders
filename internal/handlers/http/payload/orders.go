package payload

type PublicOrder struct {
	Status       string `json:"status"`
	CustomerId   int    `json:"customer_id"`
	ContractorId int    `json:"contractor_id"`
	OfferingId   int    `json:"offering_id"`
}

// Create

type OrdersCreateRequest struct {
	CustomerId int `json:"customer_id"`
	OfferingId int `json:"offering_id"`
}

type OrdersCreateResponse struct {
	PublicOrder
}

// UpdateById

type OrdersUpdateByIdResponse struct {
	IsSuccess bool `json:"is_success"`
}

// GetById

type OrdersGetByIdResponse struct {
	PublicOrder
}

// GetByUserId

type OrdersGetByUserIdResponse struct {
	PublicOrder
}
