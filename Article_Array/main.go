package main

import (
	"article/models"
	"article/storage"
	"fmt"
	"time"
)

// var inMemory storage.ArticleStorage

func main() {

	// articleStorage = make(storage.ArticleStorage)
	articleStorage := storage.NewArticleStorage()
	var a1 models.Article
	a1.ID = 1
	a1.Title = "Lorem"
	a1.Body = "Lorem ipsum"
	var p models.Person = models.Person{
		Firstname: "John",
		Lastname:  "Doe",
	}
	a1.Author = p
	t := time.Now()
	a1.CreatedAt = &t
	err := articleStorage.Add(a1)
	if err != nil {
		fmt.Println(err)
	}
	var a2 models.Article
	a2.ID = 2
	a2.Title = "Data"
	a2.Body = "data ipsum"
	var a2p models.Person = models.Person{
		Firstname: "Jane",
		Lastname:  "Doe",
	}
	a2.Author = a2p
	a2.CreatedAt = &t
	err = articleStorage.Add(a2)
	if err != nil {
		fmt.Println(err)
	}

	getByIdArticle, err := articleStorage.GetByID(2)
	fmt.Println(getByIdArticle)
	err = articleStorage.Add(a1)
	if err != nil {
		fmt.Println(err)
	}

	var updateA1 models.Article
	updateA1.ID = 1
	updateA1.Title = "Go"
	updateA1.Body = "Golang "
	var updateA1Person models.Person = models.Person{
		Firstname: "MRB",
		Lastname:  "Hero",
	}
	updateA1.CreatedAt = &t
	updateA1.Author = updateA1Person

	err = articleStorage.Update(updateA1)
	if err != nil {
		fmt.Println(err)
	}

	articleStorage.Delete(2)
	articleList := articleStorage.GetList()
	fmt.Println(articleList)
	searchArticleList := articleStorage.Search("Data")
	fmt.Println(searchArticleList)
	fmt.Println(articleStorage)
}
