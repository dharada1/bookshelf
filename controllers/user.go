package controllers

import (
	"github.com/dharada1/bookshelf/models"
)

// User is
type User struct {
}

// New User
func NewUser() User {
	return User{}
}

// Get
func (c User) Get(n int) interface{} {
	repo := models.NewUserRepository()
	user := repo.GetByID(n)
	return user
}
