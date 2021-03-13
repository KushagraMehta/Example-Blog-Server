package model

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func (t *Tag) validation(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if t.Summary == "" {
			return errors.New("summary is required")
		}
		if t.Title == "" {
			return errors.New("title is empty")
		}
	case "attach", "delete":
		if t.ID == 0 {
			return errors.New("tag id is required")
		}
	}
	return nil
}

// Create will add tag In Database
func (t *Tag) Create(db *pgxpool.Pool) error {
	if err := t.validation("create"); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx, "INSERT INTO TAGS(TITLE,SUMMARY) VALUES($1,$2);", t.Title, t.Summary); err != nil {
		return err
	}
	return nil
}

//  AttachMe will add tag to a Post
func (t *Tag) AttachMe(db *pgxpool.Pool, postID int) error {
	if err := t.validation("attach"); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx, "INSERT INTO POST_TAG(POST_ID,TAG_ID) VALUES($1,$2)", postID, t.ID); err != nil {
		return err
	}
	return nil
}

//  Delete will remove tag from a post
func (t *Tag) Delete(db *pgxpool.Pool, postID int) error {
	if err := t.validation("delete"); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx, "DELETE FROM POST_TAG WHERE POST_ID=$1 AND TAG_ID=$2", postID, t.ID); err != nil {
		return err
	}
	return nil
}

// GetTagsData bring Top tags with data by limit
func GetTagsData(db *pgxpool.Pool, limit int) ([]Tag, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var data []Tag
	if rows, err := db.Query(ctx, "SELECT ID,TITLE,SUMMARY,TOTAL_POST FROM TAGS ORDER BY TAGS.TOTAL_POST DESC LIMIT $1;", limit); err != nil {
		return []Tag{}, err
	} else {
		defer rows.Close()
		var tmp Tag
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.Title, &tmp.Summary, &tmp.TotalPost)
			data = append(data, tmp)
		}

		if rows.Err() != nil {
			fmt.Println(err)
		}
	}
	return data, nil
}

//GetTagPosts bring PostID's related to a tag
func GetTagPosts(db *pgxpool.Pool, tagID int, limit int) ([]int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var data []int
	if rows, err := db.Query(ctx, "SELECT POST_ID FROM POST_TAG LIMIT $1;", limit); err != nil {
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
