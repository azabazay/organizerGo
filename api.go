package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	types "github.com/abai/organizer/types"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(listenAddr string) error {
	r := mux.NewRouter()

	r.HandleFunc("/", s.handleGetCatFact)
	r.HandleFunc("/get_user/{id}", s.handleGetUser)

	r.HandleFunc("/create_event", s.handleCreateEvent)
	r.HandleFunc("/delete_event/{id}", s.handleDeleteEvent)
	r.HandleFunc("/get_event/{id}", s.handleGetEvent)
	r.HandleFunc("/get_user_events/{id}", s.handleGetUserEvents)

	return http.ListenAndServe(listenAddr, r)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, fact)
}

func (s *ApiServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	userID, ok := vars["id"]
	if !ok {
		fmt.Println("user ID is missing in parameters")
	}

	user, err := s.svc.GetUser(userID)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (s *ApiServer) handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	event := types.TimeTableItem{}

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	defer func() {
		r.Body.Close()
	}()

	savedEvent, err := s.svc.CreateEvent(&event)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, savedEvent)
}

func (s *ApiServer) handleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	eventID, ok := vars["id"]
	if !ok {
		fmt.Println("event ID is missing in parameters")
	}

	user, err := s.svc.DeleteEvent(eventID)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (s *ApiServer) handleGetEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	eventID, ok := vars["id"]
	if !ok {
		fmt.Println("event ID is missing in parameters")
	}

	user, err := s.svc.GetEvent(eventID)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (s *ApiServer) handleGetUserEvents(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	userID, ok := vars["id"]
	if !ok {
		fmt.Println("user ID is missing in parameters")
	}

	userEvents, err := s.svc.GetUserEvents(userID)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, userEvents)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}
