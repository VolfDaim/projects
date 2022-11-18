package models

type Transaction struct {
	ID        int `json:"id" db:"id"`
	IDService int `json:"id_service" db:"id_service"`
	IDOrder   int `json:"id_order" db:"id_order"`
	Balance   int `json:"balance" db:"balance"`
}
