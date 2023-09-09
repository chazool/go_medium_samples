package main

import (
	"fmt"
	"modulemanager/configs"
	"modulemanager/dto"
	"modulemanager/repo"
)

func init() {
	// init db connection
	configs.InitDBConnection()
}

func main() {
	users := []dto.User{
		{
			UserName: "user1",
			Password: "test1",
		},
		{
			UserName: "user2",
			Password: "test2",
		},
	}

	repo := repo.NewUserRepo()

	// save users
	err := repo.Save(users...)
	if err != nil {
		return
	}

	// get all users
	users, err = repo.FindAll()
	if err != nil {
		return
	}

	// print users
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

}
