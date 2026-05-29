package models

type Teacher struct {
	User
	ClassAssigned bool
}

func (t *Teacher) AssignClassRequest(className string) (string, [2]string, [2]string) {
	message := t.Email + ", requested assign class: " + className

	from := [2]string{"teacher", t.Name}

	to := [2]string{"admin", ""}

	return message, to, from
}
