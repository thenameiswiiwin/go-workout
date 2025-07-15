package main

import "github.com/thenameiswiiwin/go-workout/internal/app"

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("Application started successfully")
}
