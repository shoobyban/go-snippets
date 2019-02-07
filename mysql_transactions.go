package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", "test:test@(localhost:3306)/test")
	if err != nil {
		log.Fatalf("panic on connecting: %v", err)
	}
	defer db.Close()

	//	db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("can't create transaction: %v", err)
	}
	res, err := tx.Exec("INSERT INTO table2(table1_id, name) VALUES(?, ?)", 1, "other name")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	fmt.Printf("Res: %v", res)
	// commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatalf("can't roll transaction back: %v", err)
	}

	tx.Rollback()
}
