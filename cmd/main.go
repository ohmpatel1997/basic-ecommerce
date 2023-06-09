package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	httpInternal "github.com/ohmpatel1997/rundoo-task/pkg/http"
	"github.com/ohmpatel1997/rundoo-task/pkg/product"
	"github.com/ohmpatel1997/rundoo-task/pkg/storage"
)

func main() {
	port := os.Getenv("POSTGRES_PORT")
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	name := os.Getenv("POSTGRES_NAME")

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		pass,
		host,
		port,
		name,
	)

	m, err := migrate.New(
		"file://migrations",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.NewWithConfig(
		context.Background(),
		cfg,
	)
	if err != nil {
		log.Fatal(err)
	}

	handler := httpInternal.NewHandler(product.NewService(storage.NewService(pool)))
	handler.RegisterRoutes(router)

	// configure the server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	log.Println("pcc-server now running on http://localhost:9999")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("error starting the http server", err.Error())
		}
	}()

	// gracefully shutdown on interrupt
	<-stop
	log.Printf("shutting down ...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("error shutting down the server", err.Error())
	}
	log.Printf("successfully shutdown \n")
}
