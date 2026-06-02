package service

import (
	"local-event-bus/internal/models"
	"local-event-bus/internal/storage"
	"time"
)

type EventBus struct {
	storage *storage.MemoryStorage
}

func New(storage *storage.MemoryStorage) *EventBus {
	return &EventBus{
		storage: storage,
	}
}

func (b *EventBus) CreateAppointment(
	client string,
	service string,
) models.Appointment {

	appointment := models.Appointment{
		Client:    client,
		Service:   service,
		Status:    "waiting",
		CreatedAt: time.Now(),
	}

	return b.storage.Create(appointment)
}

func (b *EventBus) ListAppointments() []models.Appointment {
	return b.storage.List()
}

func (b *EventBus) CompleteAppointment(id int64) bool {
	return b.storage.UpdateStatus(id, "completed")
}

func (b *EventBus) GetAppointment(id int64) (*models.Appointment, bool) {
	return b.storage.GetByID(id)
}
