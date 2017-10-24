package models

import (
	"github.com/go-xorm/xorm"

	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

func init() {
	var err error
	// https://yhirano55.github.io/2017/03/26/hello-gin/ こちら参考にsqlite3で
	engine, err = xorm.NewEngine("sqlite3", "./development.db")
	if err != nil {
		panic(err)
	}
}

type User struct {
	ID        int    `json:"id" xorm:"'id'"`
	FirstName string `json: "first_name" xorm: "'first_name'"`
	LastName  string `json: "last_name" xorm: "'last_name'"`
}

func NewUser(id int, first_name string, last_name string) User {
	return User{
		ID:        id,
		FirstName: first_name,
		LastName:  last_name,
	}
}

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (m UserRepository) GetByID(id int) *User {
	var user = User{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}

	return nil
}
