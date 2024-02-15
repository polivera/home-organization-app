//go:build unit

package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	plainPassword := NewPlainPassword("Test.123")
	hashed, err := NewHashFromPlain(plainPassword)
	assert.NoError(t, err)

	t.Run("Test hash from plain password", func(t *testing.T) {
		t.Run("Get hashed password", func(t *testing.T) {
			assert.NotEmpty(t, hashed.Value())
		})

		t.Run("Check valid password", func(t *testing.T) {
			assert.True(t, hashed.MatchPlain(plainPassword))
		})

		t.Run("Check for invalid password", func(t *testing.T) {
			assert.False(t, hashed.MatchPlain(NewPlainPassword("InvalidPassword")))
		})
	})

	t.Run("Test from hashed password", func(t *testing.T) {
		fromHash := NewHashPassword(hashed.Value())

		t.Run("Test get hash password", func(t *testing.T) {
			assert.Equal(t, hashed.Value(), fromHash.Value())
		})

		t.Run("Check valid password", func(t *testing.T) {
			assert.True(t, fromHash.MatchPlain(plainPassword))
		})

		t.Run("Check from invalid password", func(t *testing.T) {
			assert.False(t, fromHash.MatchPlain(NewPlainPassword("InvalidPassword")))
		})
	})
}
