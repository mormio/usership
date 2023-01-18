package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Contact struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	business    bool
}

var contacts []Contact

var items []Item

var form = tview.NewForm()

var app = tview.NewApplication()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreenYellow).
	SetText("(a) to add a new contact \n(q) to quit")

func main() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 97 {
			addContactForm()
		}
		return event
	})
	if err := app.SetRoot(text, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func addContactForm() {
	contact := Contact{}

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		contact.firstName = firstName
	})

	form.AddInputField("Last Name", "", 20, nil, func(lastName string) {
		contact.lastName = lastName
	})

	form.AddInputField("Email", "", 20, nil, func(email string) {
		contact.email = email
	})

	form.AddInputField("Phone", "", 20, nil, func(phone string) {
		contact.phoneNumber = phone
	})

	form.AddCheckbox("Business", false, func(business bool) {
		contact.business = business
	})

	form.AddButton("Save", func() {
		contacts = append(contacts, contact)
	})
}
