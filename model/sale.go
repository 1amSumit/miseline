package model

import "time"

type Sale struct {
	Id        int64
	ProductId int64
	OutletId  int64
	CreatedAt time.Time
}
