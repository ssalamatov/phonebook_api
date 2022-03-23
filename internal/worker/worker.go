package worker

import (
	"time"

	"github.com/ssalamatov/phonebook_api/internal/user"
)

type Worker struct {
	User       *user.User `json:"user"`
	Id         int        `json:"id"`
	Department string     `json:"department"`
	JoinedAt   time.Time  `json:"joined_at"`
	IsActive   bool       `json:"is_active"`
}

func NewWorker(firstName string, lastName string, id int, department string, joinedAt time.Time) *Worker {
	return &Worker{
		User:       user.NewUser(firstName, lastName),
		Id:         id,
		Department: department,
		JoinedAt:   joinedAt}
}

func (w *Worker) SetActive(isActive bool) {
	w.IsActive = isActive
}
