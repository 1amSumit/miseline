package model

import "time"

type Staff struct {
	Id        int64
	Name      string
	Image     string
	Role      string
	Address   string
	Email     string
	Phone     int64
	City      string
	State     string
	Zip       int64
	Country   string
	OutletId  int64
	CreatedAt *time.Time
}
