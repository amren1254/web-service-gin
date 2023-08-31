package main

import (
	"database/sql"
	"fmt"

	"github.com/amren1254/gin-docker/database"
	"github.com/amren1254/gin-docker/router"

	"log"

	"github.com/joho/godotenv"
)

func init() {
	//initialize enviroment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading environment Vars - %v \n", err)
	}

}

func main() {
	fmt.Println(database.DatabaseRepository{DB: &sql.DB{}})
	r := router.InitRoute()
	r.Run(":8080")
}
