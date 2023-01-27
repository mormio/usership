package main

import (
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Tview
var pages = tview.NewPages()
var itemText = tview.NewTextView()
var userText = tview.NewTextView()
var app = tview.NewApplication()
var form = tview.NewForm().SetFieldBackgroundColor(tcell.ColorMediumVioletRed)
var itemsList = tview.NewList().ShowSecondaryText(false)
var usersList = tview.NewList().ShowSecondaryText(false)
var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreenYellow).
	SetText("(u) to add a new user \n(i) to add a new item \n(q) to quit")

func main() {

	DBconnect()
	AddItemsList()

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(itemsList, 0, 2, true).
			AddItem(itemText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
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

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Add User", form, true, false)
	pages.AddPage("Add Item", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
