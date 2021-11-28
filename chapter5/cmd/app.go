package main

import (
	"log"

	"github.com/PacktPublishing/Domain-Driven-Design-with-Go/chapter5/internal/hash"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"

	"github.com/PacktPublishing/Domain-Driven-Design-with-Go/chapter5/internal/user"
)

func main() {
	// create all third-party dependencies
	validate := validator.New()
	engine := gin.Default()

	// creates new hash module
	hashModule := &hash.Module{}
	hashModule.Configure("super-secret-value")

	// creates new user module
	userModule := &user.Module{}
	userModule.Configure("./database.sqlite", engine, validate, hashModule.GetHashService())

	// run server
	log.Fatalln(engine.Run())
}
