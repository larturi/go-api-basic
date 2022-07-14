package main

import (
	"fmt"
	"packages/db"
	"packages/models"
)

func main() {
	db.Connect()
	db.Ping()

	db.CreateTable(models.UserSchema, "users")

	// Crear Users
	// user := models.CreateUser("Lichi", "123456", "lichi@gmail.com")
	// fmt.Println(user)

	// Listado de Users
	users := models.ListUsers()
	fmt.Println(users)

	// Busco un User por Id
	user := models.GetUser(2)
	// Update user
	user.Username = "Candelaria"
	fmt.Println(user)

	db.Close()
}
