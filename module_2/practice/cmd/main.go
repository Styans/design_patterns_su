package main

import "practice/internal/user"

func main() {
	manager := user.UserManager{}

	manager.AddUser("Стас", "sty@example.com", "Admin")
	manager.AddUser("Мадина", "madina@example.com", "User")

	manager.ListUsers()

	manager.UpdateUser("madina@example.com", "Мадина С.", "Admin")
	manager.RemoveUser("sty@example.com")

	manager.ListUsers()
}
