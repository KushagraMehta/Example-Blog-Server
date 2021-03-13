package model

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	db *pgxpool.Pool
}

// User is the model present in the database
type User struct {
	ID             uint32    `json:"id"`
	UserName       string    `json:"username"`
	Email          string    `json:"email"`
	PasswordHashed string    `json:"password"`
	CreatedOn      time.Time `json:"created_on"`
	UpdatedOn      time.Time `json:"updated_on"`
	LastLogin      time.Time `json:"last_login"`
}

type UserInterface interface {
	Validate() error
	Init(username, email, password string) error
	SignUp(db *pgxpool.Pool) error
	Login(db *pgxpool.Pool) (int64, error)
	PutNewPassword(db *pgxpool.Pool, newPassword string) error
	FindByID(db *pgxpool.Pool, uid int64) (*User, error)
	GetLikedPost(db *pgxpool.Pool) ([]int64, error)
	PatchLike(db *pgxpool.Pool, postID int64) error
}

// Post is the model present in the database
type Post struct {
	ID        int       `json:"id"`
	AuthorID  int       `json:"author_id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	Body      string    `json:"body"`
	Published bool      `json:"published"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	LikeCount uint      `json:"like_count"`
	Views     uint      `json:"views"`
}

type PostInterface interface {
	Init(author int, title string, summary string)
	Validate(action string) error
	Create(db *pgxpool.Pool) error
	GetDraft(db *pgxpool.Pool) error
	PatchDrafted(db *pgxpool.Pool) error
	Get(db *pgxpool.Pool) error
}

// Store Data regarding commments
type Comment struct {
	ID        int       `json:"id"`
	AuthorID  int       `json:"author_id"`
	Body      string    `json:"body"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type Tag struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	TotalPost int    `json:"total_post"`
}

// Connect Will Start the Connection to the PostgreSQL
func (tmp *DB) Connect() error {
	var err error
	databaseURL := os.Getenv("DB_URL")
	tmp.db, err = pgxpool.Connect(context.Background(), databaseURL)
	return err
}
func (tmp *DB) Close() {
	tmp.db.Close()
}
