package storage

import (
	"article/models"
	"errors"
)

// type articleStorage map[int]models.Article
func NewArticleStorage() articleStorage {
	return articleStorage{
		data: make(map[int]models.Article),
	}
}

type articleStorage struct {
	data map[int]models.Article
}

//Errors varableName
var ErrorAlreadyExists = errors.New("already exists id")
var ErrorNotFound = errors.New("not found id ")
var ErrorNotUpdated = errors.New("not update because not found id")

//add new article
func (storage *articleStorage) Add(entity models.Article) error {
	if _, ok := storage.data[entity.ID]; ok {
		return ErrorAlreadyExists
	}
	storage.data[entity.ID] = entity
	return nil
}

//get article by id return map
func (storage *articleStorage) GetByID(ID int) (models.Article, error) {
	var resp models.Article
	resp, ok := storage.data[ID]
	if ok {
		return resp, nil
	}
	return resp, ErrorNotFound
}

// get all article return array
func (storage *articleStorage) GetList() []models.Article {
	var resp []models.Article
	for _, data := range storage.data {
		resp = append(resp, data)
	}
	return resp
}

//search article via title return array
func (storage *articleStorage) Search(str string) []models.Article {
	var resp []models.Article
	for _, data := range storage.data {
		if data.Content.Title == str {
			resp = append(resp, data)
		}
	}
	return resp
}

//update exist article
func (storage *articleStorage) Update(entity models.Article) error {
	if _, ok := storage.data[entity.ID]; ok {
		storage.data[entity.ID] = entity
		return nil
	}
	return ErrorNotUpdated
}

//delete article via id
func (storage *articleStorage) Delete(ID int) error {
	if _, ok := storage.data[ID]; ok {
		delete(storage.data, ID)
		return nil
	}
	return ErrorNotFound
}
