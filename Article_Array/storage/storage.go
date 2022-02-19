package storage

import (
	"article/models"
	"errors"
)

// type articleStorage map[int]models.Article
func NewArticleStorage() articleStorage {
	return articleStorage{
		data: make([]models.Article, 0, 999),
	}
}

type articleStorage struct {
	data []models.Article
}

//Errors varableName
var ErrorAlreadyExists = errors.New("already exists id")
var ErrorNotFound = errors.New("not found id ")
var ErrorNotUpdated = errors.New("not update because not found id")

//add new article
func (storage *articleStorage) Add(entity models.Article) error {
	for _, item := range storage.data {
		if item.ID == entity.ID {
			return ErrorAlreadyExists
		}
	}
	storage.data = append(storage.data, entity)
	return nil
}

//get article by id return map
func (storage *articleStorage) GetByID(ID int) (models.Article, error) {
	var resp models.Article
	for _, item := range storage.data {
		if item.ID == ID {
			resp = item
			return resp, nil
		}
	}
	return resp, ErrorNotFound
}

// get all article return array
func (storage *articleStorage) GetList() []models.Article {
	return storage.data
}

//search article via title return array
func (storage *articleStorage) Search(str string) []models.Article {
	var resp []models.Article
	for _, item := range storage.data {
		if item.Title == str {
			resp = append(resp, item)
			return resp
		}
	}
	return resp
}

//update exist article
func (storage *articleStorage) Update(entity models.Article) error {
	for index, item := range storage.data {
		if item.ID == entity.ID {
			storage.data[index] = entity
			return nil
		}
	}
	return ErrorNotUpdated
}

//delete article via id
func (storage *articleStorage) Delete(ID int) error {
	for index, item := range storage.data {
		if item.ID == ID {
			storage.data = append(storage.data[:index], storage.data[index+1:]...)
			return nil
		}
	}
	return ErrorNotFound
}
