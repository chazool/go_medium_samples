package repo

import "multidbmanager/configs"

type DBOperator struct {
	Con configs.DBConnection
}

func (db *DBOperator) Save(value interface{}) error {
	return db.Con.Statement.Create(value).Error
}

func (db *DBOperator) Find(dest interface{}) error {
	return db.Con.Statement.Find(dest).Error
}
