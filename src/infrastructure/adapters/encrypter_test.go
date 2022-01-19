package adapters

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypterAdapter_Hash(t *testing.T) {
	// arrange
	encrypter, _ := NewEncrypterAdapter()
	text := "any_text"
	regex := regexp.MustCompile(`^\$2[aby]?\$\d{1,2}\$[.\/A-Za-z0-9]{53}$`)
	// act
	hash, _ := encrypter.Hash(text)
	// assert
	assert.True(t, regex.Match([]byte(hash)))
}

func TestEncrypterAdapter_SuccessCompare(t *testing.T) {
	// arrange
	encrypter, _ := NewEncrypterAdapter()
	text := "any_text"
	hash, _ := encrypter.Hash(text)
	// act
	result, _ := encrypter.Compare(text, hash)
	// assert
	assert.True(t, result)
}

func TestEncrypterAdapter_FailCompare(t *testing.T) {
	// arrange
	encrypter, _ := NewEncrypterAdapter()
	text := "any_text"
	hash, _ := encrypter.Hash("other_text")
	// act
	result, _ := encrypter.Compare(text, hash)
	// assert
	assert.False(t, result)
}
