package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	rollno   int
	name     string
	year     uint16
	subjects map[string]int
}

var students map[int]Student
var currentRollNo int = 1

func main() {
	students = make(map[int]Student)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== Student Management CLI =====")
		fmt.Println("1. Add Student")
		fmt.Println("2. Remove Student")
		fmt.Println("3. Display All Students")
		fmt.Println("4. Display Student Marks")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 1:
			addStudentCLI(reader)
		case 2:
			removeStudentCLI(reader)
		case 3:
			displayStudentsList()
		case 4:
			displayStudentMarksCLI(reader)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

// -------- CLI HANDLERS --------

func addStudentCLI(reader *bufio.Reader) {
	var s Student

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter year: ")
	yearInput, _ := reader.ReadString('\n')
	year64, _ := strconv.ParseUint(strings.TrimSpace(yearInput), 10, 16)
	year := uint16(year64)

	subjects := make(map[string]int)

	for {
		fmt.Print("Enter subject name (or 'done'): ")
		sub, _ := reader.ReadString('\n')
		sub = strings.TrimSpace(sub)

		if sub == "done" {
			break
		}

		fmt.Print("Enter marks: ")
		marksInput, _ := reader.ReadString('\n')
		marks, _ := strconv.Atoi(strings.TrimSpace(marksInput))

		subjects[sub] = marks
	}

	s.addStudent(name, year, subjects)
	students[s.rollno] = s

	fmt.Println("Student added with Roll No:", s.rollno)
}

func removeStudentCLI(reader *bufio.Reader) {
	fmt.Print("Enter roll number to remove: ")
	input, _ := reader.ReadString('\n')
	rollno, _ := strconv.Atoi(strings.TrimSpace(input))

	removeStudent(rollno)
	fmt.Println("Student removed (if existed).")
}

func displayStudentMarksCLI(reader *bufio.Reader) {
	fmt.Print("Enter roll number: ")
	input, _ := reader.ReadString('\n')
	rollno, _ := strconv.Atoi(strings.TrimSpace(input))

	displayStudentMarks(rollno)
}

// -------- CORE FUNCTIONS --------

func (s *Student) addStudent(name string, year uint16, subjects map[string]int) {
	s.rollno = currentRollNo
	currentRollNo++
	s.name = name
	s.year = year
	s.subjects = subjects
}

func removeStudent(rollno int) {
	delete(students, rollno)
}

func displayStudentsList() {
	fmt.Println("\nRollno\tName\t\tYear")
	for _, student := range students {
		fmt.Println(student.rollno, "\t", student.name, "\t\t", student.year)
	}
}

func displayStudentMarks(rollno int) {
	student, exists := students[rollno]
	if !exists {
		fmt.Println("Student not found")
		return
	}

	fmt.Println("\n", student.rollno, "\t", student.name)
	for subject, marks := range student.subjects {
		fmt.Println(subject, ":", marks)
	}
}
