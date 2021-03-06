package main

import (
	"fmt"
	"log"

	"github.com/backupsGit/pryml/Models"
	"github.com/backupsGit/pryml/db"
)

func main() {
	if db.CheckConnection() {
		log.Fatal("Problema de conexión con la BD")
	}
	//handlers.Manejadores()

	var tm = [][]string{
		{"A", "A", "A", "A", "04", "05", "06", "07", "08"},
		{"A", "11", "12", "13", "14", "15", "16", "17", "18"},
		{"A", "21", "B", "23", "24", "25", "26", "27", "28"},
		{"A", "31", "32", "B", "34", "35", "36", "37", "38"},
		{"40", "41", "42", "43", "B", "45", "46", "47", "48"},
		{"50", "51", "52", "53", "54", "B", "56", "57", "58"},
		{"60", "61", "62", "63", "64", "65", "66", "67", "68"},
		{"70", "71", "72", "73", "74", "75", "76", "77", "78"},
		{"80", "81", "82", "83", "84", "85", "86", "87", "88"}}

	m := Models.MutantsModel{Matrix: tm, Len: 3}.GetMutants()

	fmt.Println("Matrix:", tm)
	fmt.Println("Matrix:", m)
}
