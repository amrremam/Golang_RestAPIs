package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listen string) *APIServer {
	return &APIServer{
		listenAddr: listen,
	}
}

func (s *APIServer) run() {
	router := mux.NewRouter

	router.HandleFunc("/account", s.handleAccount())
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
