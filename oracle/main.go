package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "gopkg.in/rana/ora.v4"
)

func main() {

	db, err := sql.Open("oci8", "hr/hr2@example.com:1521/xe")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

}
