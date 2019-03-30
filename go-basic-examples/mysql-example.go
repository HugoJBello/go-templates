package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type ScrapingLog struct {
	scraping_id   string    `json:"scraping_id"`
	last_piece   string    `json:"last_piece"`
	last_result string `json:"last_result"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/scraping")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	results, err := db.Query("select * from scraping_execution_log")


	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag ScrapingLog
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.scraping_id, &tag.last_piece, &tag.last_result)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.scraping_id)
	}

	// be careful deferring Queries if you are using transactions
	defer results.Close()


}