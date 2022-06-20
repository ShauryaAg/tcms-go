package routes

import (
	"net/http"

	"github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"
	"github.com/go-chi/chi"
)

func NewRouter(repositories map[string]interfaces.Repository) http.Handler {
	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/testcase", testCaseRouter(repositories["test_cases"]))
		r.Mount("/teststep", testStepRouter(repositories["test_steps"]))
	})
	return router
}
