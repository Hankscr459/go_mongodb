package main

import (
	"log"
	"twitter/db"
	"twitter/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin connect a la DB")
		return
	}
	handlers.Managements()
}
