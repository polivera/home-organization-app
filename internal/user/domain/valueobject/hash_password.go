package valueobject

import "golang.org/x/crypto/bcrypt"

type HashPassword interface {
	GetHash() string
	IsPasswordValid(passVO PlainPassword) bool
}

type hashPassword struct {
	hash string
	cost int
}

func NewHashFromPlain(plainPassword PlainPassword) (HashPassword, error) {
	var err error
	hashPass := &hashPassword{cost: 12}
	hashPass.hash, err = hashPass.buildHash(plainPassword.GetValue())
	return hashPass, err
}

func NewHashPassword(hashedPass string) HashPassword {
	return &hashPassword{hash: hashedPass}
}

func (hp *hashPassword) GetHash() string {
	return hp.hash
}

func (hp *hashPassword) IsPasswordValid(passVO PlainPassword) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hp.hash), []byte(passVO.GetValue()))
	return err == nil
}

func (hp *hashPassword) buildHash(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), hp.cost)
	return string(bytes), err
}
