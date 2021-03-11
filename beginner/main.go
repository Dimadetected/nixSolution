package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	user := "postgres"
	password := "postgres"
	port := "5432"
	host := "localhost"
	sslmode := "disable"
	dbname := "postgres"

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s dbname=%s",
		host, user, password, port, sslmode, dbname))
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	var id int
	var name string

	rows, err := db.Query("SELECT * FROM remote_cars")
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(id, name)
	}

	row := db.QueryRow("SELECT * FROM remote_cars WHERE id = 1")
	row.Scan(&id, &name)
	fmt.Println(id, name)
	row = db.QueryRow("INSERT INTO remote_cars (status) VALUES (false) RETURNING id")
	if err := row.Scan(&id); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(id)

	_, err = db.Exec("DELETE FROM remote_cars WHERE id IN (7,8,9,10,11)")
	if err != nil {
		log.Fatal(err.Error())
	}

}
