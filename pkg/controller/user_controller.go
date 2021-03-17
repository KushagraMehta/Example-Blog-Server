package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/auth"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/model"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/responses"
	"github.com/KushagraMehta/Example-Blog-Server/pkg/util"
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

func (server *Server) FindUserByID(w http.ResponseWriter, r *http.Request)   {}
func (server *Server) PutNewPassword(w http.ResponseWriter, r *http.Request) {}
func (server *Server) UserLikedPost(w http.ResponseWriter, r *http.Request)  {}
func (server *Server) PatchUserLike(w http.ResponseWriter, r *http.Request)  {}
