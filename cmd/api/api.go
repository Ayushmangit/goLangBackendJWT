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
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Server running on %s\n", srv.Addr)
	return srv.ListenAndServe()
}

func (app *application) mount() *chi.Mux {
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, env.GetString("DSN", "postgres://root:71883@localhost:5433/app?sslmode=disable"))
	if err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(conn)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)

	studentRepo := repository.NewStudentRepository(queries)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Route("/api/v1", func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", authHandler.Register)
			r.Post("/login", authHandler.Login)
			r.With(myMiddleware.AuthMiddleware).Get("/me", authHandler.Me)
		})

		r.Route("/students", func(r chi.Router) {
			r.Use(myMiddleware.AuthMiddleware)
			r.Post("/", studentHandler.CreateStudent)
		})
	})

	return r
}
