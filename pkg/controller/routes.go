package controller

import middlewares "github.com/KushagraMehta/Example-Blog-Server/pkg/middleware"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.JSON(s.Home)).Methods("GET")

	// Login
	s.Router.HandleFunc("/login", middlewares.JSON(s.Login)).Methods("POST")

	// SignUp
	s.Router.HandleFunc("/signup", middlewares.JSON(s.SignUp)).Methods("POST")

	// User
	s.Router.HandleFunc("/users/{id:[0-9]+}", middlewares.JSON(s.FindUserByID)).Methods("GET")
	s.Router.HandleFunc("/users/{id:[0-9]+}", middlewares.Auth(middlewares.JSON(s.PutNewPassword))).Methods("PUT")
	s.Router.HandleFunc("/users/{id:[0-9]+}/like", middlewares.Auth(middlewares.JSON(s.UserLikedPost))).Methods("GET")
	s.Router.HandleFunc("/users/{id:[0-9]+}/like/{postid:[0-9]+}", middlewares.Auth(middlewares.JSON(s.PatchUserLike))).Methods("PATCH")

	// Post
	s.Router.HandleFunc("/posts", middlewares.Auth(middlewares.JSON(s.PostCreate))).Methods("POST")
	s.Router.HandleFunc("/posts/{id:[0-9]+}", middlewares.JSON(s.GetPostbyID)).Methods("GET")
	s.Router.HandleFunc("/posts/top/{limit:[0-9]+}", middlewares.JSON(s.GetTopPostIDs)).Methods("GET")
	s.Router.HandleFunc("/posts/draft/{id:[0-9]+}", middlewares.Auth(middlewares.JSON(s.GetDraft))).Methods("GET")
	s.Router.HandleFunc("/posts/draft/{id:[0-9]+}", middlewares.Auth(middlewares.JSON(s.PatchDrafted))).Methods("PATCH")
	s.Router.HandleFunc("/posts/{postId:[0-9]+}/tags", middlewares.JSON(s.GetTagsOfPost)).Methods("GET")
	s.Router.HandleFunc("/posts/tag/{id:[0-9]+}/{limit:[0-9]+}", middlewares.JSON(s.GetPostsOfTag)).Methods("GET")

	// Tags
	s.Router.HandleFunc("/tags", middlewares.Auth(middlewares.JSON(s.TagCreate))).Methods("POST")
	s.Router.HandleFunc("/tags/{id:[0-9]+}", middlewares.JSON(s.GetTagData)).Methods("GET")
	s.Router.HandleFunc("/tags/top/{limit:[0-9]+}", middlewares.JSON(s.GetTopTags)).Methods("GET")
	s.Router.HandleFunc("/tags/{postid:[0-9]+}/{id:[0-9]+}", middlewares.Auth(middlewares.JSON(s.DeleteTags))).Methods("DELETE")
	s.Router.HandleFunc("/tags/{id:[0-9]+}/add/{postid:[0-9]+}", middlewares.Auth(middlewares.JSON(s.AttachMe))).Methods("POST")

	// Comments
	s.Router.HandleFunc("/comments/{postid:[0-9]+}", middlewares.JSON(s.GetComments)).Methods("GET")
	s.Router.HandleFunc("/comments/{postid:[0-9]+}", middlewares.Auth(middlewares.JSON(s.PostComment))).Methods("POST")
	s.Router.HandleFunc("/comments/{postid:[0-9]+}/{id:[0-9]+}", middlewares.Auth(middlewares.JSON(s.DeleteComment))).Methods("DELTE")

}
