package app

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Run() {
	log.Println("Initializing..")

	log.Println("Reading config file..")
	initializeConfig()

	log.Println("Initializing DB..")
	initializeDB()
	defer closeDBConnection()

	log.Println("Initializing repositories..")
	initializeRepositories()

	log.Println("Initializing usecases..")
	initializeUsecases()

	log.Println("Initializing deliveries..")
	initializeDeliveries()

	router.Run(":9000")
}

func initializeConfig() {
	// try get config in local env
	err := godotenv.Load("../../config/.env")
	if err != nil {
		// get config in docker env
		err = godotenv.Load()
		if err != nil {
			log.Fatalln("[Error] Failed to get config file")
		}
	}
}

func initializeDB() {
	var err error
	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable connect_timeout=15",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		"5432",
	)

	dbPgClient, err = sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalln("[Error] Failed to connect to postgres", err)
	}
}

func closeDBConnection() {
	dbPgClient.Close()
}
