package repo

import (
	"modulemanager/configs"
	"modulemanager/dto"
)

type UserRepo interface {
	Save(...dto.User) error
	FindAll() ([]dto.User, error)
}

type UserRepoImpl struct {
	db configs.DBOperator
}

func NewUserRepo() UserRepo {
	return UserRepoImpl{
		db: configs.DBOperator{
			Con: configs.GetDBConnection(),
		},
	}
}

func (repo UserRepoImpl) Save(user ...dto.User) error {
	err := repo.db.Save(user)
	if err != nil {
		// handle any special logic hdto
		return err
	}
	return nil
}

func (repo UserRepoImpl) FindAll() ([]dto.User, error) {
	var users []dto.User
	err := repo.db.Find(&users)
	if err != nil {
		// handle any special logic here
		return nil, err
	}

	return users, nil
}
