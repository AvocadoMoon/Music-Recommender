package daily_user

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Handler struct{
}


func NewHandler() *Handler{
	return &Handler{}
}


func (h *Handler) RegisterAnonymousUserRoutes(router *mux.Router){
	router.HandleFunc("/todaysMusic", h.handleGettingTodaysMusic).Methods("GET")
	router.HandleFunc("/calendar", h.handleGettingCalendar).Methods("GET")
}

func (h *Handler) handleGettingTodaysMusic(w http.ResponseWriter, r *http.Request){
	// get todays music from the DB and return the information
}

func (h *Handler) handleGettingCalendar(w http.ResponseWriter, r *http.Request){
	// get past music choices with their dates
}


