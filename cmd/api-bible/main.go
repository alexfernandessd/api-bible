package main

import (
	"fmt"
	"log"

	bible "github.com/alexfernandessd/api-bible/bible"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := bible.NewConfig()

	dbConn, dbErr := bible.NewConnectionMySQL(config)

	if dbErr != nil {
		log.Fatal("fail to connect with instance rds", dbErr)
	}

	connErr := dbConn.Ping()

	if connErr != nil {
		log.Fatal("failed to connect on database: ", connErr)
	}

	// Read to start de application

	rows, err := dbConn.Query("SELECT * FROM verses where id = 4000")

	for rows.Next() {
		var id int
		var version string
		var testament int
		var book int
		var chapter int
		var verse int
		var text string
		err = rows.Scan(&id, &version, &testament, &book, &chapter, &verse, &text)
		fmt.Printf("chapter: %d, verse: %d: ", chapter, verse)
		fmt.Println(text)
	}

	if err != nil {
		fmt.Println("error: ", err)
	}
}
