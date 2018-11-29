package core

import (
	"github.com/joho/godotenv"
)

func Application() {
	// load .env environment variables
	error := godotenv.Load("environments/.development.env")

	if error != nil {
		panic(error)
	}

	// initializes database
	connection, _ := SetupDatabase()

	Router(connection)
}
