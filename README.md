# MyApp - Student Management System

A command-line interface (CLI) application for managing educational institution, built in Go. The system supports multiple user roles including students, teachers, and administrators.

## 📁 Project Structure
student_management_system:.  
│   go.mod   
│   main.go    
│   myapp.exe    
│    
├───cli   
│       admin.go   
│       auth.go   
│       home.go   
│       reader.go   
│       student.go   
│       teacher.go  
│         
├───models  
│       admin.go  
│       class.go  
│       institute.go  
│       student.go  
│       teacher.go  
│       user.go  
│         
├───session  
│       session.go  
│         
└───storage  
  
## 🚀 Features
  - Multiple role authentications: Admin, Teacher, Student
  - Broadcast features for all/admins/teacher/student/class
  - View and Edit profiles
  - Student Attendance calculation and Marks allocation

## 📝 Note :- Use "admin123" as default admin access key, you can change it in models/admin.go

## 📄 License
This project is licensed under the MIT License - see the LICENSE file for details.
