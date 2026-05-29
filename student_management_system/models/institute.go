package models

import "time"

type Institute struct {
	Name    string            `json:"name"`
	Classes map[string]*Class `json:"classes"`
	Users   map[int]*User     `json:"users"`
	NextID  int               `json:"next_id"`
}

func (i *Institute) AddUser(name, email, password, role string) *User {
	id := i.NextID

	i.Users[id] = &User{
		ID:       id,
		Name:     name,
		Role:     role,
		Email:    email,
		Password: password,
	}

	i.NextID++

	return i.Users[id]
}

func (i *Institute) CreateClass(name string, teacher *Teacher) {
	i.Classes[name] = &Class{
		Name:     name,
		Teacher:  teacher,
		Students: make(map[string]*Student),
	}
}

func (i *Institute) RemoveUser(ID int) bool {
	for _, v := range i.Users {
		if v.ID == ID {
			delete(i.Users, v.ID)
			return true
		}
	}
	return false
}

func (i *Institute) GetUserID(email string) int {
	for _, v := range i.Users {
		if v.Email == email {
			return v.ID
		}
	}

	return -1
}

func (i *Institute) GetStudent(email string) *Student {
	for _, v := range i.Classes {
		for _, student := range v.Students {
			if student.Email == email {
				return student
			}
		}
	}

	return nil
}

func (i *Institute) GetTeacher(email string) *Teacher {
	for _, v := range i.Classes {
		if v.Teacher.Email == email {
			return v.Teacher
		}
	}

	return nil
}

func (i *Institute) GetClass(email string) *Class {

	for _, v := range i.Classes {
		if v.Teacher.Email == email {
			return v
		}
	}

	for _, v := range i.Classes {
		for _, student := range v.Students {
			if student.Email == email {
				return v
			}
		}
	}

	return nil
}

func (i *Institute) GetTeacherWithoutClass(email string) *Teacher {

	for _, v := range i.Users {
		if v.Email == email && v.Role == "teacher" {
			return &Teacher{*v, false}
		}
	}
	return nil
}

func (i *Institute) BroadCastMessage(message string, to, from [2]string) string {

	//from[0] - role, from[1] - name
	//to[0]- role, to[1]- class
	now := time.Now()

	now = now.Truncate(time.Minute)

	dateTimeStr := now.Format("2006-01-02 15:04")

	if from[0] == "teacher" && (to[0] == "all") {
		return "Permission not granted"
	}

	messageFinal := from[0] + ":" + from[1] + " :- " + message + " [" + dateTimeStr + "]"

	switch to[0] {
	case "all":
		for _, v := range i.Users {
			v.Inbox = append(v.Inbox, messageFinal)
		}

	case "admin":
		for _, v := range i.Users {
			if v.Role == "admin" {
				v.Inbox = append(v.Inbox, messageFinal)
			}
		}

	case "teacher":
		for _, v := range i.Users {
			if v.Role == "teacher" {
				v.Inbox = append(v.Inbox, messageFinal)
			}
		}

	case "student":
		for _, v := range i.Users {
			if v.Role == "student" {
				v.Inbox = append(v.Inbox, messageFinal)
			}
		}

	case "class":
		for _, class := range i.Classes {
			if class.Name == to[1] {

				for _, u := range i.Users {
					if u.Email == class.Teacher.Email {
						u.Inbox = append(u.Inbox, messageFinal)
					}
				}

				for _, student := range class.Students {
					for _, u := range i.Users {
						if u.Email == student.Email {
							u.Inbox = append(u.Inbox, messageFinal)
						}
					}
				}
			}
		}

	}

	return "Message broadcasted successfully"
}
