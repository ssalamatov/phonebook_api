package storage

import (
	"sync"

	"github.com/ssalamatov/phonebook_api/internal/worker"
)

type StorageImpl interface {
	Add()
	Get()
	Delete()
}

type Storage struct {
	S *sync.Map
}

func NewStorage() *Storage {
	return &Storage{S: new(sync.Map)}
}

func (s *Storage) Add(worker *worker.Worker) {
	id := worker.Id
	s.S.Store(id, worker)
}

func (s *Storage) Get(id int) (*worker.Worker, error) {
	val, ok := s.S.Load(id)
	if !ok {
		return nil, ErrNotFound
	}
	w, ok := val.(*worker.Worker)
	if !ok {
		return nil, ErrInvalidWorker
	}
	return w, nil
}

func (s *Storage) Delete(id int) error {
	_, ok := s.S.LoadAndDelete(id)
	if !ok {
		return ErrNotFound
	}
	return nil
}
