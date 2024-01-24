//go:build unit

package valueobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	plainPassword := NewPlainPassword("Test.123")
	hashed, err := NewHashFromPlain(plainPassword)
	assert.NoError(t, err)

	t.Run("Test hash from plain password", func(t *testing.T) {
		t.Run("Get hashed password", func(t *testing.T) {
			assert.NotEmpty(t, hashed.GetHash())
		})

		t.Run("Check valid password", func(t *testing.T) {
			assert.True(t, hashed.IsPasswordValid(plainPassword))
		})

		t.Run("Check for invalid password", func(t *testing.T) {
			assert.False(t, hashed.IsPasswordValid(NewPlainPassword("InvalidPassword")))
		})
	})

	t.Run("Test from hashed password", func(t *testing.T) {
		fromHash := NewHashPassword(hashed.GetHash())

		t.Run("Test get hash password", func(t *testing.T) {
			assert.Equal(t, hashed.GetHash(), fromHash.GetHash())
		})

		t.Run("Check valid password", func(t *testing.T) {
			assert.True(t, fromHash.IsPasswordValid(plainPassword))
		})

		t.Run("Check from invalid password", func(t *testing.T) {
			assert.False(t, fromHash.IsPasswordValid(NewPlainPassword("InvalidPassword")))
		})
	})
}
