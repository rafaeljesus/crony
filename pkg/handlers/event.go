package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nbari/violetear"
	"github.com/rafaeljesus/crony/pkg/models"
	"github.com/rafaeljesus/crony/pkg/render"
	"github.com/rafaeljesus/crony/pkg/repos"
	"github.com/rafaeljesus/crony/pkg/scheduler"
)

type EventsHandler struct {
	EventRepo repos.EventRepo
	Scheduler scheduler.Scheduler
}

func NewEventsHandler(r repos.EventRepo, s scheduler.Scheduler) *EventsHandler {
	return &EventsHandler{r, s}
}

func (h *EventsHandler) EventsIndex(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	expression := r.URL.Query().Get("expression")
	query := models.NewQuery(status, expression)

	events, err := h.EventRepo.Search(query)
	if err != nil {
		render.Response(w, http.StatusPreconditionFailed, err)
		return
	}

	render.JSON(w, http.StatusOK, events)
}

func (h *EventsHandler) EventsCreate(w http.ResponseWriter, r *http.Request) {
	event := models.NewEvent()
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		render.Response(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	if errors, valid := event.Validate(); !valid {
		render.Response(w, http.StatusBadRequest, errors)
		return
	}

	if err := h.EventRepo.Create(event); err != nil {
		render.Response(w, http.StatusUnprocessableEntity, "An error occurred during creating event")
		return
	}

	if err := h.Scheduler.Create(event); err != nil {
		render.Response(w, http.StatusInternalServerError, "An error occurred during scheduling event")
		return
	}

	render.JSON(w, http.StatusCreated, event)
}

func (h *EventsHandler) EventsShow(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		render.Response(w, http.StatusBadRequest, "Missing param :id")
		return
	}

	event, err := h.EventRepo.FindById(id)
	if err != nil {
		render.Response(w, http.StatusNotFound, "Event not found")
		return
	}

	render.JSON(w, http.StatusOK, event)
}

func (h *EventsHandler) EventsUpdate(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		render.Response(w, http.StatusBadRequest, "Missing param :id")
		return
	}

	event, err := h.EventRepo.FindById(id)
	if err != nil {
		render.Response(w, http.StatusNotFound, "Event not found")
		return
	}

	newEvent := models.NewEvent()
	if err := json.NewDecoder(r.Body).Decode(newEvent); err != nil {
		render.Response(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	if errors, valid := newEvent.Validate(); !valid {
		render.Response(w, http.StatusBadRequest, errors)
		return
	}

	event.SetAttributes(newEvent)
	if err := h.EventRepo.Update(event); err != nil {
		render.Response(w, http.StatusUnprocessableEntity, "An error occurred during updating event")
		return
	}

	if err := h.Scheduler.Update(event); err != nil {
		render.Response(w, http.StatusInternalServerError, "An error occurred during scheduling event")
		return
	}

	render.JSON(w, http.StatusOK, event)
}

func (h *EventsHandler) EventsDelete(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		render.Response(w, http.StatusBadRequest, "Missing param :id")
		return
	}

	event, err := h.EventRepo.FindById(id)
	if err != nil {
		render.Response(w, http.StatusNotFound, "Event not found")
		return
	}

	if err := h.EventRepo.Delete(event); err != nil {
		render.Response(w, http.StatusUnprocessableEntity, "An error occurred during deleting event")
		return
	}

	if err := h.Scheduler.Delete(event.Id); err != nil {
		render.Response(w, http.StatusInternalServerError, "An error occurred during deleting scheduled event")
		return
	}

	render.JSON(w, http.StatusNoContent, nil)
}
