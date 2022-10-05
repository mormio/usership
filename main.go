// Written according to the tutorial at: https://go.dev/doc/tutorial/database-access

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	ID      int64
	Name    string
	Contact string
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "userships",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	users, err := usersByName("Morgane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Users found: %v\n", users)
}

// usersByName queries for users that have the specified name.
func usersByName(query_name string) ([]User, error) {
	// A users slice to hold data from returned rows.
	var users []User

	rows, err := db.Query("SELECT * FROM users WHERE name = ?", query_name)
	if err != nil {
		return nil, fmt.Errorf("usersByName %q: %v", query_name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Contact); err != nil {
			return nil, fmt.Errorf("usersByName %q: %v", query_name, err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("usersByName %q: %v", query_name, err)
	}
	return users, nil
}
