package controller

import (
	"net/http"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {

	responses.JSON(w, http.StatusOK, "We are Up & running")
}
