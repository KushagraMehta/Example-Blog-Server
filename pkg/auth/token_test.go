package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {

	t.Run("TokenValid Method with valid token", func(t *testing.T) {
		token, _ := CreateToken(123)
		_, err := TokenValid(token)
		require.NoError(t, err)
	})
	t.Run("TokenValid Method with unvalid token", func(t *testing.T) {
		token, _ := CreateToken(123)
		_, err := TokenValid(token + "sdfg")
		require.Error(t, err)
	})
	t.Run("ExtractTokenID Method", func(t *testing.T) {
		want := 123
		token, _ := CreateToken(want)
		got, err := ExtractTokenID(token)
		require.NoError(t, err)
		require.Equal(t, want, got)
	})
}
