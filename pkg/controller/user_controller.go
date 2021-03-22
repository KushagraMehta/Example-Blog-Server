package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/auth"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/model"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/responses"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/util"
	"github.com/gorilla/mux"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	var user model.User
	var token string
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tmpUser := model.User{}
	err = json.Unmarshal(body, &tmpUser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, util.FormatError(err))
		return
	}

	if err := user.Init(tmpUser.UserName, tmpUser.Email, tmpUser.Password); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, util.FormatError(err))
		return
	}
	_, err = user.Login(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, util.FormatError(err))
		return
	}
	token, err = auth.CreateToken(user.ID)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, util.FormatError(err))
		return
	}

	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var user model.User
	var token string
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tmpUser := model.User{}
	err = json.Unmarshal(body, &tmpUser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, util.FormatError(err))
		return
	}

	if err := user.Init(tmpUser.UserName, tmpUser.Email, tmpUser.Password); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, util.FormatError(err))
		return
	}
	_, err = user.SignUp(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, util.FormatError(err))
		return
	}
	token, err = auth.CreateToken(user.ID)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, util.FormatError(err))
		return
	}

	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) FindUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user, err := model.FindUserByID(server.DB, int(id))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}
func (server *Server) PutNewPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var user model.User
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
	user.ID = int(id)
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.PutNewPassword(server.DB, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Done")
}
func (server *Server) UserLikedPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := model.User{
		ID: int(id),
	}
	posts, err := user.GetLikedPost(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.PutNewPassword(server.DB, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}
func (server *Server) PatchUserLike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

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
	user := model.User{
		ID: int(id),
	}
	err = user.PatchLike(server.DB, int(postid))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Done")
}
