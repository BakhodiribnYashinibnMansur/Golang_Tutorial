package storage

import (
	"article/models"
)

type Storage interface {
	Article() ArticleRepository
	Author() AuthorRepository
}

type ArticleRepository interface {
	Create(entity models.Article) (err error)
	GetList(query models.Query) (resp []models.Article, err error)
	GetByID(ID int) (resp models.Article, err error)
	Update(entity models.Article) (effectedRowsNum int64, err error)
	Delete(ID int) (effectedRowsNum int64, err error)
	Search(query models.Query) (resp []models.Article, err error)
}

type AuthorRepository interface {
	Create(entity models.Person) (err error)
	GetList(entity models.Query) (resp []models.Person, err error)
	GetById(ID int) (resp models.Person, err error)
	Update(entity models.Person) (effectedRowsNum int64, err error)
	Delete(ID int) (effectedRowsNum int64, err error)
	Search(query models.Query) (resp []models.Person, err error)
}
