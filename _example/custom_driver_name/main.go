package main

import (
	"database/sql"

	_ "github.com/rysa5/modxx"
)

func main() {
	for _, driver := range sql.Drivers() {
		println(driver)
	}
}
