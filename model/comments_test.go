package model

import (
	"testing"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T, userId, postID int) Comment {

	var randomComment Comment
	randomComment.AuthorID = userId
	randomComment.Body = util.RandomString(100)

	err := randomComment.Post(testDB.db, postID)
	require.NoError(t, err)
	return randomComment
}
func TestComments(t *testing.T) {

	t.Run("Post Method", func(t *testing.T) {
		randomUser := createRandomUser(t)
		randomPost := createRandomPost(t, randomUser)
		randomComment := createRandomComment(t, randomUser.ID, randomPost.ID)

		require.NotZero(t, randomComment.ID)
		cleanTable(t)
	})
	t.Run("Delete Method", func(t *testing.T) {
		randomUser := createRandomUser(t)
		randomPost := createRandomPost(t, randomUser)
		randomComment := createRandomComment(t, randomUser.ID, randomPost.ID)
		err := randomComment.Delete(testDB.db)
		require.NoError(t, err)
		comments, _ := GetComments(testDB.db, int64(randomPost.ID))

		require.Nil(t, comments)

		cleanTable(t)
	})
	t.Run("GetComments Method", func(t *testing.T) {
		randomUser := createRandomUser(t)
		randomPost := createRandomPost(t, randomUser)
		randomComment := createRandomComment(t, randomUser.ID, randomPost.ID)
		comments, err := GetComments(testDB.db, int64(randomPost.ID))
		require.NoError(t, err)

		require.Equal(t, randomComment, comments[0])
		cleanTable(t)
	})
}
