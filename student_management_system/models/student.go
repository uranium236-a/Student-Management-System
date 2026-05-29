package models

type Student struct {
	User
	Rollno     int
	Marks      float32
	Attendance Attendance
}

type Attendance struct {
	AverageAttendance float32
	AttendanceList    []bool
}
