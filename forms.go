package main

import (
	"strconv"

	"github.com/rivo/tview"
)

func AddUserForm() *tview.Form {

	user := User{}

	form.AddInputField("user id", "", 20, nil, func(userID string) {
		ID, _ := strconv.Atoi(userID)
		user.ID = int64(ID)
	})

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
		users = append(users, user)
		addUsersList()
		pages.SwitchToPage("Menu")
	})

	return form
}
