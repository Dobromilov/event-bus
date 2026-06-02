package storage

import (
	"sync"

	"local-event-bus/internal/models"
)

type MemoryStorage struct {
	mu sync.RWMutex

	appointments []models.Appointment
	lastID       int64
}

func New() *MemoryStorage {
	return &MemoryStorage{
		appointments: make([]models.Appointment, 0),
	}
}

func (s *MemoryStorage) Create(a models.Appointment) models.Appointment {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++
	a.ID = s.lastID

	s.appointments = append(s.appointments, a)

	return a
}

func (s *MemoryStorage) List() []models.Appointment {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Appointment, len(s.appointments))
	copy(result, s.appointments)

	return result
}

func (s *MemoryStorage) GetByID(id int64) (*models.Appointment, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, a := range s.appointments {
		if a.ID == id {
			copyA := a
			return &copyA, true
		}
	}

	return nil, false
}

func (s *MemoryStorage) UpdateStatus(id int64, status string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.appointments {
		if s.appointments[i].ID == id {
			s.appointments[i].Status = status
			return true
		}
	}

	return false
}
