package models

type Class struct {
	Name     string
	Teacher  *Teacher
	Students map[string]*Student
	DayCount int
}

func (c *Class) AddStudent(rollno int, user User) Student {

	newStudent := &Student{
		User:   user,
		Rollno: rollno,
	}

	c.Students[user.Email] = newStudent

	return *newStudent

}

func (c *Class) AssignTeacher(teacher User) Teacher {

	assignedTeacher := &Teacher{teacher, true}

	c.Teacher = assignedTeacher

	return *assignedTeacher
}

func (c *Class) GetStudentByRollNo(rollno int) *Student {
	for _, student := range c.Students {
		if student.Rollno == rollno {
			return student
		}
	}
	return nil
}
