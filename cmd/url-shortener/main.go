package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: ini logger: slog

	// TODO: init storage: sqlite3

	// TODO: init router: chi, chi render

	// TODO: run server
}
