package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"twitter/middleware"
	"twitter/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Managements() {
	router := mux.NewRouter()
	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), handler))
}