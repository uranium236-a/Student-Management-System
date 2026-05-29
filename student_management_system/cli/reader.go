package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(prompt string) (string, error) {
	fmt.Print("\n", prompt)

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	return input, nil
}

func ReadUser(role string) (string, string, string, string, error) {
	var err error
	name, err := ReadInput("Enter your name: ")
	if err != nil {
		fmt.Println(err)
		return "", "", "", "", err
	}

	email, err := ReadInput("Enter your email: ")
	if err != nil {
		fmt.Println(err)
		return "", "", "", "", err
	}

	password, err := ReadInput("Enter your password: ")
	if err != nil {
		fmt.Println(err)
		return "", "", "", "", err
	}

	return name, role, email, password, nil
}

func ReadStudent() (int, string, string, string, string) {
	rollnos, _ := ReadInput("Roll no: ")
	rollno, _ := strconv.Atoi(rollnos)
	name, _, email, password, _ := ReadUser("student")
	class, _ := ReadInput("Class assigned: ")

	return rollno, name, email, password, class
}

func ReadTeacher() (string, string, string, string) {
	name, _, email, password, _ := ReadUser("teacher")
	class, _ := ReadInput("Class assigned: ")

	return name, email, password, class
}
