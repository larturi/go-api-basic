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

func ExistsTables(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)

	if err != nil {
		fmt.Println(err)
	}

	return rows.Next()
}

func CreateTable(schemma string, tableName string) {
	if !ExistsTables(tableName) {
		_, err := db.Exec(schemma)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	db.Exec(sql)
}

//Polimorfismo a Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()

	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// Polimorfismo a Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()

	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
