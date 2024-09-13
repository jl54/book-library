package main

import (
	"book-library/db"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type SimpleResponse struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

func renderIndex(w http.ResponseWriter, r *http.Request) {
	var t, err = template.ParseFiles("templates/layout.html", "templates/body.html")

	if err != nil {
		log.Fatalf("Unable to parse files %v\n", err)
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Fatalf("Unable to parse template %v\n", err)
	}
}

func handlePublic() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
}

func main() {
	dbpool, err := db.NewPostgres()

	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	registerRoutes(dbpool)

	http.Handle("/public/", handlePublic())
	http.HandleFunc("/", renderIndex)
	// ListenAndServe always returns an error,
	// since it only returns when an unexpected error occurs.
	// In order to log that error we wrap the
	// function call with log.Fatal
	log.Fatal(http.ListenAndServe(":8080", nil))
}
