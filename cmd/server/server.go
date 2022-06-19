package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/cors"

	"github.com/ShauryaAg/tcms-go/internal/db"
	"github.com/ShauryaAg/tcms-go/internal/db/mongo"
	httptransport "github.com/ShauryaAg/tcms-go/internal/http"
)

func main() {

	ctx := context.Background()

	store := db.MongoStore{}
	client, err := store.Connect(ctx)
	defer client.Disconnect(ctx)
	if err != nil {
		log.Fatal("err", err)
	}

	repo := mongo.NewMongoTestCaseRepository(client.Database("tcms"))
	handler := httptransport.TestCaseHandler{
		Store: *repo,
	}

	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Post("/testcase", handler.CreateTestCase)
	})

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
		Handler:      c.Handler(r),
	}

	fmt.Printf("Starting Server....")
	log.Fatal(srv.ListenAndServe())
}
