package main

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type User struct {
	Id    int64
	Name  string
	Email string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Email)
}

func main() {
	db := pg.Connect(&pg.Options{
		User:     "goadmin",
		Password: "denis_denis",
		Database: "gobase",
	})
	defer db.Close()

	db.CreateTable(&User{}, &orm.CreateTableOptions{})

	db.Insert(&User{
		Name:  "Denis",
		Email: "Denis6705@gmail.com",
	})

	var users []User
	db.Model(&users).Select()

	fmt.Println(users)

}
