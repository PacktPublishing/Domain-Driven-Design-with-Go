package hash

import "github.com/PacktPublishing/Domain-Driven-Design-with-Go/chapter5/internal/hash/application"

// Module is a struct that defines all dependencies inside hash module
type Module struct {
	hashService application.HashService
}

// Configure setups all dependencies
func (m *Module) Configure(secret string) {
	m.hashService = application.NewMd5HashService(secret)
}

// GetHashService returns an instance of application.HashService
func (m Module) GetHashService() application.HashService {
	return m.hashService
}
