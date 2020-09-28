package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sangianpatrick/go-mux-recovery-demo/httphandler"
	"github.com/sangianpatrick/go-mux-recovery-demo/middleware"
	"github.com/sirupsen/logrus"
)

// default port will be set to "8080"
var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
}

func main() {
	logger := logrus.New()
	recoveryHandlerMiddleware := &middleware.MyRecoveryHandlerLogger{
		Logger: logger,
	}

	router := mux.NewRouter()

	httphandler.NewHTTPHandler(router)

	handler := handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(false),
		handlers.RecoveryLogger(recoveryHandlerMiddleware),
	)

	http.ListenAndServe(fmt.Sprintf(":%s", port), handler(router))
}
