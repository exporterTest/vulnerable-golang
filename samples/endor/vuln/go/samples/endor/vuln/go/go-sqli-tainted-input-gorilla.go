package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

func main() {
	// 1. Using schema.Decoder - Unsanitized Input
	var user User
	decoder := schema.NewDecoder()
	req := http.Request{}                                  // Simulate a request
	req.URL.Query().Add("username", "malicious' OR 1=1--") // Unsanitized input
	err := decoder.Decode(&user, req.URL.Query())
	if err != nil {
		fmt.Println(err)
		return
	}
	// Vulnerable SQL query:
	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", user.Username)
	// ruleid: go-sqli-tainted-input-gorilla
	_, err = db.Exec(query) // Assuming db is a database connection
	if err != nil {
		// Handle error
	}

}

type User struct {
	Username string `schema:"username"`
}

func runStartupQueries(db *sql.DB, startupQueries string) error {
	// ok: go-sqli-tainted-input-gorilla
	if _, err := db.Exec(startupQueries); err != nil {
		return err
	}
	// ok: go-sqli-tainted-input-gorilla
	if _, err := db.Exec(builtinStartupQueries); err != nil {
	}
}
