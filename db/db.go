package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:root@tcp(localhost:3308)/go-api-mysql"

var db *sql.DB

func Connect() {
	conection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	db = conection

	fmt.Println("La conexión a la BD ha sido exitosa")
}

func Close() {
	db.Close()
	fmt.Println("La conexión a la BD ha finalizado")
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
