package model

import (
	"testing"

	"github.com/KushagraMehta/Example-Blog-Server/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomTag(t *testing.T) Tag {
	var randomtag Tag
	randomtag.Title = util.RandomString(10)
	randomtag.Summary = util.RandomString(100)
	err := randomtag.Create(testDB.db)
	require.NoError(t, err)

	return randomtag
}

func findTag(n int, array []Tag) bool {
	for _, v := range array {
		if n == v.ID {
			return true
		}
	}
	return false
}

func TestTags(t *testing.T) {

	t.Run("Create Tag", func(t *testing.T) {
		randomtag := createRandomTag(t)
		require.NotZero(t, randomtag.ID)
		cleanTable(t)
	})
	t.Run("AttachMe Method", func(t *testing.T) {
		randomtag := createRandomTag(t)
		randomUser := createRandomUser(t)
		randomPost := createRandomPost(t, randomUser)
		err := randomtag.AttachMe(testDB.db, randomPost.ID)
		require.NoError(t, err)
		require.NotZero(t, randomtag.ID)
		cleanTable(t)
	})
	t.Run("GetTagsOfPost Method", func(t *testing.T) {
		randomUser := createRandomUser(t)
		randomPost := createRandomPost(t, randomUser)
		randomTag1 := createRandomTag(t)
		randomTag2 := createRandomTag(t)
		randomTag1.AttachMe(testDB.db, randomPost.ID)
		randomTag2.AttachMe(testDB.db, randomPost.ID)

		tags, err := GetTagsOfPost(testDB.db, randomPost.ID)
		require.NoError(t, err)

		require.True(t, find(randomTag1.ID, tags))
		require.True(t, find(randomTag1.ID, tags))

		cleanTable(t)
	})
	t.Run("Delete Method", func(t *testing.T) {
		randomTag := createRandomTag(t)
		randomUser := createRandomUser(t)
		randomPost := createRandomPost(t, randomUser)

		randomTag.AttachMe(testDB.db, randomPost.ID)

		err := DeleteTags(testDB.db, randomTag.ID, randomPost.ID)
		require.NoError(t, err)
		tags, _ := GetTagsOfPost(testDB.db, randomPost.ID)
		require.Nil(t, tags)
		cleanTable(t)
	})
	t.Run("Get Tag Method", func(t *testing.T) {

		var tmp Tag
		var randomTags []Tag
		var storeDataOfFirstTag []int
		totalTags := 10
		totalPost := 5

		for i := 0; i < totalTags; i++ {
			tmp = createRandomTag(t)
			for j := 0; j < totalPost; j++ {
				randomUser := createRandomUser(t)
				randomPost := createRandomPost(t, randomUser)
				tmp.AttachMe(testDB.db, randomPost.ID)

				if i == 0 {
					storeDataOfFirstTag = append(storeDataOfFirstTag, randomPost.ID)
				}
			}
			randomTags = append(randomTags, tmp)
		}

		t.Run("GetTopTags Test", func(t *testing.T) {
			tagsGot, err := GetTopTags(testDB.db, 10)
			require.NoError(t, err)
			for _, v := range tagsGot {
				require.True(t, findTag(v.ID, randomTags))
			}
			cleanTable(t)
		})
		t.Run("GetPostsOfTag Test", func(t *testing.T) {
			postID, err := GetPostsOfTag(testDB.db, randomTags[0].ID, 10)
			require.NoError(t, err)

			for _, v := range postID {
				require.True(t, find(v, storeDataOfFirstTag))
			}
			cleanTable(t)
		})
	})
	t.Run("GetTagData Method", func(t *testing.T) {
		randomTag := createRandomTag(t)

		randomTagClone, err := GetTagData(testDB.db, randomTag.ID)

		require.NoError(t, err)
		require.Equal(t, randomTag, randomTagClone)
		cleanTable(t)
	})
}
