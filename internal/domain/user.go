package domain

type User struct {
	Id                    string `json:"id" db:"id"`
	Email                 string `json:"email" db:"email"`
	Password              string `json:"password" db:"password"`
	Name                  string `json:"name" db:"name"`
	Role                  string `json:"role" db:"role"`
	NumberCompletedOrders int    `json:"number_completed_orders" db:"number_completed_orders"`
}
