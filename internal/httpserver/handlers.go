package httpserver

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"local-event-bus/internal/service"
	"local-event-bus/internal/storage"
)

type AppointmentHandler struct {
	bus *service.EventBus
}

func NewAppointmentHandler(bus *service.EventBus) *AppointmentHandler {
	return &AppointmentHandler{
		bus: bus,
	}
}

func Register(mux *http.ServeMux) {
	store := storage.New()

	bus := service.New(store)

	handler := NewAppointmentHandler(bus)

	mux.HandleFunc("POST /appointments", handler.Create)
	mux.HandleFunc("GET /appointments", handler.List)
	mux.HandleFunc("GET /appointments/{id}", handler.GetByID)
	mux.HandleFunc("POST /appointments/{id}/complete", handler.Complete)
}

type CreateAppointmentRequest struct {
	Client  string `json:"client"`
	Service string `json:"service"`
}

// Create godoc
//
//	@Summary	Create appointment
//	@Tags		appointments
//	@Accept		json
//	@Produce	json
//	@Param		request	body		CreateAppointmentRequest	true	"appointment"
//	@Success	201		{object}	models.Appointment
//	@Router		/appointments [post]
func (h *AppointmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateAppointmentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	appointment := h.bus.CreateAppointment(
		req.Client,
		req.Service,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(appointment)
}

// List godoc
//
//	@Summary	List appointments
//	@Tags		appointments
//	@Produce	json
//	@Success	200	{array}	models.Appointment
//	@Router		/appointments [get]
func (h *AppointmentHandler) List(w http.ResponseWriter, r *http.Request) {
	appointments := h.bus.ListAppointments()

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(appointments)
}

// GetByID godoc
//
//	@Summary	Get appointment by id
//	@Tags		appointments
//	@Produce	json
//	@Param		id	path		int	true	"appointment id"
//	@Success	200	{object}	models.Appointment
//	@Router		/appointments/{id} [get]
func (h *AppointmentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/appointments/")

	id, err := strconv.ParseInt(path, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	appointment, ok := h.bus.GetAppointment(id)
	if !ok {
		http.Error(w, "appointment not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(appointment)
}

// Complete godoc
//
//	@Summary	Complete appointment
//	@Tags		appointments
//	@Param		id	path	int	true	"appointment id"
//	@Success	200
//	@Router		/appointments/{id}/complete [post]
func (h *AppointmentHandler) Complete(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/appointments/")
	path = strings.TrimSuffix(path, "/complete")

	id, err := strconv.ParseInt(path, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	ok := h.bus.CompleteAppointment(id)
	if !ok {
		http.Error(w, "appointment not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
