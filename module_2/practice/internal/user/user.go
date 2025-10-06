package user

import "fmt"

type User struct {
	Name  string
	Email string
	Role  string
}

type UserManager struct {
	users []User
}

func (m *UserManager) AddUser(name, email, role string) {
	if m.findIndexByEmail(email) != -1 {
		fmt.Println("User already exists.")
		return
	}
	m.users = append(m.users, User{Name: name, Email: email, Role: role})
	fmt.Println("User added:", name)
}

func (m *UserManager) RemoveUser(email string) {
	index := m.findIndexByEmail(email)
	if index == -1 {
		fmt.Println("User not found.")
		return
	}
	m.users = append(m.users[:index], m.users[index+1:]...)
	fmt.Println("User removed:", email)
}

func (m *UserManager) UpdateUser(email, newName, newRole string) {
	index := m.findIndexByEmail(email)
	if index == -1 {
		fmt.Println("User not found.")
		return
	}
	m.users[index].Name = newName
	m.users[index].Role = newRole
	fmt.Println("User updated:", newName)
}

func (m *UserManager) ListUsers() {
	for _, u := range m.users {
		fmt.Printf("Name: %s | Email: %s | Role: %s\n", u.Name, u.Email, u.Role)
	}
}

func (m *UserManager) findIndexByEmail(email string) int {
	for i, u := range m.users {
		if u.Email == email {
			return i
		}
	}
	return -1
}
