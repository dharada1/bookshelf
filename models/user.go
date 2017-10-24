package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

// init ...
func init() {
	var err error
	// driveNameにDB、dataSourceNameにユーザ名：パスワード/ DB名をそれぞれ指定する。
	engine, err = xorm.NewEngine("mysql", "bookshelfman:password@/bookshelf")
	fmt.Print(engine)
	fmt.Print("I'm on xorm initialize function\n")
	if err != nil {
		panic(err)
	}
}

// User is ...
type User struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"name" xorm:"'nickname'"`
}

// NewUser ...
func NewUser(id int, username string) User {
	return User{
		ID:       id,
		Username: username,
	}
}

// UserRepository is
type UserRepository struct { // ここ記事だとtypoしてる
}

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// GetByID ...
func (m UserRepository) GetByID(id int) *User {
	var user = User{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}
	return nil
}
