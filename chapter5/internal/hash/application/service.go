package application

import (
	"crypto/md5"
	"fmt"
)

// HashService defines an application service inside hash module
type HashService interface {
	Do(text string) string
}

// md5HashService is an actual implementation of HashService
type md5HashService struct {
	secret string
}

// NewMd5HashService creates new md5HashService
func NewMd5HashService(secret string) HashService {
	return &md5HashService{
		secret: secret,
	}
}

// Do executes md5 hashing
func (m *md5HashService) Do(text string) string {
	full := fmt.Sprintf("%s#%s", m.secret, text)
	data := []byte(full)
	return fmt.Sprintf("%x", md5.Sum(data))
}
