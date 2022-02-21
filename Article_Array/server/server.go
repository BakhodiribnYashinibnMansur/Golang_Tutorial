package main

import (
	"article/models"
	"article/storage"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var articleStorage storage.ArticleStorage

func main() {
	// initialize database
	articleStorage = storage.NewArticleStorage()
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

	//setup router
	router := gin.Default()

	//routers names
	router.GET("/article", GetArticleList)
	router.POST("/article", CreateArticle)
	router.GET("/article/search", SearchArticle)
	router.GET("/article/id:id", GetArticleByID)
	router.DELETE("/article/id:id", DeleteArticle)
	router.POST("/article/update", UpdateArticle)
	router.Run(":8080")
}

//get all article list
func GetArticleList(ctx *gin.Context) {
	resp := articleStorage.GetList()
	//response body json.
	ctx.JSON(200, resp)
}

// create article
func CreateArticle(ctx *gin.Context) {
	var article models.Article
	//read request json data.
	err := ctx.BindJSON(&article)
	if err != nil {
		log.Println(err)
		ctx.JSON(422, err.Error())
		return
	}
	err = articleStorage.Add(article)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, "message : Success Created Article")
}

func SearchArticle(ctx *gin.Context) {
	//read router query e.x. {{base_url}}article/search?query=Lorem
	searchText := ctx.Query("query")
	if searchText != "" {
		searchedList := articleStorage.Search(searchText)
		if len(searchedList) == 0 {
			//response json data gin.H
			ctx.JSON(400, gin.H{"message": " no result search ", "data": searchedList})
			return
		} else {
			ctx.JSON(200, searchedList)
			return
		}
	} else {
		ctx.JSON(400, "bad input")
	}
}

//get article by id
func GetArticleByID(ctx *gin.Context) {
	//read router param e.x. {{base_url}}article/id7
	paramID := ctx.Param("id")
	if paramID != "" {
		id, err := strconv.Atoi(paramID)
		if err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}
		getArticle, err := articleStorage.GetByID(id)
		if err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}
		ctx.JSON(200, gin.H{"data": getArticle})
	}
}

//delete Article
func DeleteArticle(ctx *gin.Context) {
	paramID := ctx.Param("id")
	if paramID != "" {
		id, err := strconv.Atoi(paramID)
		if err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			log.Println(err)
			return
		}
		err = articleStorage.Delete(id)
		if err != nil {
			ctx.JSON(200, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"message": "success deleted"})
	}
}

//update article via data
func UpdateArticle(ctx *gin.Context) {
	var article models.Article
	err := ctx.BindJSON(&article)
	if err != nil {
		log.Println(err)
		ctx.JSON(422, err.Error())
		return
	}
	err = articleStorage.Update(article)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, "message : Success Created Article")
}
