package main

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

var form = tview.NewForm()

func AddUserForm() {
	user := User{}

	form.AddInputField("name", "", 20, nil, func(Name string) {
		user.Name = Name
	})

	form.AddInputField("contact", "", 20, nil, func(Contact string) {
		user.Contact = Contact
	})

	form.AddInputField("contact 2", "", 20, nil, func(Contact2 string) {
		user.Contact2 = Contact2
	})

	form.AddButton("Save", func() {
		id, _ := AddUser(user)
		_ = id
		pages.SwitchToPage("Menu")
	})
}

func AddItemForm() {
	item := Item{}

	form.AddInputField("name", "", 20, nil, func(Name string) {
		item.Name = Name
	})

	form.AddInputField("description", "", 20, nil, func(Description string) {
		item.Description = Description
	})

	form.AddInputField("current user id", "", 20, nil, func(CurrentUserID string) {
		u, err := strconv.Atoi(CurrentUserID)
		item.CurrentUserID = int64(u)
		if err != nil {
			fmt.Println("current user id should be int")
			return
		}
	})

	form.AddButton("Save", func() {
		id, _ := AddItem(item)
		_ = id
		pages.SwitchToPage("Menu")
	})
}
