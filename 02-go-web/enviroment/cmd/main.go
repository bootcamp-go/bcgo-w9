package main

import (
	"app/cmd/application"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	os.Setenv("NEW_KEY", "NEW_VALUE")

	username := os.Getenv("APP_USERNAME")
	storageClimatePath := os.Getenv("APP_FILEPATH")

	// application
	app := application.NewApplication(username, storageClimatePath)

	// run
	if err := app.Run(); err != nil {
		panic(err)
	}
}