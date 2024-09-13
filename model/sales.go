package model

import "time"

type Sales struct {
	Id         int64
	OutletId   int64
	ProductId  int64
	Quantity   int
	TotalPrice float64
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
