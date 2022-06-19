package service

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ShauryaAg/tcms-go/internal/db"
	httptransport "github.com/ShauryaAg/tcms-go/internal/http"
	"github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"
)

func NewTestCaseService(store db.Store, router http.Handler, repo interfaces.Repository) *Service {
	handler := httptransport.TestCaseHandler{
		Repository: repo,
	}

	mux := router.(*chi.Mux)
	mux.Post("/testcase", handler.CreateTestCase)

	return NewService(store, router)
}
