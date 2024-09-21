package model

import (
	"time"

	"sumit.com/mise-link/db"
)

type Staff struct {
	Id        int64
	Name      string `form:"name"`
	Image     string `form:"image"`
	Role      string `form:"role"`
	Address   string `form:"address"`
	Email     string `form:"email"`
	Phone     int64  `form:"phone"`
	City      string `form:"city"`
	State     string `form:"state"`
	Zip       int64  `form:"zip"`
	Country   string `form:"country"`
	OutletId  int64
	CreatedAt *time.Time
}

func (s *Staff) Save() error {
	now := time.Now()
	s.CreatedAt = &now

	query := `INSERT INTO staff (name, image, role, address, email,phone , city , state, zip, country, outlet_id, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(s.Name, s.Image, s.Role, s.Address, s.Email, s.Phone, s.City, s.State, s.Zip, s.Country, s.OutletId, s.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
