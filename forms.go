package main

import (
	"strconv"

	"github.com/rivo/tview"
)

func AddUserForm() *tview.Form {

	user := User{}
	form.AddInputField("name", "", 20, nil, func(Name string) {
		user.Name = Name
	})

	form.AddInputField("contact", "", 20, nil, func(Contact string) {
		user.Contact = Contact
	})

	form.AddInputField("contact2", "", 20, nil, func(Contact2 string) {
		user.Contact2 = Contact2
	})

	form.AddButton("Save", func() {
		// users = append(users, user)
		id, _ := AddUser(user)
		_ = id
		AddUsersList()
		pages.SwitchToPage("Menu")
	})

	form.AddButton("Cancel", func() {
		pages.SwitchToPage("Menu")
	})

	return form
}

func AddItemForm() *tview.Form {

	item := Item{}
	form.AddInputField("name", "", 20, nil, func(Name string) {
		item.Name = Name
	})

	form.AddInputField("description", "", 20, nil, func(Description string) {
		item.Description = Description
	})

	form.AddInputField("current user id", "", 20, nil, func(CurrentUserID string) {
		id, _ := strconv.Atoi(CurrentUserID)
		item.CurrentUserID = int32(id)
	})

	form.AddButton("Save", func() {
		// users = append(users, user)
		id, _ := AddItem(item)
		_ = id
		AddItemsList()
		pages.SwitchToPage("Menu")
	})

	form.AddButton("Cancel", func() {
		pages.SwitchToPage("Menu")
	})

	return form
}
