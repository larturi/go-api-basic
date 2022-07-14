package main

import (
	"packages/db"
	"packages/models"
)

func main() {
	db.Connect()
	db.Ping()

	db.CreateTable(models.UserSchema, "users")

	db.Close()
}
