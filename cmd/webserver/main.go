package main

import (
	"log/slog"

	"github.com/wellls/api-example-golang/config/logger"
)

func main() {
	logger.InitLogger()

	slog.Info("starting api")
}
