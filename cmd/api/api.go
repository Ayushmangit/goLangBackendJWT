package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
	"github.com/Ayushmangit/goLangBackendJWT/internal/env"
	"github.com/Ayushmangit/goLangBackendJWT/internal/handler"
	myMiddleware "github.com/Ayushmangit/goLangBackendJWT/internal/middleware"
	"github.com/Ayushmangit/goLangBackendJWT/internal/repository"
	"github.com/Ayushmangit/goLangBackendJWT/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
	log.Printf("The Server is Listening on port : %s\n", srv.Addr)
	return srv.ListenAndServe()
}

func (app *application) mount() *chi.Mux {

	ctx := context.Background()
	r := chi.NewRouter()

	conn, err := pgxpool.New(ctx, env.GetString("DSN", "postgres://root:71883@localhost:5433/app?sslmode=disable"))
	if err != nil {
		log.Fatal(err)
	}
	queries := sqlc.New(conn)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		_, err := userService.Register(ctx, "test@example.com", "password123")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("user Created"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", authHandler.Register)
			r.Post("/login", authHandler.Login)
			r.With(myMiddleware.AuthMiddleware).Get("/me", authHandler.Me)
		})
	})
	return r
}
