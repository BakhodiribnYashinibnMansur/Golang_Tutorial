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
var ErrorSearch = errors.New("not found data")

//Successes varableName
var SuccessUpdated = errors.New("SuccessFull Updated")
var SuccessAdded = errors.New("SuccessFull Added")
var SuccessDeleted = errors.New("SuccessFull Deleted")
var SuccessGetByID = errors.New("SuccessFull Got by ID")

//add new article
func (storage *articleStorage) Add(entity models.Article) error {
	for _, item := range storage.data {
		if item.ID == entity.ID {
			return ErrorAlreadyExists
		}
	}
	storage.data = append(storage.data, entity)
	return SuccessAdded
}

//get article by id return map
func (storage *articleStorage) GetByID(ID int) (models.Article, error) {
	var resp models.Article
	for _, item := range storage.data {
		if item.ID == ID {
			resp = item
			return resp, SuccessGetByID
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
			return SuccessUpdated
		}
	}
	return ErrorNotUpdated
}

//delete article via id
func (storage *articleStorage) Delete(ID int) error {
	for index, item := range storage.data {
		if item.ID == ID {
			storage.data = append(storage.data[:index], storage.data[index+1:]...)
			return SuccessDeleted
		}
	}
	return ErrorNotFound
}
