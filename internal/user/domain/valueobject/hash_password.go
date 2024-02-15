package valueobject

import (
	"github.com/polivera/home-organization-app/internal/common/valueobject"
	"golang.org/x/crypto/bcrypt"
)

type HashPassword interface {
	valueobject.ValueObject[string]
	MatchPlain(passVO PlainPassword) bool
}

type hashPassword struct {
	hash string
	cost int
}

func NewHashFromPlain(plainPassword PlainPassword) (HashPassword, error) {
	var err error
	hashPass := &hashPassword{cost: 12}
	hashPass.hash, err = hashPass.buildHash(plainPassword.Value())
	return hashPass, err
}

func NewHashPassword(hashedPass string) HashPassword {
	return hashPassword{hash: hashedPass}
}

func (hp hashPassword) Value() string {
	return hp.hash
}

func (hp hashPassword) IsValid() bool {
	return true
}

func (hp hashPassword) MatchPlain(passVO PlainPassword) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hp.hash), []byte(passVO.Value()))
	return err == nil
}

func (hp hashPassword) buildHash(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), hp.cost)
	return string(bytes), err
}
