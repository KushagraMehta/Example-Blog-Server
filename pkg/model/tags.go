package model

import (
	"context"
	"errors"
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

// Create will add tag In Database. REQUIRE: Title, Summary
func (t *Tag) Create(db *pgxpool.Pool) error {
	if err := t.validation("create"); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx, "INSERT INTO TAGS(TITLE,SUMMARY) VALUES($1,$2) returning id;", t.Title, t.Summary).Scan(&t.ID); err != nil {
		return err
	}
	return nil
}

//  AttachMe will add tag to a Post. REQUIRE: TagID,PostID
func (t *Tag) AttachMe(db *pgxpool.Pool, postID int) error {
	if err := t.validation("attach"); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx, "INSERT INTO POST_TAGS(POST_ID,TAG_ID) VALUES($1,$2)", postID, t.ID); err != nil {
		return err
	}
	return nil
}

//  Delete will remove tag from a post. REQUIRE: TagID, PostID
func DeleteTags(db *pgxpool.Pool, tagID, postID int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := db.Exec(ctx, "DELETE FROM POST_TAGS WHERE POST_ID=$1 AND TAG_ID=$2", postID, tagID); err != nil {
		return err
	}
	return nil
}

// GetTopTags bring Top tags with data by limit, REQUIRE:Limit
func GetTopTags(db *pgxpool.Pool, limit int) ([]Tag, error) {

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
			return []Tag{}, err
		}
	}
	return data, nil
}

//GetPostsOfTag bring PostID's related to a tag
func GetPostsOfTag(db *pgxpool.Pool, tagID int, limit int) ([]int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var data []int
	if rows, err := db.Query(ctx, "SELECT POST_ID FROM POST_TAGS LIMIT $1;", limit); err != nil {
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

// GetTagsOfPost Will Return all the tagsID on a post REQUIRE: postID
func GetTagsOfPost(db *pgxpool.Pool, postID int) ([]int, error) {

	var data []int
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if rows, err := db.Query(ctx, "select tag_id from post_tags where post_id=$1;", postID); err != nil {
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

// // GetTagData Will return Tag object From Database REQUIRE: TagID
func GetTagData(db *pgxpool.Pool, tagID int) (Tag, error) {

	var data Tag
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.QueryRow(ctx, "select id,title,summary,total_post from tags where id=$1;", tagID).Scan(&data.ID, &data.Title, &data.Summary, &data.TotalPost); err != nil {
		return Tag{}, err
	}
	return data, nil
}
