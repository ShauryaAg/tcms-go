package routes

import (
	"net/http"

	httptransport "github.com/ShauryaAg/tcms-go/internal/http"
	httpinterfaces "github.com/ShauryaAg/tcms-go/pkg/http/interfaces"
	"github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"
)

func testCaseRouter(repository interfaces.Repository) http.Handler {
	mux := NewChiRouter()
	handler := httptransport.TestCaseHandler{
		Handler: httptransport.Handler {
			Router:     mux,
			Repository: repository,
		},
	}

	mux.Route("/", func(r httpinterfaces.Router) {
		r.Post("/", handler.CreateTestCase)
		r.Get("/{ID}", handler.GetTestCase)
	})

	return mux
}
