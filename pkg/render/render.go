package render

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

func Response(w http.ResponseWriter, code int, v interface{}) (err error) {
	wrapDefaults(w, code)

	var s string
	switch code {
	case http.StatusBadRequest:
		s = "invalid_request"
	case http.StatusUnprocessableEntity:
		s = "invalid_event"
	case http.StatusNotFound:
		s = "event_not_found"
	}

	r := &response{s, v}

	err = encode(w, r)

	return
}

func JSON(w http.ResponseWriter, code int, v interface{}) (err error) {
	wrapDefaults(w, code)

	if v == nil || code == http.StatusNoContent {
		return
	}

	err = encode(w, v)

	return
}

func wrapDefaults(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
}

func encode(w http.ResponseWriter, v interface{}) (err error) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	if err = enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}
