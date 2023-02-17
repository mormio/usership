package forms

import (
	"strconv"

	queries "github.com/dopaminegirl19/usership/pkg/queries"
	structs "github.com/dopaminegirl19/usership/pkg/structs"
	utils "github.com/dopaminegirl19/usership/pkg/utils"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var pages = tview.NewPages()
var form = tview.NewForm().
	SetFieldBackgroundColor(tcell.ColorRosyBrown).
	SetButtonBackgroundColor(tcell.ColorRosyBrown)

func AddUserForm() *tview.Form {

	user := structs.User{}
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
		id, _ := queries.AddUser(user)
		_ = id
		utils.AddUsersList()
		pages.SwitchToPage("Menu")
	})

	form.AddButton("Cancel", func() {
		pages.SwitchToPage("Menu")
	})

	return form
}

func AddItemForm() *tview.Form {

	item := structs.Item{}
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
		id, _ := queries.AddItem(item)
		_ = id
		utils.AddItemsList()
		pages.SwitchToPage("Menu")
	})

	form.AddButton("Cancel", func() {
		pages.SwitchToPage("Menu")
	})

	return form
}
