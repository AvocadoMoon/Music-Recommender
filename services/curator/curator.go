package music_curator


import (
	"net/http"
	"github.com/gorilla/mux"
)

type Handler struct{
}


func NewHandler() *Handler{
	return &Handler{}
}


func (h *Handler) RegisterCuratorRoutes(router *mux.Router){
	router.HandleFunc("/login", h.login).Methods("POST")
	router.HandleFunc("/submitMusic", h.submitMusic).Methods("POST")
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request){
	// provide credentials to gain an authentication token
}

func (h *Handler) submitMusic(w http.ResponseWriter, r *http.Request){
	// submit music to be chosen to the data base.
}

