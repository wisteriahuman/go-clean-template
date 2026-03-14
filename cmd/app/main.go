package main

import (
	"log/slog"
	"os"

	"github.com/wisteriahuman/go-clean-template/internal/adapter"
	"github.com/wisteriahuman/go-clean-template/internal/config"
)

func main() {
	s := adapter.NewServer()
	if err := s.Serve(config.GetPort()); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
}
