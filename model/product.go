package model

import "time"

type Product struct {
	Id          int64
	Name        string
	Image       string
	Type        string
	Price       float64
	Quantity    int
	InventoryId int64
	CreatedAt   *time.Time
}
