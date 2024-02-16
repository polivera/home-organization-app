//go:build unit

package valueobject_test

import (
	"github.com/polivera/home-organization-app/internal/user/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	plainPassword := valueobject.NewPlainPassword("Test.123")
	hashed, err := valueobject.NewHashFromPlain(plainPassword)
	assert.NoError(t, err)

	t.Run("Test hash from plain password", func(t *testing.T) {
		t.Run("Get hashed password", func(t *testing.T) {
			assert.NotEmpty(t, hashed.Value())
		})

		t.Run("Check valid password", func(t *testing.T) {
			assert.True(t, hashed.MatchPlain(plainPassword))
		})

		t.Run("Check for invalid password", func(t *testing.T) {
			assert.False(t, hashed.MatchPlain(valueobject.NewPlainPassword("InvalidPassword")))
		})
	})

	t.Run("Test from hashed password", func(t *testing.T) {
		fromHash := valueobject.NewHashPassword(hashed.Value())

		t.Run("Test get hash password", func(t *testing.T) {
			assert.Equal(t, hashed.Value(), fromHash.Value())
		})

		t.Run("Check valid password", func(t *testing.T) {
			assert.True(t, fromHash.MatchPlain(plainPassword))
		})

		t.Run("Check from invalid password", func(t *testing.T) {
			assert.False(t, fromHash.MatchPlain(valueobject.NewPlainPassword("InvalidPassword")))
		})
	})
}
