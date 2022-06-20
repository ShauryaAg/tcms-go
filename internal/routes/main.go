package routes

import (
	"net/http"

	"github.com/go-chi/chi/middleware"

	httpinterfaces "github.com/ShauryaAg/tcms-go/pkg/http/interfaces"
	"github.com/ShauryaAg/tcms-go/pkg/repository/interfaces"
)

func NewRouter(repositories map[string]interfaces.Repository) http.Handler {
	router := NewChiRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Route("/api/v1", func(r httpinterfaces.Router) {
		r.Mount("/testcase", testCaseRouter(repositories["test_cases"]))
		r.Mount("/teststep", testStepRouter(repositories["test_steps"]))
	})
	return router
}
