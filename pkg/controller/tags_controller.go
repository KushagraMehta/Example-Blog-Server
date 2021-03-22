package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/model"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/responses"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/util"
	"github.com/gorilla/mux"
)

func (server *Server) TagCreate(w http.ResponseWriter, r *http.Request) {
	var tag model.Tag
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &tag)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = tag.Create(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, util.FormatError(err))
		return
	}

	responses.JSON(w, http.StatusOK, tag)
}
func (server *Server) GetTagData(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	tagID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tag, err := model.GetTagData(server.DB, int(tagID))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, tag)
}
func (server *Server) GetTopTags(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	limit, err := strconv.ParseInt(vars["limit"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tags, err := model.GetTopTags(server.DB, int(limit))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, tags)
}
func (server *Server) DeleteTags(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	postid, err := strconv.ParseInt(vars["postid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = model.DeleteTags(server.DB, int(id), int(postid))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, "Done")
}

func (server *Server) AttachMe(w http.ResponseWriter, r *http.Request) {

	var tag model.Tag
	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	postid, err := strconv.ParseInt(vars["postid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tag.ID = int(id)
	err = json.Unmarshal(body, &tag)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = tag.AttachMe(server.DB, int(postid))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Done")
}
