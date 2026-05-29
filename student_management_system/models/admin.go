package models

type Admin struct {
	User
}

func CheckAdminAccess(inputKey string) bool {
	const AdminAccessKey = "admin123"

	if inputKey == AdminAccessKey {
		return true
	}

	return false
}
