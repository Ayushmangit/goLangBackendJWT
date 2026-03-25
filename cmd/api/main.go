package main

import (
	"github.com/Ayushmangit/goLangBackendJWT/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("PORT", ":3333"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()
	app.run(mux)
}
