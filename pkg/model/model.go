package model

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// User is the model present in the database
type User struct {
	ID             int       `json:"id,omitempty"`
	UserName       string    `json:"username,omitempty"`
	Email          string    `json:"email,omitempty"`
	PasswordHashed string    `json:"-"`
	CreatedOn      time.Time `json:"created_on,omitempty"`
	UpdatedOn      time.Time `json:"updated_on,omitempty"`
	LastLogin      time.Time `json:"last_login,omitempty"`
}

type UserInterface interface {
	GetLikedPost(db *pgxpool.Pool) ([]int, error)
	Init(username, email, password string) error
	Login(db *pgxpool.Pool) (int, error)
	PatchLike(db *pgxpool.Pool, postID int) error
	PutNewPassword(db *pgxpool.Pool, newPassword string) error
	SignUp(db *pgxpool.Pool) (int, error)
}

// Post is the model present in the database
type Post struct {
	ID        int       `json:"id"`
	AuthorID  int       `json:"author_id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary,omitempty"`
	Body      string    `json:"body,omitempty"`
	Published bool      `json:"published,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
	UpdatedOn time.Time `json:"updated_on,omitempty"`
	LikeCount int       `json:"like_count,omitempty"`
	Views     int       `json:"views,omitempty"`
}

type PostInterface interface {
	Create(db *pgxpool.Pool) error
	Get(db *pgxpool.Pool) error
	GetDraft(db *pgxpool.Pool) error
	Init(author int, title string, summary string)
	PatchDrafted(db *pgxpool.Pool) error
}

// Store Data regarding commments
type Comment struct {
	ID        int       `json:"id,omitempty"`
	AuthorID  int       `json:"author_id,omitempty"`
	Body      string    `json:"body,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
	UpdatedOn time.Time `json:"updated_on,omitempty"`
}

type Tag struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Summary   string `json:"summary,omitempty"`
	TotalPost int    `json:"total_post,omitempty"`
}
