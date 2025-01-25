package http

import (
	"net/http"

	"github.com/babaYaga451/go-zomato/common/json"
)

func (h *OrderCommandHandler) writeJsonError(w http.ResponseWriter, status int, message string) {
	type envelope struct {
		Error string `json:"error"`
	}

	json.WriteJSON(w, status, &envelope{Error: message})
}

func (h *OrderCommandHandler) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err)

	h.writeJsonError(w, http.StatusInternalServerError, "internal server error")
}

func (h *OrderCommandHandler) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Warnw("bad request", "method", r.Method, "path", r.URL.Path, "error", err)

	h.writeJsonError(w, http.StatusBadRequest, "bad request error")
}

func (h *OrderCommandHandler) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Warnw("not found", "method", r.Method, "path", r.URL.Path, "error", err)

	h.writeJsonError(w, http.StatusNotFound, "not found error")
}
