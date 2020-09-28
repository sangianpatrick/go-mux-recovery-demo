package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPHandler is an entity of http handler.
type HTTPHandler struct {
	router *mux.Router
}

// NewHTTPHandler is a constructor of concrete HTTPHandler
func NewHTTPHandler(router *mux.Router) {
	hh := HTTPHandler{
		router: router,
	}

	hh.router.HandleFunc("/handle-panic", hh.HandlePanic).Methods(http.MethodGet)
}

// HandlePanic will simulate the panic event
func (hh HTTPHandler) HandlePanic(w http.ResponseWriter, r *http.Request) {
	// simulate the event when panic happens
	myData := make([]string, 2)
	fourth := myData[3]

	result := make(map[string]interface{})
	result["success"] = true
	result["data"] = fourth
	result["message"] = "The fourth data"

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
