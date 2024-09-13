package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "postgres://postgres:root@localhost/shopdb?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

}

func createTables() {
	createUserTable := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    )`

	createOutlerTable := `
	CREATE TABLE IF NOT EXISTS outlets (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            address TEXT NOT NULL,
			image TEXT NOT NULL,
			city TEXT NOT NULL,
			state TEXT NOT NULL,
            zip TEXT NOT NULL,
            country TEXT NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			user_id INT,
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
			)`

	createInventoryTable := `
	CREATE TABLE IF NOT EXISTS inventory(
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
        total_item INT NOT NULL,
        outlet_id INT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (outlet_id) REFERENCES outlets (id) ON DELETE CASCADE
	)`

	createProductTable := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		image TEXT NOT NULL,
		type TEXT NOT NULL,
		quantity INT NOT NULL,
        description TEXT NOT NULL,
        price DECIMAL NOT NULL,
        inventory_id INT,
		FOREIGN KEY (inventory_id) REFERENCES inventory (id) ON DELETE CASCADE,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	createSalesTable := `
	CREATE TABLE IF NOT EXISTS sales (
		id SERIAL PRIMARY KEY,
		outlet_id INT NOT NULL,
		product_id INT NOT NULL,
        quantity INT NOT NULL,
        total_price DECIMAL NOT NULL,
        FOREIGN KEY (outlet_id) REFERENCES outlets (id) ON DELETE CASCADE,
		FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	createStaffTable := `
	CREATE TABLE IF NOT EXISTS staff(
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		image TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
        role TEXT NOT NULL,
		address TEXT NOT NULL,
		phone TEXT NOT NULL,
		city TEXT NOT NULL,
		state TEXT NOT NULL,
        zip TEXT NOT NULL,
        country TEXT NOT NULL,
        outlet_id INT,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       	FOREIGN KEY (outlet_id) REFERENCES outlets (id) ON DELETE CASCADE
	)
	`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create user table " + err.Error())
	}

	_, err = DB.Exec(createOutlerTable)

	if err != nil {
		panic("Could not create Outlet table " + err.Error())
	}
	_, err = DB.Exec(createInventoryTable)

	if err != nil {
		panic("Could not create Inventory table " + err.Error())
	}
	_, err = DB.Exec(createProductTable)

	if err != nil {
		panic("Could not create Products table " + err.Error())
	}
	_, err = DB.Exec(createSalesTable)

	if err != nil {
		panic("Could not create Sales table " + err.Error())
	}
	_, err = DB.Exec(createStaffTable)

	if err != nil {
		panic("Could not create Staff table " + err.Error())
	}
}
