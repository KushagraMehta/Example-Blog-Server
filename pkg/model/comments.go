package model

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func (c *Comment) validation(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if c.AuthorID == 0 {
			return errors.New("author id is required")
		}
		if c.Body == "" {
			return errors.New("body is empty")
		}
		return nil
	case "delete":
		if c.ID == 0 {
			return errors.New("comment id is require")
		}
		return nil
	}
	return nil
}

// Post will put comment on a post. REQUIRE: postID,AuthorID,Body
func (c *Comment) Post(db *pgxpool.Pool, postID int) error {
	if err := c.validation("create"); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx, "INSERT INTO comments(AUTHOR_ID,BODY) VALUES($1,$2) returning id,created_on,updated_on;", c.AuthorID, c.Body).Scan(&c.ID, &c.CreatedOn, &c.UpdatedOn); err != nil {
		return err
	}
	if _, err := db.Exec(ctx, "INSERT INTO POST_COMMENTS(COMMENT_ID,POST_ID) VALUES($1,$2);", c.ID, postID); err != nil {
		return err
	}
	return nil
}

// Delete a comment from a post. REQUIRE: CommentID
func (c *Comment) Delete(db *pgxpool.Pool) error {
	if err := c.validation("delete"); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx, "DELETE FROM COMMENTS WHERE ID=$1;", c.ID); err != nil {
		return err
	}
	return nil
}

// GetComments will return All the comments on a post,REQUIRE: PostID
func GetComments(db *pgxpool.Pool, postID int64) ([]Comment, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var comments []Comment

	if rows, err := db.Query(ctx, "SELECT COMMENTS.ID,COMMENTS.AUTHOR_ID,COMMENTS.CREATED_ON,COMMENTS.UPDATED_ON,COMMENTS.BODY FROM POST_COMMENTS LEFT OUTER JOIN COMMENTS ON COMMENTS.ID = POST_COMMENTS.COMMENT_ID WHERE POST_COMMENTS.POST_ID=$1;", postID); err != nil {
		return comments, err
	} else {
		defer rows.Close()

		var tmp Comment
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.AuthorID, &tmp.CreatedOn, &tmp.UpdatedOn, &tmp.Body)
			comments = append(comments, tmp)
		}

		if rows.Err() != nil {
			return comments, err
		}
	}
	return comments, nil
}
