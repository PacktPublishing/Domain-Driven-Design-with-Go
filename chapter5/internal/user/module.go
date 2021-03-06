package user

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"

	hashApplication "github.com/PacktPublishing/Domain-Driven-Design-with-Go/chapter5/internal/hash/application"
	"github.com/PacktPublishing/Domain-Driven-Design-with-Go/chapter5/internal/user/application"
	"github.com/PacktPublishing/Domain-Driven-Design-with-Go/chapter5/internal/user/infrastructure"
	"github.com/PacktPublishing/Domain-Driven-Design-with-Go/chapter5/internal/user/presentation"
)

// Module is a struct that defines all dependencies inside user module
type Module struct{}

// Configure setups all dependencies
func (m *Module) Configure(db *gorm.DB, engine *gin.Engine, validate *validator.Validate, service hashApplication.HashService) {
	repository := infrastructure.NewUserRepository(db)
	useCase := application.NewRegistrationUseCase(repository)
	controller := presentation.NewUserController(useCase, validate, service)
	engine.POST("/users", controller.Register)
}
