package server

import (
	"github.com/gorilla/mux"
)

// Endpoint defines an http endpoint to be queried from an external source
type Endpoint struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

// NewRouter creates the routes to listen
func NewRouter(h *Handler) *mux.Router {

	var endpoints = []Endpoint{
		{
			"Index",
			"GET",
			"/",
			h.Index,
		},
		{
			"TestGET",
			"GET",
			"/test",
			h.handleGET,
		},
		{
			"TestPOST",
			"POST",
			"/test",
			h.handlePOST,
		},
		{
			"TestPUT",
			"PUT",
			"/test",
			h.handlePUT,
		},
		{
			"TestDELETE",
			"DELETE",
			"/test",
			h.handleDELETE,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, endpoint := range endpoints {
		router.
			Methods(endpoint.Method).
			Path(endpoint.Pattern).
			Name(endpoint.Name).
			Handler(endpoint.HandlerFunc)
	}

	return router

}
