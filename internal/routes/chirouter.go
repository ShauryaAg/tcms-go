package routes

import (
	"net/http"

	"github.com/go-chi/chi"

	httpinterfaces "github.com/ShauryaAg/tcms-go/pkg/http/interfaces"
)

func NewChiRouter() ChiRouter {
	return ChiRouter{
		mux: chi.NewMux(),
	}
}

type ChiRouter struct {
	mux *chi.Mux
}

func (r ChiRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r ChiRouter) URLParam(request http.Request, key string) string {
	return chi.URLParam(&request, key)
}

func (r ChiRouter) Post(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Post(pattern, handlerFn)
}

func (r ChiRouter) Get(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Get(pattern, handlerFn)
}

func (r ChiRouter) Put(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Put(pattern, handlerFn)
}

func (r ChiRouter) Patch(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Patch(pattern, handlerFn)
}

func (r ChiRouter) Delete(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Delete(pattern, handlerFn)
}

func (r ChiRouter) Head(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Head(pattern, handlerFn)
}

func (r ChiRouter) Options(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Options(pattern, handlerFn)
}

func (r ChiRouter) Connect(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Connect(pattern, handlerFn)
}

func (r ChiRouter) Trace(pattern string, handlerFn http.HandlerFunc) {
	r.mux.Trace(pattern, handlerFn)
}

func (r ChiRouter) Route(pattern string, fn func(r httpinterfaces.Router)) httpinterfaces.Router {
	chifn := func(ro chi.Router) { fn(ChiRouter{mux: ro.(*chi.Mux)}) }

	return ChiRouter{
		mux: r.mux.Route(pattern, chifn).(*chi.Mux),
	}
}

func (r ChiRouter) Group(fn func(r httpinterfaces.Router)) httpinterfaces.Router {
	chifn := func(ro chi.Router) { fn(ChiRouter{mux: ro.(*chi.Mux)}) }

	return ChiRouter{
		mux: r.mux.Group(chifn).(*chi.Mux),
	}
}

func (r ChiRouter) Mount(pattern string, h http.Handler) {
	r.mux.Mount(pattern, h)
}

func (r ChiRouter) Handle(pattern string, h http.Handler) {
	r.mux.Handle(pattern, h)
}

func (r ChiRouter) HandleFunc(pattern string, h http.HandlerFunc) {
	r.mux.HandleFunc(pattern, h)
}

func (r ChiRouter) Method(method, pattern string, h http.Handler) {
	r.mux.Method(method, pattern, h)
}

func (r ChiRouter) MethodFunc(method, pattern string, h http.HandlerFunc) {
	r.mux.MethodFunc(method, pattern, h)
}

func (r ChiRouter) Use(middlewares ...func(http.Handler) http.Handler) {
	r.mux.Use(middlewares...)
}

func (r ChiRouter) With(middlewares ...func(http.Handler) http.Handler) httpinterfaces.Router {
	return ChiRouter{
		mux: r.mux.With(middlewares...).(*chi.Mux),
	}
}

func (r ChiRouter) NotFound(h http.HandlerFunc) {
	r.mux.NotFound(h)
}

func (r ChiRouter) MethodNotAllowed(h http.HandlerFunc) {
	r.mux.MethodNotAllowed(h)
}
