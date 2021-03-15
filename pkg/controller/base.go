package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	DB     *pgxpool.Pool
	Router *http.ServeMux
}

func (server *Server) Initialize() {

	var err error

	DBURL := os.Getenv("DB_URL")
	if server.DB, err = pgxpool.Connect(context.Background(), DBURL); err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database")
	}
	server.Router = http.NewServeMux()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
