package cli

import (
	"fmt"
	"student_management_system/models"
	"student_management_system/session"
)

func Signup(i *models.Institute) *models.User {

	fmt.Println("-----Sign Up-----")

	role, _ := ReadInput("Enter role\nadmin\nteacher\n>> ")

	name, role, email, password, err := ReadUser(role)

	if err != nil {
		return nil
	}

	for _, v := range i.Users {
		if v.Email == email {
			fmt.Println("Account already exists")
			return nil
		}
	}

	if role == "admin" {
		inputKey, err := ReadInput("Enter Admin Access Key: ")
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if models.CheckAdminAccess(inputKey) {
			return i.AddUser(name, email, password, role)

		} else {
			fmt.Println("Access Denied")
		}
	} else if role == "teacher" {
		return i.AddUser(name, email, password, role)
	} else {
		fmt.Println("Invalid role choice")
	}

	return nil
}

func Login(i *models.Institute) *models.User {

	email, _ := ReadInput("Enter email: ")
	password, _ := ReadInput("Enter password: ")

	for _, v := range i.Users {
		if v.Email == email {
			if v.Password == password {
				return v
			}
		}
	}

	return nil

}

func Logout() {
	session.CurrentUser = nil
}
