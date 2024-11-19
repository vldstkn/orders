package domain

type User struct {
	Id                    string  `json:"id" db:"id"`
	CreatedAt             string  `json:"created_at" db:"created_at"`
	UpdatedAt             string  `json:"updated_at" db:"updated_at"`
	Email                 string  `json:"email" db:"email"`
	Password              string  `json:"password" db:"password"`
	Name                  string  `json:"name" db:"name"`
	Role                  string  `json:"role" db:"role"`
	Rating                float32 `json:"rating" db:"rating"`
	NumberCompletedOrders int     `json:"number_completed_orders" db:"number_completed_orders"`
}

type Offering struct {
	Id           int    `json:"id" db:"id"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
	Status       string `json:"status" db:"status"`
	Title        string `json:"title" db:"title"`
	Description  string `json:"description" db:"description"`
	Price        int    `json:"price" db:"price"`
	ContractorId int    `json:"contractor_id" db:"contractor_id"`
}

type Order struct {
	Id           string `json:"id" db:"id"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
	Status       string `json:"status" db:"status"`
	CustomerId   int    `json:"customer_id" db:"customer_id"`
	ContractorId int    `json:"contractor_id" db:"contractor_id"`
	OfferingId   int    `json:"offering_id" db:"offering_id"`
}
