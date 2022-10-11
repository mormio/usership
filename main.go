package main

import (
	"github.com/rivo/tview"
)

var items []Item
var app = tview.NewApplication()

func main() {
	if err := app.SetRoot(tview.NewBox(), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
