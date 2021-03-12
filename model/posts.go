package model

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func (p *Post) Init(author int, title string, summary string) {
	p.AuthorID = author
	p.Title = title
	p.Summary = summary
}
func (p *Post) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if p.AuthorID == 0 {
			return errors.New("required author id")
		}
		if p.Title == "" {
			return errors.New("required proper title")
		}
		return nil
	case "draft":
		if p.AuthorID == 0 {
			return errors.New("required author id")
		}
		if p.ID == 0 {
			return errors.New("required post id")
		}
		return nil
	default:
		if p.ID == 0 {
			return errors.New("required post id")
		}
		if p.AuthorID == 0 {
			return errors.New("required author id")
		}
		if p.Title == "" {
			return errors.New("required proper title")
		}
		return nil
	}
}

//CreatePost will create a draft of post on the database
func (p *Post) Create(db *pgxpool.Pool) error {
	if err := p.Validate("create"); err != nil {
		return err
	}
	var postID int
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if rows, err := db.Query(ctx, "INSERT INTO posts(AUTHOR_ID,TITLE,SUMMARY) VALUES($1,$2,$3) returning id;", p.AuthorID, p.Title, p.Summary); err != nil {
		return err
	} else {
		defer rows.Close()

		rows.Next()
		rows.Scan(&postID)

		if rows.Err() != nil {
			return err
		}
	}
	p.ID = postID
	return nil
}

// GetDraftPost get the Drafted Post from Database
func (p *Post) GetDraft(db *pgxpool.Pool) error {
	if err := p.Validate("draft"); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx, "SELECT TITLE,SUMMARY,BODY FROM posts WHERE ID=$1 AND AUTHOR_ID=$2;", p.ID, p.AuthorID).Scan(&p.Title, &p.Summary, &p.Body); err != nil {
		return err
	}
	return nil
}

// PatchDraftedPost Update's Drafted Post, Update the values of Post object before calling it
func (p *Post) PatchDrafted(db *pgxpool.Pool) error {
	if err := p.Validate("draft"); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx, "UPDATE POSTS SET SUMMARY=$1,BODY=$2,PUBLISHED=$3,UPDATED_ON=current_timestamp WHERE ID=$4 AND AUTHOR_ID=$5;", p.Summary, p.Body, p.Published, p.ID, p.AuthorID); err != nil {
		return err
	}
	return nil
}

// Get will update the post object with data stored in database REQUIRE: PostID
func (p *Post) Get(db *pgxpool.Pool) error {
	if p.ID == 0 {
		return errors.New("post id is required")
	}
	ctx1, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx1, "SELECT AUTHOR_ID,TITLE,SUMMARY,UPDATED_ON,LIKE_COUNT,VIEWS,BODY FROM POSTS WHERE ID=$1", p.ID).Scan(&p.AuthorID, &p.Title, &p.Summary, &p.UpdatedOn, &p.LikeCount, &p.Views, &p.Body); err != nil {
		return err
	}
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx2, "UPDATE POSTS SET VIEWS=VIEWS+1 WHERE ID=$1;", p.ID); err != nil {
		return err
	}
	return nil
}

// GetTop return ID's of top viewed Posts. length is the limit of post require. REQUIRE:length
func GetTop(db *pgxpool.Pool, length int) ([]int, error) {
	var data []int
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if rows, err := db.Query(ctx, " SELECT ID FROM POSTS ORDER BY posts.views DESC LIMIT $1;", length); err != nil {
		return []int{}, err
	} else {
		defer rows.Close()
		var tmp int
		for rows.Next() {
			rows.Scan(&tmp)
			data = append(data, tmp)
		}

		if rows.Err() != nil {
			fmt.Println(err)
		}
	}
	return data, nil
}
