package valueobject

import "golang.org/x/crypto/bcrypt"

type HashPassword interface {
	GetHash() string
	IsPasswordValid(passVO PlainPassword) bool
}

type hashPassword struct {
	hash string
}

func NewHashFromPlain(plainPassword PlainPassword) (HashPassword, error) {
	var err error
	hashPass := &hashPassword{}
	hashPass.hash, err = hashPass.buildHash(plainPassword.GetValue())
	return hashPass, err
}

func (hp *hashPassword) GetHash() string {
	return hp.hash
}

func (hp *hashPassword) IsPasswordValid(passVO PlainPassword) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hp.hash), []byte(passVO.GetValue()))
	return err == nil
}

func (hp *hashPassword) buildHash(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	return string(bytes), err
}
