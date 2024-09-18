package entity

import "time"

type Product struct {
	ClientId int       `json:"client_id"`
	Item     string    `json:"item"`
	Quantity int       `json:"quantity"`
	Kind     string    `json:"kind"`
	Date     time.Time `json:"date"`
}