package routes

import (
	"net/http"

	"github.com/go-chi/chi"

	httptransport "github.com/ShauryaAg/tcms-go/internal/http"
	"github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"
)

func testStepRouter(repository interfaces.Repository) http.Handler {
	handler := httptransport.TestStepHandler{
		Repository: repository,
	}

	mux := chi.NewRouter()
	mux.Route("/", func(r chi.Router) {
		r.Post("/", handler.CreateTestStep)
	})

	return mux
}
