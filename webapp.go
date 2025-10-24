package main

import (
	"log"
	"time"
	"net/http"

	"github.com/gorilla/mux"
	_ "net/http/pprof"

	"webapp/handlers"
)

// Main
func main() {

	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", handlers.Home)
	router.HandleFunc("/home", handlers.Home)
	router.HandleFunc("/openwebuichat", handlers.OpenWebuiChat)

	// pprof thread /debug/pprof
	go http.ListenAndServe("localhost:6060", nil)

	// Serve
	log.Print("Serve Http on 8085")
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8085",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
