package main

import (
	"context"
	"log"

	"github.com/Ayushmangit/goLangBackendJWT/internal/db/sqlc"
	"github.com/Ayushmangit/goLangBackendJWT/internal/env"
	"github.com/Ayushmangit/goLangBackendJWT/internal/seeder"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	dsn := env.GetString("DSN", "postgres://root:71883@localhost:5433/app?sslmode=disable")

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to the Database")

	queries := sqlc.New(conn)
	err = seeder.Seed(ctx, queries)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("seeding completed!")
}
