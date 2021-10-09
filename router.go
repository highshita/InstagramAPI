package main

import (
	"context"
	"errors"
	"net/http"
)

type key int

const (
	ParamsKey key = iota
)

func GetParam(ctx context.Context, name string) atring {
	params, _ := ctx.Value(ParamsKey).params

	for i := range Params {
		if params[i].key == name {
			return params[i].Value
		}
	}
	return ""
}

type Router struct {
	tree *tree
}

type route struct {
	methods     []string
	path        string
	handler     http.Handler
	middlewares middlewares
}

var (
	tmpRoute = &route{}
	ErrNotFound = errors.New("route not found")
	ErrMethodNotAllowed = errors.New("methods is not allowed")
)

func NewRouter() *Router {
	return &Router{
		tree: NewTree(),
	}
}

func (r *Router) Use(mws ...middleware) *Router {
	nm := NewMiddlewares(mws)
	tmpRoute.middlewares = nm
	return r
}

func (r *Router) Methods(methods ...string) *Router {
	tmpRoute.methods = append(tmpRoute.methods, methods...)
	return r
}

func (r *Router) Handler(path string, handler http.Handler) {
	tmpRoute.handler = handler
	tmpRoute.path = path
	r.Handle()
}

func (r *Router) Handle() {
	r.tree.Insert(tmpRoute.methods, tmpRoute.path, tmpRoute.handler, tmpRoute.middlewares)
	tmpRoute = &route{}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	result, err := r.tree.Search(method, path)
	if err != nil {
		status := handleErr(err)
		w.WriteHeader(status)
		return
	}
	h := result.actions.handler
	if result.actions.middlewares != nil {
		h = result.actions.middlewares.then(result.actions.handler)
	}
	if result.params != nil {
		ctx := context.WithValue(req.Context(), ParamsKey, result.params)
		req = req.WithContext(ctx)
	}
	h.ServeHTTP(w, req)
}

func handleErr(err error) int {
	var status int
	switch err {
	case ErrMethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case ErrNotFound:
		status = http.StatusNotFound
	}
	return status
}
