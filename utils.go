package main

func addUsersList() {
	usersList.Clear()
	for index, user := range users {
		usersList.AddItem(user.Name+" "+user.Contact, " ", rune(49+index), nil)
	}
}

func setConcatText(user *User) {
	userText.Clear()
	text := user.Name + " " + user.Contact + "\n" + user.Contact2
	userText.SetText(text)
}
