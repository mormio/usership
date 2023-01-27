package main

import (
	"fmt"
	"strconv"
)

func AddUsersList() {
	usersList.Clear()

	rows, err := db.Query("SELECT id, name, contact, COALESCE(contact2, '') FROM users")
	if err != nil {
		fmt.Printf("ItemsByString: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Contact, &u.Contact2); err != nil {
			fmt.Printf("ItemsByUser: %v", err)
		}
		users = append(users, u)
		usersList.AddItem(strconv.Itoa(int(u.ID))+" "+u.Name, " ", rune('☺'), nil)
	}
}

func SetConcatText(user *User) {
	userText.Clear()
	text := user.Name + "\n" + user.Contact + "\n" + user.Contact2
	userText.SetText(text)
}
