package api

import (
	"database/sql"
	"music-recommender/services/daily-user"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	db   *sql.DB // Pointer to SQL DB
	addr string
}

func CreateMainServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		db:   db,
		addr: addr,
	}
}

func (a APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	
	anonymous_user_handler := daily_user.NewHandler()
	anonymous_user_handler.RegisterAnonymousUserRoutes(subrouter)

	return http.ListenAndServe(a.addr, router)
}
