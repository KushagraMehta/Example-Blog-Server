package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	DB     *pgxpool.Pool
	Router *mux.Router
}

func (server *Server) Initialize() {

	var err error

	DBURL := os.Getenv("DATABASE_URL")
	if server.DB, err = pgxpool.Connect(context.Background(), DBURL); err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Println("We are connected to the database")
	}
	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(port string) {

	fmt.Printf("Listening to port %s\n", port)
	srv := &http.Server{
		Handler:      server.Router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	if err := srv.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}

	os.Exit(0)
}
