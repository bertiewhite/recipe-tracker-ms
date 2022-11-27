package handler

import (
	"fmt"
	"net/http"
)

const (
	Host              = "localhost"
	Port              = "8080"
	HelloWorldMessage = "Hello World!"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, HelloWorldMessage)
}

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct{}

func (h *Handler) AddRoute(path string, action func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, action)
}

func (h *Handler) Start() error {
	err := http.ListenAndServe(Host+":"+Port, nil)
	return err
}
