package main

import (
	"log"

	"github.com/backupsGit/pryml/db"
)

func main() {
	if db.CheckConnection() {
		log.Fatal("Problema de conexión con la BD")
	}
	//handlers.Manejadores()
}
