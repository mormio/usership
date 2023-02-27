package main

import (
	dbConnect "github.com/dopaminegirl19/usership/pkg/dbConnect"
	forms "github.com/dopaminegirl19/usership/pkg/forms"
	utils "github.com/dopaminegirl19/usership/pkg/utils"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Tview
var pages = tview.NewPages()
var itemText = tview.NewTextView()
var userText = tview.NewTextView()
var app = tview.NewApplication()
var form = tview.NewForm().
	SetFieldBackgroundColor(tcell.ColorRosyBrown).
	SetButtonBackgroundColor(tcell.ColorRosyBrown)
var itemsList = tview.NewList().ShowSecondaryText(false)
var usersList = tview.NewList().ShowSecondaryText(false)
var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorTomato).
	SetText("(u) to view users \n(i) to views items \n(q) to quit")

func main() {

	appDB := dbConnect.DBconnect()
	utils.AddItemsList(appDB)

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
			forms.AddUserForm(appDB)
			pages.SwitchToPage("Add User")
		case event.Rune() == 105:
			form.Clear(true)
			forms.AddItemForm(appDB)
			pages.SwitchToPage("Add Item")
		}
		return event
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("View Users", flex, true, true)
	pages.AddPage("Add User", form, true, false)
	pages.AddPage("View Items", flex, true, true)
	pages.AddPage("Add Item", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
