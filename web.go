package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var port string = "3000"

func message(format string, a ...interface{}) {
	fmt.Printf(fmt.Sprintf("tug-example: %s", format), a...)
}

func initializeDatabase() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=postgres dbname=postgres sslmode=disable", host, port))

	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS fruits (name varchar UNIQUE, color varchar);")

	if err != nil {
		return nil, err
	}

	db.Exec("INSERT INTO fruits (name, color) VALUES ('apple', 'red');")
	db.Exec("INSERT INTO fruits (name, color) VALUES ('orange', 'orange');")
	db.Exec("INSERT INTO fruits (name, color) VALUES ('lemon', 'yellow');")

	return db, nil
}

func main() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	db, err := initializeDatabase()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to postgres: %s\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!\n")
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "Environment:\n")
		for _, env := range os.Environ() {
			fmt.Fprintf(w, "%s\n", env)
		}
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "Database:\n")
		fruits, err := db.Query("SELECT * FROM fruits ORDER BY name")
		if err != nil {
			fmt.Fprintf(w, "Database Error: %s\n", err)
		} else {
			for fruits.Next() {
				var name, color string
				fruits.Scan(&name, &color)
				fmt.Fprintf(w, "%s\t%s\n", name, color)
			}
		}
	})

	message("listening on port %s\n", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
