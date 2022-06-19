package service

import (
	"net/http"

	"github.com/ShauryaAg/tcms-go/internal/db"
)

func NewService(store db.Store, router http.Handler) *Service {
	return &Service{
		Store:  store,
		Router: &router,
	}
}

type Service struct {
	Store  db.Store
	Router *http.Handler
}
