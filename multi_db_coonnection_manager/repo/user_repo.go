package repo

import (
	"multidbmanager/configs"
	"multidbmanager/dto"
)

type UserRepo interface {
	Save(...dto.User) error
	FindAll() ([]dto.User, error)
}

type UserRepoImpl struct {
	db DBOperator
}

func NewUserRepo(dbConIdentificationName string) UserRepo {
	return UserRepoImpl{
		db: DBOperator{
			Con: configs.GetDBConnection(dbConIdentificationName),
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
