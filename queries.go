package main

import "fmt"

// usersByName queries for users that have the specified name.
func UsersByName(query_name string) ([]User, error) {
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
