//go:build unit

package valueobject_test

import (
	"github.com/polivera/home-organization-app/internal/user/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlainPassword(t *testing.T) {
	t.Run("Check password is valid and get value", func(t *testing.T) {
		validPassList := [4]string{
			"Test.123",
			"SomethingSpecial34!",
			"formSTest664.",
			"Ts1!@$%^&*()-=_+,./<>?;':\"[]{}",
		}
		for _, pass := range validPassList {
			plainPass := valueobject.NewPlainPassword(pass)
			assert.True(t, plainPass.IsValid())
			assert.Equal(t, pass, plainPass.Value())
		}
	})
	t.Run("Check password is invalid", func(t *testing.T) {
		validPassList := [4]string{
			"Test0123",
			"somethingspecial34!",
			"FORMSTEST664",
			"T1!@$%^&*()-=_+,./<>?;':\"[]{}",
		}
		for _, pass := range validPassList {
			plainPass := valueobject.NewPlainPassword(pass)
			assert.False(t, plainPass.IsValid())
		}
	})
}
