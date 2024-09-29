package main

import (
	"log"

	"github.com/ad-bak/chi/internal/env"
	"github.com/ad-bak/chi/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("API_ADDR", ":8080"),
	}
	store := store.NewPostgresStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
