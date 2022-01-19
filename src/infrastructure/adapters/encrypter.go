package adapters

import "golang.org/x/crypto/bcrypt"

type EncrypterAdapter struct{}

func (encrypterAdapter *EncrypterAdapter) Hash(text string) (string, error) {
	const BCRYPT_COST = 10
	hash, err := bcrypt.GenerateFromPassword([]byte(text), BCRYPT_COST)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (encrypterAdapter *EncrypterAdapter) Compare(text string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	if err != nil {
		return false, nil
	}
	return true, nil
}

func NewEncrypterAdapter() (*EncrypterAdapter, error) {
	return &EncrypterAdapter{}, nil
}
