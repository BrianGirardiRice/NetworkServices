package main

import (
	"log/slog"
	"net/http"

	"github.com/BrianGirardiRice/NetworkServices/taskHandler"
)

func main() {
	handler := taskHandler.NewServer()
	srv := http.Server{
		Addr:    ":5318",
		Handler: handler,
	}
	slog.Info("Listening on port 5318")
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server closed", "error", err)
	} else {
		slog.Info("Server closed", "error", err)
	}
}
