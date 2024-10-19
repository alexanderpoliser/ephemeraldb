package main

import (
	"ephemeraldb/internal/db"
	"ephemeraldb/internal/menu"
)

func main() {
	db := db.NewNoSQLDB()
	menu.RunMenu(db)
}
