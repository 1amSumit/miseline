package model

import (
	"time"

	"sumit.com/mise-link/db"
)

type Product struct {
	Id          int64
	Name        string  `form:"name"`
	Image       string  `form:"image"`
	Type        string  `form:"type"`
	Price       float64 `form:"price"`
	Description string  `form:"description"`
	Quantity    int     `form:"quantity"`
	InventoryId int64
	CreatedAt   *time.Time
}

func (p *Product) Save() error {
	now := time.Now()
	p.CreatedAt = &now
	query := "INSERT INTO products (name, image, type, price,description quantity,  inventory_id, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.Name, p.Image, p.Type, p.Price, p.Description, p.Quantity, p.InventoryId, p.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}
