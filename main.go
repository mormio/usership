package main

import (
	"strconv"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var users = make([]User, 0)

// Tview
var pages = tview.NewPages()
var userText = tview.NewTextView()
var app = tview.NewApplication()
var form = tview.NewForm()
var UsersList = tview.NewList().ShowSecondaryText(false)
var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreenYellow).
	SetText("(u) to add a new user \n(i) to add a new item \n(q) to quit")

func main() {

	DBconnect()

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(UsersList, 0, 1, true), 0, 6, true).
		AddItem(text, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch {
		case event.Rune() == 113:
			app.Stop()
		case event.Rune() == 117:
			form.Clear(true)
			addUserForm()
			pages.SwitchToPage("Add User")
		case event.Rune() == 105:
			form.Clear(true)
			// AddItemForm()
			pages.SwitchToPage("Add Item")
		}
		return event
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Add User", form, true, false)
	pages.AddPage("Add Item", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func addUsersList() {
	UsersList.Clear()
	for index, user := range users {
		UsersList.AddItem(user.Name+" "+user.Contact, " ", rune(49+index), nil)
	}
}

func addUserForm() *tview.Form {

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

func setConcatText(user *User) {
	userText.Clear()
	text := user.Name + " " + user.Contact + "\n" + user.Contact2
	userText.SetText(text)
}
