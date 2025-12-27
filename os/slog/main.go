package main

import (
	"log/slog"
	"os"
)

func main() {
	//	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)) (isso é um log default, pra aumentar precisa adicionar umas coisas - T
	// Tipo pra colocar um debuger)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	data := slog.With(slog.String("response", "value ok"))
	logger.Info("Hello world", slog.String("key1", "value1"), slog.String("key2", "value2"), slog.Int("key3", 5), slog.Group("request", slog.String("response", "value ok")))
	slog.Debug("Debug message")
	data.Info("message")
}
