package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	var greeting string
	err := server.DB.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, fmt.Sprintf("We are Up and running..!!\n%s", greeting))
}
