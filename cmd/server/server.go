package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ShauryaAg/tcms-go/internal/db"
	_ "github.com/ShauryaAg/tcms-go/internal/db/mongo"
	"github.com/ShauryaAg/tcms-go/internal/routes"
	"github.com/rs/cors"
)

func main() {

	router := routes.NewRouter(db.Repositories)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Accept", "Accept-Language", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	srv := &http.Server{
		Addr:         ":" + "8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      c.Handler(router),
	}

	log.Println("Starting Server....")
	log.Fatal(srv.ListenAndServe())
}
