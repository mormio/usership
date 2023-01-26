package main

import (
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Tview
var userList = tview.NewList().ShowSecondaryText(false)
var app = tview.NewApplication()
var pages = tview.NewPages()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreenYellow).
	SetText("(u) to add a new user \n(i) to add a new item \n(q) to quit")

func main() {

	DBconnect()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch {
		case event.Rune() == 113:
			app.Stop()
		case event.Rune() == 117:
			form.Clear(true)
			AddUserForm()
			pages.SwitchToPage("Add User")
		case event.Rune() == 105:
			form.Clear(true)
			AddItemForm()
			pages.SwitchToPage("Add Item")
		}
		return event
	})

	pages.AddPage("Menu", text, true, true)
	pages.AddPage("Add User", form, true, false)
	pages.AddPage("Add Item", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
