package routes

import (
	"net/http"

	httptransport "github.com/ShauryaAg/tcms-go/internal/http"
	httpinterfaces "github.com/ShauryaAg/tcms-go/pkg/http/interfaces"
	"github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"
)

func testStepRouter(repository interfaces.Repository) http.Handler {
	mux := NewChiRouter()
	handler := httptransport.TestStepHandler{
		Router:     mux,
		Repository: repository,
	}

	mux.Route("/", func(r httpinterfaces.Router) {
		r.Post("/", handler.CreateTestStep)
	})

	return mux
}
