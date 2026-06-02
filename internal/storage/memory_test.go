package storage

import (
	"testing"

	"local-event-bus/internal/models"
)

func TestMemoryStorageCreateAndGetByID(t *testing.T) {
	store := New()

	created := store.Create(models.Appointment{
		Client:  "Анна",
		Service: "hair",
		Status:  "waiting",
	})

	found, ok := store.GetByID(created.ID)
	if !ok {
		t.Fatal("expected appointment to be found")
	}

	if found.Client != "Анна" {
		t.Fatalf("expected client Анна, got %s", found.Client)
	}
}

func TestMemoryStorageUpdateStatus(t *testing.T) {
	store := New()

	created := store.Create(models.Appointment{
		Client:  "Мария",
		Service: "nails",
		Status:  "waiting",
	})

	ok := store.UpdateStatus(created.ID, "completed")
	if !ok {
		t.Fatal("expected status to be updated")
	}

	found, _ := store.GetByID(created.ID)

	if found.Status != "completed" {
		t.Fatalf("expected completed, got %s", found.Status)
	}
}
