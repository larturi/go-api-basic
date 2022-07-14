package main

import "packages/db"

func main() {
	db.Connect()
	db.Ping()
	db.Close()
}
