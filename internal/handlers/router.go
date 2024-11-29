package handlers

import (
	"database/sql"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "music_library/docs"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/swagger/*", httpSwagger.WrapHandler)

	router.HandleFunc("/songs", GetSongsHandler(db)).Methods(http.MethodGet)
	router.HandleFunc("/songs", AddSongHandler(db)).Methods(http.MethodPost)

	router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	router.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))
	router.Handle("/", http.FileServer(http.Dir("./public")))

	return router
}
