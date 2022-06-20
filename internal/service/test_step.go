package service

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ShauryaAg/tcms-go/internal/db"
	httptransport "github.com/ShauryaAg/tcms-go/internal/http"
	"github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"
)

func NewTestStepService(store db.Store, router http.Handler, repo interfaces.Repository) *Service {
	handler := httptransport.TestStepHandler{
		Repository: repo,
	}

	mux := router.(*chi.Mux)
	mux.Post("/teststep", handler.CreateTestStep)

	return NewService(store, router)
}
