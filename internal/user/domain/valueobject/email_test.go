//go:build unit

package valueobject

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmail(t *testing.T) {
	t.Run("Test email is valid", func(t *testing.T) {
		validEmails := [5]string{
			"test@test.local",
			"from@test.co.uk",
			"bull@frog.ll",
			"test@test",
			"bull@frog.llambada",
		}
		for _, validEmail := range validEmails {
			email := NewEmail(validEmail)
			assert.True(t, email.IsValid(), fmt.Sprintf("email %s should be valid", validEmails))
			assert.Equal(t, validEmail, email.Value())
		}
	})

	t.Run("Invalid email addresses", func(t *testing.T) {
		invalidEmails := [4]string{
			"from@test..uk",
			"notat.lambda.com",
			"joro\\m/ba@camapario.com",
			"bull@frog/something",
		}
		for _, invalidEmail := range invalidEmails {
			email := NewEmail(invalidEmail)
			assert.False(t, email.IsValid(), fmt.Sprintf("email %s should be invalid", invalidEmail))
			assert.Equal(t, invalidEmail, email.Value())
		}
	})
}
