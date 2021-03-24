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

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	var user model.User
	var token string
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	type password struct {
		Password string `json:"password"`
	}
	var pass password
	tmpUser := model.User{}
	if err = json.Unmarshal(body, &tmpUser); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(body, &pass); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Init(tmpUser.UserName, tmpUser.Email, pass.Password)
	if _, err = user.Login(server.DB); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if token, err = auth.CreateToken(user.ID); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var token string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tmpUser := model.User{}
	type password struct {
		Password string `json:"password"`
	}
	var pass password
	if err = json.Unmarshal(body, &tmpUser); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err = json.Unmarshal(body, &pass); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Init(tmpUser.UserName, tmpUser.Email, pass.Password)
	if _, err = user.SignUp(server.DB); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if token, err = auth.CreateToken(user.ID); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
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
