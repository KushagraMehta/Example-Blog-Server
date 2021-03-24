package model

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func (p *Post) Init(author int, title string, summary string) {
	p.AuthorID = author
	p.Title = title
	p.Summary = summary
}
func (p *Post) validate(action string) error {
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
	}
	return errors.New("wrong validaton option")
}

//Create will create a draft of post on the database. REQUIRE:AuthorID, Title, ?Summary
func (p *Post) Create(db *pgxpool.Pool) error {
	if err := p.validate("create"); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx, "INSERT INTO posts(AUTHOR_ID,TITLE,SUMMARY) VALUES($1,$2,$3) returning id,created_on,updated_on;", p.AuthorID, p.Title, p.Summary).Scan(&p.ID, &p.CreatedOn, &p.UpdatedOn); err != nil {
		return err
	}
	return nil
}

// GetDraftPost get the Drafted Post from Database. REQUIRE:AuthorID, ID
func (p *Post) GetDraft(db *pgxpool.Pool) error {
	if err := p.validate("draft"); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx, "SELECT TITLE,SUMMARY,published,created_on,updated_on,BODY FROM posts WHERE ID=$1 AND AUTHOR_ID=$2;", p.ID, p.AuthorID).Scan(&p.Title, &p.Summary, &p.Published, &p.CreatedOn, &p.UpdatedOn, &p.Body); err != nil {
		return err
	}
	return nil
}

// PatchDrafted Update's Drafted Post, Update the values of Post object before calling it. REQUIRE: AuthorID, PostID
func (p *Post) PatchDrafted(db *pgxpool.Pool) error {
	if err := p.validate("draft"); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx, "UPDATE POSTS SET SUMMARY=$1,BODY=$2,PUBLISHED=$3,UPDATED_ON=current_timestamp WHERE ID=$4 AND AUTHOR_ID=$5;", p.Summary, p.Body, p.Published, p.ID, p.AuthorID); err != nil {
		return err
	}
	return nil
}

// Get update the post object with published post stored in database REQUIRE: PostID
func GetPostbyID(db *pgxpool.Pool, postID int) (Post, error) {

	returnPost := Post{ID: postID}
	if postID == 0 {
		return Post{}, errors.New("post id is required")
	}

	ctx1, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx1, "SELECT AUTHOR_ID,TITLE,SUMMARY,created_on,UPDATED_ON,LIKE_COUNT,VIEWS,BODY,published FROM POSTS WHERE ID=$1 AND PUBLISHED='true';", returnPost.ID).Scan(&returnPost.AuthorID, &returnPost.Title, &returnPost.Summary, &returnPost.CreatedOn, &returnPost.UpdatedOn, &returnPost.LikeCount, &returnPost.Views, &returnPost.Body, &returnPost.Published); err != nil {
		return Post{}, err
	}
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := db.Exec(ctx2, "UPDATE POSTS SET VIEWS=VIEWS+1 WHERE ID=$1;", returnPost.ID); err != nil {
		return Post{}, err
	}
	return returnPost, nil
}

// GetTop return ID's of top viewed Posts. limit is the max number of post require. REQUIRE:limit
func GetTopPostIDs(db *pgxpool.Pool, limit int) ([]int, error) {
	var data []int
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if rows, err := db.Query(ctx, " SELECT ID FROM POSTS ORDER BY posts.views DESC LIMIT $1;", limit); err != nil {
		return []int{}, err
	} else {
		defer rows.Close()
		var tmp int
		for rows.Next() {
			rows.Scan(&tmp)
			data = append(data, tmp)
		}

		if rows.Err() != nil {
			return []int{}, err
		}
	}
	return data, nil
}
