//go:build unit

package valueobject_test

import (
	"testing"

	"github.com/polivera/home-organization-app/internal/user/domain/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestUsernameVO(t *testing.T) {
	t.Run("Check get value", func(t *testing.T) {
		username := valueobject.NewUsername("testUsername")
		assert.Equal(t, "testUsername", username.Value())
	})

	t.Run("Test is valid username", func(t *testing.T) {
		validUsernameList := []string{
			"testValid", "otherValidUser", "validusername", "valid45User",
		}
		invalidUsernameList := []string{
			"@#roberto", "#str@nges!mbols", "someth!ng", "very wrong",
		}

		for _, validUsername := range validUsernameList {
			uName := valueobject.NewUsername(validUsername)
			assert.True(t, uName.IsValid(), validUsername+" should be valid")
		}

		for _, invalidUsername := range invalidUsernameList {
			uName := valueobject.NewUsername(invalidUsername)
			assert.False(t, uName.IsValid(), invalidUsername+" should be invalid")
		}
	})
}
