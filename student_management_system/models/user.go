package models

type User struct {
	ID       int
	Name     string
	Role     string
	Email    string
	Password string //Subject to change
	Inbox    []string
	//History  []string
}
