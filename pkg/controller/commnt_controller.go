package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/auth"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/model"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseInt(vars["postid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	comments, err := model.GetComments(server.DB, postID)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, comments)
}
func (server *Server) PostComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	postID, err := strconv.ParseInt(vars["postid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var comment model.Comment
	json.Unmarshal(body, &comment)
	comment.AuthorID, err = auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = comment.Post(server.DB, int(postID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, comment)
}
func (server *Server) DeleteComment(w http.ResponseWriter, r *http.Request) {

	var comment model.Comment
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	json.Unmarshal(body, &comment)
	err = comment.Delete(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Done")
}
