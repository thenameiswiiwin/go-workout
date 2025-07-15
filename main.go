package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/thenameiswiiwin/go-workout/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("Application started successfully")

	http.HandleFunc("/health", healthCheck)
	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: OK")
}
