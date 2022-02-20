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

	//add new article

	//1 article
	var a1 models.Article
	a1.ID = 1
	a1.Title = "Lorem"
	a1.Body = "Lorem ipsum"
	var a1p models.Person = models.Person{
		Firstname: "John",
		Lastname:  "Doe",
	}
	a1.Author = a1p
	t := time.Now()
	a1.CreatedAt = &t
	err := articleStorage.Add(a1)
	if err != nil {
		fmt.Println(err)
	}

	//2 - article
	var a2 models.Article
	a2.ID = 2
	a2.Title = "Hello Nova"
	a2.Body = "NASA"
	var a2p models.Person = models.Person{
		Firstname: "John",
		Lastname:  "Carter",
	}
	a2.Author = a2p
	a2.CreatedAt = &t
	err = articleStorage.Add(a2)
	if err != nil {
		fmt.Println(err)
	}

	//3 article
	var a3 models.Article
	a3.ID = 3
	a3.Title = "Why 13 Reasons"
	a3.Body = "Cinema"
	var a3p models.Person = models.Person{
		Firstname: "Hanna",
		Lastname:  "Backer",
	}
	a3.Author = a3p
	a3.CreatedAt = &t
	err = articleStorage.Add(a3)
	if err != nil {
		fmt.Println(err)
	}
	err = articleStorage.Add(a1)
	if err != nil {
		fmt.Println(err)
	}

	//update 1 article update
	var updateA1 models.Article
	updateA1.ID = 1
	updateA1.Title = "Go"
	updateA1.Body = "Golang"
	var updateA1Person models.Person = models.Person{
		Firstname: "MRB",
		Lastname:  "Hero",
	}
	updateA1.Author = updateA1Person

	//delete article  by id
	deleteArticle := articleStorage.Delete(2)

	//Update article by ID
	updateArticle := articleStorage.Update(updateA1)

	//Search  Article via Article Title
	searchArticleList := articleStorage.Search("Why 13 Reasons")
	if len(searchArticleList) == 0 {
		fmt.Println(storage.ErrorSearch)
	}
	//get all article list
	articleList := articleStorage.GetList()

	//get by id
	getArticleByID, err := articleStorage.GetByID(1)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println()
	fmt.Println("Get List    : ", articleList)
	fmt.Println()
	fmt.Println("Create storage: ", articleStorage)
	fmt.Println()
	fmt.Println("Get by ID: ", getArticleByID)
	fmt.Println()
	fmt.Println("Delete by ID: ", deleteArticle)
	fmt.Println()
	fmt.Println("Update by ID: ", updateArticle)
	fmt.Println()
	fmt.Println("Search by string:  ", searchArticleList)
	fmt.Println()
	fmt.Println("Get List    : ", articleList)
	fmt.Println()
}
