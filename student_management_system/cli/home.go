package cli

import (
	"fmt"
	"student_management_system/models"
	"student_management_system/session"
)

func Home(institute *models.Institute) {

	if session.CurrentUser == nil {
		choice, err := ReadInput("\nSelect your choice\n1. Signup\n2. Login\n3. Exit\n>> ")
		if err != nil {
			fmt.Println(err)
		}
		switch choice {
		case "1":
			session.CurrentUser = Signup(institute)
		case "2":
			session.CurrentUser = Login(institute)
			if session.CurrentUser == nil {
				fmt.Println("User doesn't exists")
				return
			}

		case "3":
			session.Running = false
			return
		}
	} else {
		switch session.CurrentUser.Role {
		case "admin":
			admin(institute)
		case "teacher":
			teacher(institute)
		case "student":
			student(institute)
		}

	}

}
