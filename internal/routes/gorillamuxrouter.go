package routes

import (
	"net/http"

	httpinterfaces "github.com/ShauryaAg/tcms-go/pkg/http/interfaces"
	"github.com/gorilla/mux"
)

func NewGorillaMuxRouter() *GorillaMuxRouter {
	return &GorillaMuxRouter{
		mx: *mux.NewRouter(),
	}
}

type GorillaMuxRouter struct {
	mx mux.Router
}

func (ro *GorillaMuxRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ro.mx.ServeHTTP(w, r)
}

func (r *GorillaMuxRouter) URLParam(request http.Request, key string) string {
	return mux.Vars(&request)[key]
}

func (r *GorillaMuxRouter) Post(pattern string, handlerFn http.HandlerFunc) {
	r.mx.HandleFunc(pattern, handlerFn).Methods("POST")
}

func (r *GorillaMuxRouter) Get(pattern string, handlerFn http.HandlerFunc) {
	r.mx.HandleFunc(pattern, handlerFn).Methods("GET")
}

func (r *GorillaMuxRouter) Put(pattern string, handlerFn http.HandlerFunc) {
	r.mx.HandleFunc(pattern, handlerFn).Methods("PUT")
}

func (r *GorillaMuxRouter) Patch(pattern string, handlerFn http.HandlerFunc) {
	r.mx.HandleFunc(pattern, handlerFn).Methods("PATCH")
}

func (r *GorillaMuxRouter) Delete(pattern string, handlerFn http.HandlerFunc) {
	r.mx.HandleFunc(pattern, handlerFn).Methods("DELETE")
}

func (r *GorillaMuxRouter) Head(pattern string, handlerFn http.HandlerFunc) {
	r.mx.HandleFunc(pattern, handlerFn).Methods("HEAD")
}

func (r *GorillaMuxRouter) Options(pattern string, handlerFn http.HandlerFunc) {
	r.mx.HandleFunc(pattern, handlerFn).Methods("OPTION")
}

func (r *GorillaMuxRouter) Connect(pattern string, handlerFn http.HandlerFunc) {
	r.mx.HandleFunc(pattern, handlerFn).Methods("CONNECT")
}

func (r *GorillaMuxRouter) Trace(pattern string, handlerFn http.HandlerFunc) {
	r.mx.HandleFunc(pattern, handlerFn).Methods("TRACE")
}

func (r *GorillaMuxRouter) Route(pattern string, fn func(r httpinterfaces.Router)) httpinterfaces.Router {
	ro := NewGorillaMuxRouter()
	fn(ro)
	r.mx.PathPrefix(pattern).Handler(http.StripPrefix(pattern, ro))
	return ro
}

func (r *GorillaMuxRouter) Router(pattern string, fn func(r httpinterfaces.Router)) httpinterfaces.Router {
	ro := NewGorillaMuxRouter()
	fn(ro)
	r.mx.PathPrefix(pattern).Handler(http.StripPrefix(pattern, ro))
	return ro
}

func (r *GorillaMuxRouter) Use(middlewares ...func(http.Handler) http.Handler) {
	for _, wh := range middlewares {
		r.mx.Use(mux.MiddlewareFunc(wh))
	}
}

func (r *GorillaMuxRouter) Handle(pattern string, h http.Handler) {
	r.mx.PathPrefix(pattern).Handler(http.StripPrefix(pattern, h))
}

func (r *GorillaMuxRouter) HandleFunc(pattern string, h http.HandlerFunc) {
	r.mx.HandleFunc(pattern, h)
}

func (r *GorillaMuxRouter) Mount(pattern string, h http.Handler) {
	r.mx.PathPrefix(pattern).Handler(http.StripPrefix(pattern, h))
}

func (r *GorillaMuxRouter) Method(method, pattern string, h http.Handler) {
	r.mx.PathPrefix(pattern).Handler(http.StripPrefix(pattern, h)).Methods(method)
}

func (r *GorillaMuxRouter) MethodFunc(method, pattern string, h http.HandlerFunc) {
	r.mx.HandleFunc(pattern, h).Methods(method)
}

func (r *GorillaMuxRouter) NotFound(h http.HandlerFunc) {
	r.mx.NotFoundHandler = h
}

func (r *GorillaMuxRouter) MethodNotAllowed(h http.HandlerFunc) {
	r.mx.MethodNotAllowedHandler = h
}
