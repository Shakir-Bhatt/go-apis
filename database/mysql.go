package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error Loading .env file", err)
	}
	return os.Getenv(key)
}

func NewMysqlConnection() *sql.DB {
	var (
		host     = getEnvVariable("DB_HOST")
		port     = getEnvVariable("DB_PORT")
		user     = getEnvVariable("DB_USERNAME")
		dbname   = getEnvVariable("DB_DATABASE")
		password = getEnvVariable("DB_PASSWORD")
	)

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	// defer db.Close()
	return db
}
