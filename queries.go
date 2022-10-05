package main

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID      int64
	Name    string
	Contact string
}

type Item struct {
	ID            int64
	Name          string
	CurrentUserID int64
}

func UserByID(query_userid int64) (User, error) {
	var user User

	row := db.Query("SELECT * FROM users WHERE id = ?", query_userid)
	if err := row.Scan(&user.ID, &user.Name, &user.Contact); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("UserByID %d: no such user", query_userid)
		}
		return user, fmt.Errorf("UserByID %d: %v", query_userid, err)
	}
	return user, nil
}

// ItemsByUser queries for items that are in the possession of a user (by user ID)
func ItemsByUser(query_userid int64) ([]Item, error) {
	// an items slice to hold data from returned rows.
	var items []Item

	rows, err := db.Query("SELECT * FROM items WHERE current_user_id = ?", query_userid)
	if err != nil {
		return nil, fmt.Errorf("ItemsByUser %q: %v", query_userid, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name, &i.CurrentUserID); err != nil {
			return nil, fmt.Errorf("ItemsByUser %q: %v", query_userid, err)
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ItemsByUser %q: %v", query_userid, err)
	}
	return items, nil
}

// UsersByName queries for users that have the specified name.
func UsersByName(query_name string) ([]User, error) {
	// A users slice to hold data from returned rows.
	var users []User

	rows, err := db.Query("SELECT * FROM users WHERE name = ?", query_name)
	if err != nil {
		return nil, fmt.Errorf("UsersByName %q: %v", query_name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Contact); err != nil {
			return nil, fmt.Errorf("UsersByName %q: %v", query_name, err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("usersByName %q: %v", query_name, err)
	}
	return users, nil
}
