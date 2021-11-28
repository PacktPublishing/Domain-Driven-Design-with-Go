package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/PacktPublishing/Domain-Driven-Design-with-Go/chapter5/internal/hash"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"

	"github.com/PacktPublishing/Domain-Driven-Design-with-Go/chapter5/internal/user"
)

func main() {
	// create all third-party dependencies
	validate := validator.New()
	engine := gin.Default()

	// initialize connection to the database
	db, err := gorm.Open(sqlite.Open("./database.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// creates new hash module
	hashModule := &hash.Module{}
	hashModule.Configure("super-secret-value")

	// creates new user module
	userModule := &user.Module{}
	userModule.Configure(db, engine, validate, hashModule.GetHashService())

	// run server
	log.Fatalln(engine.Run())
}
