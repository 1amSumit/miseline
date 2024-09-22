package model

import (
	"time"

	"sumit.com/mise-link/db"
)

type Inventory struct {
	Id        int64
	Name      string `form:"name"`
	Type      string `form:"type"`
	TotalItem int64  `form:"total_item"`
	OutletId  int64
	CreatedAt *time.Time
}

func (i *Inventory) Save() error {

	now := time.Now()
	i.CreatedAt = &now

	query := `INSERT INTO inventory (name, type, total_item, outlet_id, created_at) VALUES ($1, $2, $3, $4, $5)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(i.Name, i.Type, i.TotalItem, i.OutletId, i.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
