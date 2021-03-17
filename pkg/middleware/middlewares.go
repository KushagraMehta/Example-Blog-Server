package middlewares

import (
	"errors"
	"net/http"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/auth"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/responses"
)

func JSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
