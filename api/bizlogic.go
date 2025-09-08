package api

import (
	"Project/dataservice"
	"Project/model"
	"database/sql"
)

type IbizLogic interface {
	CreateBookLogic(book model.Book) error
}

type bizlogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *bizlogic {
	return &bizlogic{DB: db}
}

func (bl *bizlogic) CreateBookLogic(book model.Book) error {
	//validations
	if err := dataservice.CreateBook(bl.DB, book); err != nil {
		return err
	}
	return nil
}
