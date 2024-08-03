package main

import (
	"auth/internal/app"
	"auth/internal/config"
)

func main() {
	cfg := config.MustLoad()
	app.New(cfg).Run()
}
