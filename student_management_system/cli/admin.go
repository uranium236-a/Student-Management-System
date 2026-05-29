package cli

import (
	"fmt"
	"student_management_system/models"
	"student_management_system/session"
)

func admin(institute *models.Institute) {
	fmt.Println("----------Admin Panel----------")

	fmt.Println("\n1. Profile")
	fmt.Println("2. Add User")
	fmt.Println("3. Broadcast")
	fmt.Println("4. View")
	fmt.Println("5. Inbox")
	fmt.Println("6. Assign Class")
	fmt.Println("7. Logout")

	choice, err := ReadInput(">> ")
	if err != nil {
		fmt.Println(err)
		return
	}

	switch choice {
	case "1":
		profile(institute)

	case "2":
		ch, _ := ReadInput("1. student\n2. teacher\n>> ")

		switch ch {
		case "1":
			rollno, name, email, password, class := ReadStudent()

			_, exists := institute.Classes[class]

			if exists {
				studentUser := institute.AddUser(name, email, password, "student")
				student := institute.Classes[class].AddStudent(rollno, *studentUser)
				fmt.Println("Student added successfully: ", student.Name)
			} else {
				fmt.Println("Class doesn't exist")
			}

		case "2":
			name, email, password, class := ReadTeacher()

			_, exists := institute.Classes[class]

			teacherUser := institute.AddUser(name, email, password, "teacher")

			if exists {
				teacher := institute.Classes[class].AssignTeacher(*teacherUser)
				fmt.Println("Teacher added successfully: ", teacher.Name)
			} else {
				choice2, _ := ReadInput("Class doesn't exist, create class?[y/n]: ")

				if choice2 == "y" {
					institute.CreateClass(class, nil)
					teacher := institute.Classes[class].AssignTeacher(*teacherUser)
					fmt.Println("Teacher added successfully: ", teacher.Name)
				}
			}
		}

	case "3":
		message, _ := ReadInput("Enter message: ")
		var to [2]string
		target, _ := ReadInput("to(all/admin/teacher/student/class)")

		to[0] = target

		if to[0] == "class" {
			classname, _ := ReadInput("Enter class name: ")
			to[1] = classname
		}

		from := [2]string{"admin", session.CurrentUser.Name}

		status := institute.BroadCastMessage(message, to, from)

		fmt.Println(status)

	case "4":
		view(institute)

	case "5":
		for i := len(session.CurrentUser.Inbox) - 1; i >= 0; i-- {
			fmt.Println(session.CurrentUser.Inbox[i])
		}

	case "6":
		class, _ := ReadInput("Enter class name: ")

		_, exists := institute.Classes[class]

		email, _ := ReadInput("Enter teacher email: ")
		teacher := institute.GetTeacherWithoutClass(email)

		if teacher == nil {
			fmt.Println("Teacher doesn't exist")
			return
		}

		if exists {
			fmt.Println("Class exists.. continuing the process.. ")
			teach := institute.Classes[class].AssignTeacher(teacher.User)
			fmt.Println("Teacher added successfully: ", teach.Name)
		} else {
			choice2, _ := ReadInput("Class doesn't exist, create class?[y/n]: ")

			if choice2 == "y" {
				institute.CreateClass(class, nil)
				teach := institute.Classes[class].AssignTeacher(teacher.User)
				fmt.Println("Teacher added successfully: ", teach.Name)
			}
		}

	case "7":
		Logout()
	}
}

func profile(institute *models.Institute) {
	fmt.Println("-------------Profile-------------")

	fmt.Println("\nAdmin ID: ", session.CurrentUser.ID)
	fmt.Println("Name: ", session.CurrentUser.Name)
	fmt.Println("Email: ", session.CurrentUser.Email)
	fmt.Println()
	fmt.Println("1. Edit Profile \t 2. Log Out")
	fmt.Println("3. Delete Account \t 4. Exit Profile")
	choice, err := ReadInput(">> ")
	if err != nil {
		fmt.Println(err)
		return
	}

	switch choice {
	case "1":
		fmt.Println("Enter 'no' for no change")
		name, _ := ReadInput("Enter name: ")
		password, _ := ReadInput("Enter password: ")

		if name != "no" {
			institute.Users[session.CurrentUser.ID].Name = name
		}

		if password != "no" {
			institute.Users[session.CurrentUser.ID].Password = password
		}
		fmt.Println("Profile edited successfully")

	case "2":
		Logout()
		return

	case "3":
		done := institute.RemoveUser(session.CurrentUser.ID)
		if done {
			fmt.Println("Account Deleted Successfully")
			Logout()
			return
		} else {
			fmt.Println("Something went wrong")
			Logout()
			return
		}

	case "4":
		return

	}

	profile(institute)
}

func view(institute *models.Institute) {

	fmt.Println("----------VIEW---------")
	fmt.Println("1. Classes")
	fmt.Println("2. Teachers")
	fmt.Println("3. Students")
	fmt.Println("4. Exit View")

	choice, _ := ReadInput(">> ")

	switch choice {
	case "1":
		fmt.Println("Class \t Teacher \t No. of Students")
		for _, class := range institute.Classes {
			fmt.Println(class.Name, "\t", class.Teacher.Name, "\t", len(class.Students))
		}

	case "2":
		fmt.Println("Class \t Teacher \t Email")
		for _, class := range institute.Classes {
			fmt.Println(class.Name, "\t", class.Teacher.Name, "\t", class.Teacher.Email)
		}

	case "3":
		choice2, _ := ReadInput("1. By Class\n2. Specific Student\n>> ")

		switch choice2 {
		case "1":
			className, _ := ReadInput("Enter class name: ")
			fmt.Println("Roll no\tName\tEmail\tAttendance\t Marks")
			for _, v := range institute.Classes[className].Students {
				fmt.Printf("%v\t%v\t%v\t%v\t%v\n",
					v.Rollno,
					v.Name,
					v.Email,
					v.Attendance.AverageAttendance,
					v.Marks,
				)
			}
		case "2":
			email, _ := ReadInput("Enter student email: ")
			student := institute.GetStudent(email)

			fmt.Println("Rollno: ", student.Rollno)
			fmt.Println("Name: ", student.Name)
			fmt.Println("Email: ", student.Email)
			fmt.Println("Attendance: ", student.Attendance.AverageAttendance)
			fmt.Println("Marks: ", student.Marks)
		}

	case "4":
		return
	}

	view(institute)
}
