package music_curator

import (
	"errors"
	"music-recommender/db"
	"music-recommender/types"
	"music-recommender/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{
	musicDB *db.MusicDB
}


func NewHandler(mdb *db.MusicDB) *Handler{
	return &Handler{mdb}
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
	var submitSong types.SubmitSong
	err := utils.DecodeJSONBody(w, r, &submitSong)
	var mal *utils.MalformedRequest
	if err != nil{
		errors.As(err, &mal)
		http.Error(w, mal.Msg, mal.Status)
		return
	}
	h.musicDB.InsertNewSong(&submitSong)
}

