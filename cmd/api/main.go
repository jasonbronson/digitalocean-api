package main

import (
	"docontroller/config"
	"docontroller/controller"
	"docontroller/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/digitalocean/godo"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	log.Println("Starting API")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.Config{
		GodoClient: godo.NewFromToken(os.Getenv("DO_API_KEY")),
	}
	r := mux.NewRouter()

	r.Use(utils.LoggingMiddleware)
	r.HandleFunc("/", HomeHandler)
	a := r.PathPrefix("/api/v1").Subrouter()
	a.HandleFunc("/droplets/all", func(w http.ResponseWriter, r *http.Request) {
		controller.GetDropletsHandler(w, r, &config)
	}).Methods("GET")
	a.HandleFunc("/droplets/create", func(w http.ResponseWriter, r *http.Request) {
		controller.CreateDropletHandler(w, r, &config)
	}).Methods("POST")
	a.HandleFunc("/droplets/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.DeleteDropletHandler(w, r, &config)
	}).Methods("DELETE")

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
