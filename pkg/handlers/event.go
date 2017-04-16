package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nbari/violetear"
	"github.com/rafaeljesus/crony/pkg/models"
	"github.com/rafaeljesus/crony/pkg/repos"
	"github.com/rafaeljesus/crony/pkg/response"
	"github.com/rafaeljesus/crony/pkg/scheduler"
)

type EventsHandler struct {
	EventRepo repos.EventRepo
	Scheduler scheduler.Scheduler
}

func NewEventsHandler(r repos.EventRepo, s scheduler.Scheduler) *EventsHandler {
	return &EventsHandler{
		EventRepo: r,
		Scheduler: s,
	}
}

func (h *EventsHandler) EventsIndex(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	expression := r.URL.Query().Get("expression")
	query := models.NewQuery(status, expression)

	events, err := h.EventRepo.Search(query)
	if err != nil {
		response.JSON(w, http.StatusPreconditionFailed, err)
		return
	}

	response.JSON(w, http.StatusOK, events)
}

func (h *EventsHandler) EventsCreate(w http.ResponseWriter, r *http.Request) {
	event := new(models.Event)
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}

	if err := h.EventRepo.Create(event); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.Scheduler.Create(event); err != nil {
		response.JSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, event)
}

func (h *EventsHandler) EventsShow(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}

	event, err := h.EventRepo.FindById(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, err)
		return
	}

	response.JSON(w, http.StatusOK, event)
}

func (h *EventsHandler) EventsUpdate(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].(string))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}

	event, err := h.EventRepo.FindById(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, err)
		return
	}

	e := new(models.Event)
	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		response.JSON(w, http.StatusNotFound, err)
		return
	}

	event.Status = e.Status
	event.Expression = e.Expression
	event.Url = e.Url
	event.Retries = e.Retries
	event.Timeout = e.Timeout

	if err := h.EventRepo.Update(event); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.Scheduler.Update(event); err != nil {
		response.JSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, event)
}

func (h *EventsHandler) EventsDelete(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	id, err := strconv.Atoi(params[":id"].([]string)[0])
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}

	event, err := h.EventRepo.FindById(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, err)
		return
	}

	if err := h.EventRepo.Delete(event); err != nil {
		response.JSON(w, http.StatusNotFound, err)
		return
	}

	if err := h.Scheduler.Delete(event.Id); err != nil {
		response.JSON(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
