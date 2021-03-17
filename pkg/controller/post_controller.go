package controller

import "net/http"

func (server *Server) PostCreate(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) GetPostbyID(w http.ResponseWriter, r *http.Request)   {}
func (server *Server) GetTopPostIDs(w http.ResponseWriter, r *http.Request) {}
func (server *Server) GetDraft(w http.ResponseWriter, r *http.Request)      {}
func (server *Server) PatchDrafted(w http.ResponseWriter, r *http.Request)  {}
func (server *Server) GetTagsOfPost(w http.ResponseWriter, r *http.Request) {}
func (server *Server) GetPostsOfTag(w http.ResponseWriter, r *http.Request) {}
