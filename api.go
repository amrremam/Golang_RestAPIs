package main


import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "Application/json")
	return json.NewDecoder(w).Encode(v)
}


type apiFunc func(w http.ResponseWriter, r *http.Request)

type ApiError struct {
	Error string
}


func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}


type APIServer struct {
	listenAddr string
}


func NewAPIServer(listen string) *APIServer {
	return &APIServer{
		listenAddr: listen,
	}
}


func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	log.Println("JSON API Server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}


func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.GetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.CreateAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.DeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}


func (s *APIServer) GetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}


func (s *APIServer) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}


func (s *APIServer) DeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}


func (s *APIServer) Transfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
