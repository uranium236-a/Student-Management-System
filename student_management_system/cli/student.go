package cli

import (
	"fmt"
	"student_management_system/models"
	"student_management_system/session"
)

func student(institute *models.Institute) {
	fmt.Println("-------Student-------")

	currentStudent := institute.GetStudent(session.CurrentUser.Email)

	session.CurrentClass = institute.GetClass(currentStudent.Email)

	fmt.Println("\nClass: ", session.CurrentClass.Name)
	fmt.Println("Roll no: ", currentStudent.Rollno)
	fmt.Println("Name: ", session.CurrentUser.Name)
	fmt.Println("Email: ", session.CurrentUser.Email)
	fmt.Println("Attendance: ", currentStudent.Attendance.AverageAttendance)
	fmt.Println("Marks: ", currentStudent.Marks)

	fmt.Println("\n1. Inbox\n2. Exit")
	choice, _ := ReadInput(">> ")

	switch choice {
	case "1":
		fmt.Println("------------Inbox------------")
		for i := len(session.CurrentUser.Inbox) - 1; i >= 0; i-- {
			fmt.Println(session.CurrentUser.Inbox[i])
		}
		fmt.Println("------------------------------")

	case "2":
		Logout()

	}
}
