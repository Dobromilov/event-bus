package service

import (
	"testing"

	"local-event-bus/internal/storage"
)

func TestCreateAppointment(t *testing.T) {
	store := storage.New()
	bus := New(store)

	appointment := bus.CreateAppointment("Анна", "hair")

	if appointment.ID != 1 {
		t.Fatalf("expected id 1, got %d", appointment.ID)
	}

	if appointment.Client != "Анна" {
		t.Fatalf("expected client Анна, got %s", appointment.Client)
	}

	if appointment.Service != "hair" {
		t.Fatalf("expected service hair, got %s", appointment.Service)
	}

	if appointment.Status != "waiting" {
		t.Fatalf("expected status waiting, got %s", appointment.Status)
	}
}

func TestListAppointments(t *testing.T) {
	store := storage.New()
	bus := New(store)

	bus.CreateAppointment("Анна", "hair")
	bus.CreateAppointment("Мария", "nails")

	appointments := bus.ListAppointments()

	if len(appointments) != 2 {
		t.Fatalf("expected 2 appointments, got %d", len(appointments))
	}
}

func TestCompleteAppointment(t *testing.T) {
	store := storage.New()
	bus := New(store)

	appointment := bus.CreateAppointment("Анна", "hair")

	ok := bus.CompleteAppointment(appointment.ID)
	if !ok {
		t.Fatal("expected appointment to be completed")
	}

	updated, found := bus.GetAppointment(appointment.ID)
	if !found {
		t.Fatal("expected appointment to be found")
	}

	if updated.Status != "completed" {
		t.Fatalf("expected status completed, got %s", updated.Status)
	}
}
