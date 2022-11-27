package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	Host              = "localhost"
	Port              = "8080"
	HelloWorldMessage = "Hello World!"
)

type Message struct {
	Message string `json:"message"`
}

// Future endpoints shouldn't be defined here I don't think
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	fmt.Println(fmt.Sprintf("Handling request %+v", r))
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	m := Message{HelloWorldMessage}
	json.NewEncoder(w).Encode(m)
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
