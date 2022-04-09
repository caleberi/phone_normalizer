package main

import (
	"fmt"
	phonedb "phone/db"

	_ "github.com/lib/pq"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	pass      = "postgres"
	dbName    = "normalized_phone_table"
	tableName = "phone_table"
)

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, pass)
	must(phonedb.ResetDB("postgres", psqlInfo, dbName))
	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbName)
	must(phonedb.Migrate("postgres", psqlInfo, dbName, tableName))
	db, err := phonedb.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	if err := db.Seed(tableName); err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Retrieving phones by ID .... 🍔")

	ids := []int{1, 2}
	for _, id := range ids {
		found, err := phonedb.GetPhoneById(db, tableName, id)
		must(err)
		fmt.Printf("found phone{%d}:%s\n", id, found)
	}

	fmt.Println()
	fmt.Println("Retrieving all phones ....🍯 ")
	allphones, err := phonedb.GetPhones(db, tableName)
	must(err)
	for _, phone := range allphones {
		fmt.Printf("🚩 %d %s \n", phone.Id, phone.Phone_number)
	}

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
