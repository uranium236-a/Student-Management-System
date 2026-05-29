package cli

import (
	"fmt"
	"strconv"
	"strings"
	"student_management_system/models"
	"student_management_system/session"
)

func teacher(institute *models.Institute) {
	fmt.Println("--------Teacher--------")

	currentTeacher := institute.GetTeacher(session.CurrentUser.Email)

	if currentTeacher == nil {
		currentTeacher = institute.GetTeacherWithoutClass(session.CurrentUser.Email)
	}

	if !currentTeacher.ClassAssigned {
		fmt.Println("No class assigned")

		className, _ := ReadInput("Enter class name(you want to be assigned): ")

		msg := institute.BroadCastMessage(currentTeacher.AssignClassRequest(className))

		fmt.Println(msg)

		fmt.Println("Login again after some time when class is assigned")

		Logout()
		return
	}

	session.CurrentClass = institute.GetClass(currentTeacher.Email)

	fmt.Println("Teacher ID: ", session.CurrentUser.ID)
	fmt.Println("Name: ", session.CurrentUser.Name)
	fmt.Println("Email: ", session.CurrentUser.Email)
	fmt.Println("Class: ", session.CurrentClass.Name)

	fmt.Println("\n1. Inbox\n2. Broadcast Message\n3. View\n4. Attendance\n5. Allocate Marks\n6. Logout")
	choice, _ := ReadInput(">> ")

	switch choice {
	case "1":
		fmt.Println("------------Inbox------------")
		for i := len(session.CurrentUser.Inbox) - 1; i >= 0; i-- {
			fmt.Println(session.CurrentUser.Inbox[i])
		}
		fmt.Println("-----------------------------")
	case "2":
		msg, _ := ReadInput("Enter message: ")
		to := [2]string{"class", session.CurrentClass.Name}
		from := [2]string{"teacher", session.CurrentUser.Name}

		status := institute.BroadCastMessage(msg, to, from)

		fmt.Println(status)

	case "3":
		fmt.Println("Roll no\tName\tEmail\tAttendance\tMarks")

		for _, v := range institute.Classes[session.CurrentClass.Name].Students {
			fmt.Printf("%v\t%v\t%v\t%v\t%v\n",
				v.Rollno,
				v.Name,
				v.Email,
				v.Attendance.AverageAttendance,
				v.Marks,
			)
		}

	case "4":
		fmt.Println("--------Attendance------")
		fmt.Println("1. Mark Attendance")
		fmt.Println("2. View student attendance")
		choice2, _ := ReadInput(">> ")

		switch choice2 {
		case "1":
			markAttendance(institute)
		case "2":
			rollnostr, _ := ReadInput("Enter rollno: ")
			rollno, err := strconv.Atoi(rollnostr)

			if err != nil {
				fmt.Println("Invalid rollno")
			} else {
				student := institute.Classes[session.CurrentClass.Name].GetStudentByRollNo(rollno)
				fmt.Println("Name: ", student.Name)
				fmt.Println("Average attendance: ", student.Attendance.AverageAttendance)
				fmt.Println("Attendance list: ")
				fmt.Println(student.Attendance.AttendanceList)
			}
		}

	case "5":
		rollnostr, _ := ReadInput("Enter rollno: ")
		rollno, err := strconv.Atoi(rollnostr)
		if err != nil {
			fmt.Println("Invalid Rollno")
			return
		}
		marksstr, _ := ReadInput("Enter marks: ")

		marks, err := strconv.ParseFloat(marksstr, 32)

		if err != nil {
			fmt.Println("Invalid Marks")
			return
		}

		student := institute.Classes[session.CurrentClass.Name].GetStudentByRollNo(rollno)

		institute.Classes[session.CurrentClass.Name].Students[student.Email].Marks = float32(marks)

	case "6":
		Logout()
	}
}

func markAttendance(institute *models.Institute) {

	absentStr, _ := ReadInput("Enter absent roll numbers(space-separated between): ")

	absentSL := strings.Fields(absentStr)

	for _, v := range institute.Classes[session.CurrentClass.Name].Students {

		isAbsent := false

		for _, a := range absentSL {
			rollno, err := strconv.Atoi(a)

			if err != nil {
				continue
			}

			if v.Rollno == rollno {
				isAbsent = true
			}
		}

		v.Attendance.AttendanceList = append(v.Attendance.AttendanceList, !(isAbsent))
	}

	institute.Classes[session.CurrentClass.Name].DayCount++

	fmt.Println("Attendance Marked Successfully")

	//Calculating Averages:

	for _, v := range institute.Classes[session.CurrentClass.Name].Students {

		var sum int = 0

		for i := 0; i < len(v.Attendance.AttendanceList); i++ {
			if v.Attendance.AttendanceList[i] {
				sum++
			}
		}

		v.Attendance.AverageAttendance = float32((sum / institute.Classes[session.CurrentClass.Name].DayCount) * 100)
	}

	fmt.Println("Average Attendance calculated successfully")
}
