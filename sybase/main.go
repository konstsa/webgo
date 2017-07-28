package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	db, err := sql.Open("mssql", "server=10.1.108.203;port=5000;user id=test;password=123456;database=oper")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("norm")
	}
	defer db.Close()

	// Run query with multiple return rows
	rows, err := db.Query("select count(1) as cnt from oper..tCounter")
	if err != nil {
		log.Fatalf("Select failed1: %s", err)
	}
	for rows.Next() {
		var Id int
		err = rows.Scan(&Id)
		if err != nil {
			log.Fatalf("Select failed2: %s", err)
		}
		log.Printf("Id: %0.2f", Id)
	}
	//defer rows.Close()

}
