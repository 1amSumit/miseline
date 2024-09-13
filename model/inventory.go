package model

import "time"

type Inventory struct {
	Id        int64
	Name      string
	Type      string
	TotalItem int64
	OutletId  int64
	CreateAt  *time.Time
}
