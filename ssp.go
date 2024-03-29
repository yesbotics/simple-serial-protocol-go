package ssp

import (
	"fmt"
	"log/slog"
)

func Hello() {

	slog.SetLogLoggerLevel(slog.LevelDebug)

	slog.Debug("Hi")
	fmt.Println("Hi. Welcome!")
}
