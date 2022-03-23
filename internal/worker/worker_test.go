package worker

import (
	"testing"
	"time"
)

func CheckValue(t *testing.T, actual interface{}, expected interface{}) {
	if expected != actual {
		t.Errorf("Expected \n%s. Got \n%s\n", expected, actual)
	}
}
func TestCreateNewWorker(t *testing.T) {
	name := "username"
	lastName := "lastname"
	id := 1
	department := "It"
	joined_at := time.Now().Add(time.Second * 5)
	w := NewWorker(name, lastName, id, department, joined_at)

	CheckValue(t, w.User.FirstName, name)
	CheckValue(t, w.User.LastName, lastName)
	CheckValue(t, w.Id, id)
	CheckValue(t, w.Department, department)
	CheckValue(t, w.JoinedAt, joined_at)

}
func TestSetIsActive(t *testing.T) {
	w := NewWorker("aaa", "bbb", 1, "It", time.Now())

	if w.IsActive {
		t.Errorf("Default state of worker is false. Got %v", w.IsActive)
	}

	w.SetActive(true)
	if !w.IsActive {
		t.Errorf("State of worker is true. Got %v", w.IsActive)
	}
}
