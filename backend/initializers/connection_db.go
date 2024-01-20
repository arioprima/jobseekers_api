package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func ConnectDB(config *Config) (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Connect to database error", err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("Ping to database error", err)
	} else {
		log.Println("Connect to database successfully")
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db, err
}
