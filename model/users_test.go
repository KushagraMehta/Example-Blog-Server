package model

import (
	"testing"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	var testUser User
	err := testUser.Init(util.RandomOwner(), util.RandomEmail(), util.RandomString(6))
	require.NoError(t, err)

	_, err = testUser.SignUp(testDB.db)
	require.NoError(t, err)
	require.NotEmpty(t, testUser.ID)
	require.NotZero(t, testUser.CreatedOn)
	require.NotZero(t, testUser.UpdatedOn)
	require.NotZero(t, testUser.LastLogin)

	return testUser
}

func find(n int, array []int) bool {
	for _, v := range array {
		if n == v {
			return true
		}
	}
	return false
}

func TestLogin(t *testing.T) {

	t.Run("Proper login With-out any error", func(t *testing.T) {
		randomUser := createRandomUser(t)
		cloneUser := User{
			UserName:       randomUser.UserName,
			PasswordHashed: randomUser.PasswordHashed,
		}
		_, err := cloneUser.Login(testDB.db)
		require.NoError(t, err)
		require.Equal(t, randomUser.ID, cloneUser.ID)
		require.Equal(t, randomUser.Email, cloneUser.Email)
		require.Equal(t, randomUser.CreatedOn, cloneUser.CreatedOn)
		require.Equal(t, randomUser.UpdatedOn, cloneUser.UpdatedOn)
		cleanTable(t)
	})
	t.Run("Login With wrong Password", func(t *testing.T) {
		randomUser := createRandomUser(t)
		cloneUser := User{
			UserName:       randomUser.UserName,
			PasswordHashed: util.RandomString(6),
		}
		_, err := cloneUser.Login(testDB.db)
		if assert.Error(t, err) {
			require.Equal(t, err.Error(), "no rows in result set")
		}
		cleanTable(t)
	})
	t.Run("Login With wrong email", func(t *testing.T) {
		cloneUser := User{
			Email:          util.RandomEmail(),
			PasswordHashed: util.RandomString(6),
		}
		_, err := cloneUser.Login(testDB.db)
		if assert.Error(t, err) {
			require.Equal(t, err.Error(), "no rows in result set")
		}
		cleanTable(t)
	})

}
func TestPutNewPassword(t *testing.T) {

	t.Run("update password With username", func(t *testing.T) {
		randomUser := createRandomUser(t)
		cloneUser := User{
			ID:             randomUser.ID,
			UserName:       randomUser.UserName,
			PasswordHashed: randomUser.PasswordHashed,
		}
		err := cloneUser.PutNewPassword(testDB.db, util.RandomString(6))
		require.NoError(t, err)
		cleanTable(t)
	})
	t.Run("update password With Email", func(t *testing.T) {
		randomUser := createRandomUser(t)
		cloneUser := User{
			ID:             randomUser.ID,
			Email:          randomUser.Email,
			PasswordHashed: randomUser.PasswordHashed,
		}
		err := cloneUser.PutNewPassword(testDB.db, util.RandomString(6))
		require.NoError(t, err)
		cleanTable(t)
	})

}

func TestGetLikedPost(t *testing.T) {

	totalPost := 10
	randomUser := createRandomUser(t)
	randomPosts := make([]Post, totalPost)

	for i := 0; i < totalPost; i++ {
		randomPosts[i] = createRandomPubishedPost(t)
		err := randomUser.PatchLike(testDB.db, randomPosts[i].ID)
		require.NoError(t, err)
	}
	likedPost, err := randomUser.GetLikedPost(testDB.db)
	require.NoError(t, err)

	for i := 0; i < totalPost; i++ {
		require.True(t, find(randomPosts[i].ID, likedPost))
	}
	cleanTable(t)
}
func TestPatchLike(t *testing.T) {

	totalUser := 10
	randomPost := createRandomPubishedPost(t)
	randomUsers := make([]User, totalUser)

	for i := 0; i < totalUser; i++ {
		randomUsers[i] = createRandomUser(t)
		err := randomUsers[i].PatchLike(testDB.db, randomPost.ID)
		require.NoError(t, err)
	}
	err := randomUsers[5].PatchLike(testDB.db, randomPost.ID)
	require.NoError(t, err)
	updatedRandomPost, _ := GetPostbyID(testDB.db, randomPost.ID)
	require.Equal(t, 9, updatedRandomPost.LikeCount)
	cleanTable(t)
}

func TestFindUserByID(t *testing.T) {
	randomUser := createRandomUser(t)
	got, err := FindUserByID(testDB.db, randomUser.ID)
	require.NoError(t, err)

	require.Equal(t, randomUser.UserName, got.UserName)
	require.Equal(t, randomUser.LastLogin, got.LastLogin)
	require.Equal(t, randomUser.Email, got.Email)
	cleanTable(t)
}
