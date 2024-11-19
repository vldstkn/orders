package payload

type PublicOffering struct {
	Id             string `json:"id"`
	Status         string `json:"status"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Price          int    `json:"price"`
	ContractorId   int    `json:"contractor_id"`
	ContractorName int    `json:"contractor_name"`
}

// Create

type OfferingsCreateRequest struct {
	Title       string `json:"string"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type OfferingsCreateResponse struct {
	Id int `json:"id"`
}

// DeleteById

type OfferingsDeleteByIdResponse struct {
	Id        int  `json:"id"`
	IsSuccess bool `json:"is_success"`
}

// UpdateById

type OfferingsUpdateByIdRequest struct {
	Title       string `json:"string,omitempty"`
	Description string `json:"description,omitempty"`
	Price       int    `json:"price,omitempty"`
}

type OfferingsUpdateByIdResponse struct {
	IsSuccess bool `json:"is_success"`
}

// GetById

type OfferingsGetByIdResponse struct {
	PublicOffering
}

// GetByTitle

type OfferingsGetByTitleResponse struct {
	Offerings []PublicOffering `json:"offerings"`
	Count     int              `json:"count"`
}

// GetByUserId

type OfferingsGetByUserIdResponse struct {
	Offerings []PublicOffering `json:"offerings"`
	Count     int              `json:"count"`
}
