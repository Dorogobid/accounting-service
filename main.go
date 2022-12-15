package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/Dorogobid/accounting-service/delivery/httpserver"
	"github.com/Dorogobid/accounting-service/internal/domain/usecases"
	_ "github.com/lib/pq"
)

//Main function
func main() {
	ctx := context.Background()	
	db, err := sql.Open("postgres", "user=as-user password=postgres dbname=as-db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	uc := usecases.New(ctx, db)
	
	server := httpserver.New(uc)
	server.Start()
}