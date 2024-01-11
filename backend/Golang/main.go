package main

import (
	"github.com/arioprima/Jobseeker/tree/main/backend/Golang/initializers"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//checkConeection to database success or not
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Println("Load config error", err)
	}
	_, err = initializers.ConnectDB(&config)
	if err != nil {
		log.Println("Connect to database error", err)
	} else {
		log.Println("Connect to database successfully")
	}
}
