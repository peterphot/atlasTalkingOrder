package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

var RegisterRoutes = func(router *mux.Router) {
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/", fs)
	router.HandleFunc("/getTalkingOrder", GenerateTalkingOrder).Methods("GET")
}
