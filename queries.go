package main

import (
	"database/sql"
	"fmt"
)

// User struct info contains data fro 1 row of users table
type User struct {
	ID       int64
	Name     string
	Contact  string
	Contact2 string
}

// Item struct info contains data fro 1 row of items table
type Item struct {
	ID            int64
	Name          string
	CurrentUserID int64
}

// UserByID queries for the user with a specific ID
func UserByID(queryUserid int64) (User, error) {
	// An album to hold data from the returned row.
	var u User

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", queryUserid)
	if err := row.Scan(&u.ID, &u.Name, &u.Contact, &u.Contact2); err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("UserByID %d: no such user", queryUserid)
		}
		return u, fmt.Errorf("UserByID %d: %v", queryUserid, err)
	}
	return u, nil
}

// ItemsByUser queries for items that are in the possession of a user (by user ID)
func ItemsByUser(queryUserid int64) ([]Item, error) {
	// an items slice to hold data from returned rows.
	var items []Item

	rows, err := db.Query("SELECT * FROM items WHERE current_user_id = ?", queryUserid)
	if err != nil {
		return nil, fmt.Errorf("ItemsByUser %q: %v", queryUserid, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name, &i.CurrentUserID); err != nil {
			return nil, fmt.Errorf("ItemsByUser %q: %v", queryUserid, err)
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ItemsByUser %q: %v", queryUserid, err)
	}
	return items, nil
}

// UsersByName queries for users that have the specified name.
func UsersByName(queryName string) ([]User, error) {
	// A users slice to hold data from returned rows.
	var users []User

	rows, err := db.Query("SELECT * FROM users WHERE name = ?", queryName)
	if err != nil {
		return nil, fmt.Errorf("UsersByName %q: %v", queryName, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Contact); err != nil {
			return nil, fmt.Errorf("UsersByName %q: %v", queryName, err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("usersByName %q: %v", queryName, err)
	}
	return users, nil
}
