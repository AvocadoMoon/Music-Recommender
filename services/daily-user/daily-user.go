package daily_user

import (
	"errors"
	"music-recommender/db"
	"music-recommender/types"
	"music-recommender/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Handler struct{
	musicDB *db.MusicDB
}


func NewHandler(mdb *db.MusicDB) *Handler{
	return &Handler{mdb}
}


func (h *Handler) RegisterAnonymousUserRoutes(router *mux.Router){
	router.HandleFunc("/todaysMusic", h.handleGettingTodaysMusic).Methods("GET")
	router.HandleFunc("/calendar", h.handleGettingCalendar).Methods("GET")
	router.HandleFunc("/todaysMusic", h.submitAVote).Methods("POST")
}


// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

func (h *Handler) submitAVote(w http.ResponseWriter, r *http.Request){
	var vote types.SubmitVotePayload
	err := utils.DecodeJSONBody(w, r, vote)
	if err != nil{
		var mr *utils.MalformedRequest
		if errors.As(err, &mr){
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Error().Msg(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	h.musicDB.UpdateTodaysRanking(vote)
	var todaysRanking *types.TodaysRankingPayload = h.musicDB.GetTodaysRanking()
	utils.WriteJSON(w, 200, todaysRanking)
}

func (h *Handler) handleGettingTodaysMusic(w http.ResponseWriter, r *http.Request){
	// get todays music from the DB and return the information
	todaysMusic := h.musicDB.GetTodaysMusic()
	utils.WriteJSON(w, 200, todaysMusic)
}

func (h *Handler) handleGettingCalendar(w http.ResponseWriter, r *http.Request){
	// get past music choices with their dates
	calendar := h.musicDB.GetCalendarsMusic()
	utils.WriteJSON(w, 200, calendar)
}


