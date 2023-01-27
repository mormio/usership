package main

import (
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
		users = append(users, user)
		pages.SwitchToPage("Menu")
	})

	return form
}
