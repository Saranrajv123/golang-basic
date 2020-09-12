package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type UserData struct {
	Name           string `json:"name"`
	CountryOfBirth string `json:"country_of_birth"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "test"
)

func main() {
	fmt.Println("go postgresql")

	psqlQuery := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlQuery, "psqlquery")

	db, err := sql.Open("postgres", psqlQuery)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Successfully connected")

	selectedResult, err := db.Query("select first_name, country_of_birth from person")

	if err != nil {
		panic(err)
	}

	for selectedResult.Next() {
		var user UserData

		err = selectedResult.Scan(&user.Name, &user.CountryOfBirth)
		if err != nil {
			panic(err)
		}

		fmt.Println(user, "userData")
	}

	// insert, err := db.Query("insert into person values('1001','Saran', 'raj', 'saranraj@gmail.com', 'Male', '1990-08-16', 'India')")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(insert, "insert")

	// fmt.Println("Successfully inserted")
}
