package postgres

import (
	"article/models"
	"article/storage"

	"github.com/jmoiron/sqlx"
)

type authorRepo struct {
	db *sqlx.DB
}

func NewAuthorRepo(db *sqlx.DB) storage.AuthorRepository {
	return &authorRepo{db: db}
}

func (r *authorRepo) Create(entity models.Person) (err error) {
	return
}

func (r *authorRepo) GetList(entity models.Query) (resp []models.Person, err error) {
	return
}
func (r *authorRepo) Search(entity models.Query) (resp []models.Person, err error) {
	return
}

func (r *authorRepo) GetById(ID int) (resp models.Person, err error) {
	return
}

func (r *authorRepo) Update(entity models.Person) (effectedRowsNum int64, err error) {
	return
}

func (r *authorRepo) Delete(ID int) (effectedRowsNum int64, err error) {
	return
}
