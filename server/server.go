package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/sleep46/gs-Proxy/proxy"
)

type Server struct {
	port  int
	proxy *proxy.Proxy
}

func NewServer(port int, p *proxy.Proxy) *http.Server {
	s := &Server{
		port:  port,
		proxy: p,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleRoot)
	mux.HandleFunc("/GET/", s.handleGet)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	response := "Welcome to gs-Proxy <> ver_1.0\n\n"
	response += "Usage:\n"
	response += "GET /: Displays usage and configuration settings\n"
	response += "GET /GET/{key}: Returns the value associated with the provided key\n"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/GET/")
	value, err := s.proxy.Get(key)
	if err != nil {
		log.Printf("Error retrieving value for key '%s' from proxy: %v\n", key, err)
		errorMessage := fmt.Sprintf("Error retrieving value for key '%s' from proxy: %s\n", key, err.Error())
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	response := fmt.Sprintf("Value for key '%s': %s\n", key, value)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
