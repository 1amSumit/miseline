package model

import (
	"time"

	"sumit.com/mise-link/db"
)

type Outlet struct {
	Id       int64
	Name     string `form:"name"`
	Address  string `form:"address"`
	Image    string `form:"image"`
	City     string `form:"city"`
	State    string `form:"state"`
	Zip      int64  `form:"zip"`
	Country  string `form:"country"`
	CreateAt *time.Time
	UserId   int64
}

func (o *Outlet) Save() error {
	now := time.Now()
	o.CreateAt = &now

	query := `INSERT INTO outlets (name, address, image, city, state, zip, country, created_at, user_id) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(o.Name, o.Address, o.Image, o.City, o.State, o.Zip, o.Country, o.CreateAt, o.UserId)

	return err
}
