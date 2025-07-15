package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/thenameiswiiwin/go-workout/internal/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the backend server on")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/health", healthCheck)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Listening on port %d", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: OK")
}
