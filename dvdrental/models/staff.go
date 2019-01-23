package models

type Staff struct {
	StaffID   int     `json:"staff_id"`
	Email     string  `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	AddressID int     `json:"address_id"`
	StoreID   int     `json:"store_id"`
}
