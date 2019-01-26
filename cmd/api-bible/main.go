package main

import (
	"fmt"
	"log"
	"net/http"

	bible "github.com/alexfernandessd/api-bible/bible"
	"github.com/facebookgo/grace/gracehttp"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := bible.NewConfig()

	dbConn, dbErr := bible.NewConnectionMySQL(config)
	if dbErr != nil {
		log.Fatal("fail to connect with instance rds", dbErr)
	}

	service := bible.NewService(dbConn)
	handler := createServerHandler(service)

	fmt.Printf("Starting %s on port %d ...", config.APP, config.Port)
	err := gracehttp.Serve(&http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: handler,
	})

	if err != nil {
		log.Fatal("Failed to start de application", err)
	}

}
