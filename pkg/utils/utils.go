package utils

import (
	queries "github.com/dopaminegirl19/usership/pkg/queries"
	structs "github.com/dopaminegirl19/usership/pkg/structs"

	"fmt"
	"strconv"

	tview "github.com/rivo/tview"
)

var itemText = tview.NewTextView()
var userText = tview.NewTextView()
var itemsList = tview.NewList().ShowSecondaryText(false)
var usersList = tview.NewList().ShowSecondaryText(false)

func AddUsersList() {
	var users = make([]structs.User, 0)
	usersList.Clear()

	rows, err := db.Query("SELECT id, name, contact, COALESCE(contact2, '') FROM users")
	if err != nil {
		fmt.Printf("AddUsersList: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var u structs.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Contact, &u.Contact2); err != nil {
			fmt.Printf("AddUsersList: %v", err)
		}
		users = append(users, u)
		usersList.AddItem(strconv.Itoa(int(u.ID))+" "+u.Name, " ", rune('☺'), nil)
	}

	usersList.SetSelectedFunc(func(index int, name string, contact string, shortcut rune) {
		SetUserText(&users[index])
	})
}

func SetUserText(user *structs.User) {
	userText.Clear()
	text := user.Name + "\n" + user.Contact + "\n" + user.Contact2
	userText.SetText(text)
}

func AddItemsList() {
	var items = make([]structs.Item, 0)
	itemsList.Clear()

	rows, err := db.Query("SELECT id, name, COALESCE(description, '') as description, current_user_id FROM items")
	if err != nil {
		fmt.Printf("AddItemsList: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var i structs.Item
		if err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.CurrentUserID); err != nil {
			fmt.Printf("AddItemsList: %v", err)
		}
		items = append(items, i)
		itemsList.AddItem(strconv.Itoa(int(i.ID))+" "+i.Name, " ", rune('☺'), nil)
	}

	itemsList.SetSelectedFunc(func(index int, name string, description string, shortcut rune) {
		SetItemText(&items[index])
	})
}

func SetItemText(item *structs.Item) {
	itemText.Clear()

	// Get current user info
	current_user, _ := queries.UserByID(int32(item.CurrentUserID))

	text := item.Name + "\n" + item.Description + "\n\nCurrent user: \n" + current_user.Name + "\n" + current_user.Contact + "\n" + current_user.Contact2

	// Set text object
	itemText.SetText(text)
}
