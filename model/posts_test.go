package model

import (
	"testing"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T, author User) Post {
	var testPost Post
	testPost.Init(author.ID, util.RandomString(10), util.RandomString(50))

	err := testPost.Create(testDB.db)
	require.NoError(t, err)

	return testPost
}
func createRandomPubishedPost(t *testing.T) Post {
	randomUser := createRandomUser(t)
	randomPost := createRandomPost(t, randomUser)
	randomPost.Body = util.RandomString(1000)
	randomPost.Published = true

	err := randomPost.PatchDrafted(testDB.db)
	require.NoError(t, err)
	return randomPost

}
func TestPostCreate(t *testing.T) {
	t.Run("Creating Proper Post Without any error", func(t *testing.T) {
		randomUser := createRandomUser(t)
		randomPost := createRandomPost(t, randomUser)
		require.NotZero(t, randomPost.ID)
		require.NotZero(t, randomPost.CreatedOn)
		require.NotZero(t, randomPost.UpdatedOn)
		require.Equal(t, randomPost.CreatedOn, randomPost.UpdatedOn)
		cleanTable(t)
	})
}
func TestPatchDrafted(t *testing.T) {
	t.Run("Patching Post with random data", func(t *testing.T) {
		randomUser := createRandomUser(t)
		randomPost := createRandomPost(t, randomUser)
		randomPost.Body = util.RandomString(100)
		randomPost.Published = true

		err := randomPost.PatchDrafted(testDB.db)
		require.NoError(t, err)

		randomPostClone := Post{AuthorID: randomUser.ID, ID: randomPost.ID}

		err = randomPostClone.GetDraft(testDB.db)
		require.NoError(t, err)

		require.Equal(t, randomPost.Title, randomPostClone.Title)
		require.Equal(t, randomPost.Summary, randomPostClone.Summary)
		require.Equal(t, randomPost.Summary, randomPostClone.Summary)
		require.Equal(t, randomPost.Body, randomPostClone.Body)
		require.Equal(t, randomPost.CreatedOn, randomPostClone.CreatedOn)
		require.Equal(t, randomPost.Published, randomPostClone.Published)

	})
	cleanTable(t)
}

func TestGetPostbyID(t *testing.T) {
	t.Run("Get Published Post", func(t *testing.T) {
		expectedPost := createRandomPubishedPost(t)
		actualPost, err := GetPostbyID(testDB.db, expectedPost.ID)
		require.NoError(t, err)
		require.Equal(t, expectedPost.ID, actualPost.ID)
		require.Equal(t, expectedPost.AuthorID, actualPost.AuthorID)
		require.Equal(t, expectedPost.Body, actualPost.Body)
		require.Equal(t, expectedPost.CreatedOn, actualPost.CreatedOn)
		require.Equal(t, expectedPost.LikeCount, actualPost.LikeCount)
		require.Equal(t, expectedPost.Published, actualPost.Published)
		require.Equal(t, expectedPost.Title, actualPost.Title)
		require.Equal(t, expectedPost.Summary, actualPost.Summary)
		cleanTable(t)
	})
	t.Run("Get unPublished/Non-Existing Post", func(t *testing.T) {
		actualPost, err := GetPostbyID(testDB.db, 4321)
		if assert.Error(t, err) {
			require.Equal(t, err.Error(), "no rows in result set")
		}
		require.Zero(t, actualPost.ID)
		require.Zero(t, actualPost.AuthorID)
		require.Zero(t, actualPost.Body)
		require.Zero(t, actualPost.CreatedOn)
		require.Zero(t, actualPost.Summary)
		require.Zero(t, actualPost.Title)
		cleanTable(t)
	})
}

//Sometime Can give Wrong Answer Need futher inspection. But can be solve by DB migartion
func TestGetTopPostIDs(t *testing.T) {
	randomPost1 := createRandomPubishedPost(t)
	randomPost2 := createRandomPubishedPost(t)

	viewOn1 := 100
	viewOn2 := 50

	for i := 0; i < viewOn1; i++ {
		GetPostbyID(testDB.db, randomPost1.ID)
	}
	for i := 0; i < viewOn2; i++ {
		GetPostbyID(testDB.db, randomPost2.ID)
	}
	topPostArray, err := GetTopPostIDs(testDB.db, 2)
	require.NoError(t, err)
	require.Equal(t, randomPost1.ID, topPostArray[0])
	require.Equal(t, randomPost2.ID, topPostArray[1])
	cleanTable(t)
}
