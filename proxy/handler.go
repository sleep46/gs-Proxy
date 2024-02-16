package proxy

import (
	"fmt"
	"net/http"
	"strings"
)

type Handler struct {
	proxy *Proxy
}

func NewHandler(proxy *Proxy) *Handler {
	return &Handler{proxy: proxy}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/GET/")
	value, err := h.proxy.Get(key)
	if err != nil {
		errorMessage := fmt.Sprintf("Error retrieving value for key '%s' from proxy: %v\n", key, err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	response := fmt.Sprintf("Value for key '%s': %s\n", key, value)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
