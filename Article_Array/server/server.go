package main

import (
	"article/models"
	"article/storage"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "article/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var articleStorage storage.ArticleStorage

// @title           Article API Docs
// @version         1.1
// @description     This is a sample article server celler server.
// @termsOfService  http://github.com/Golang_Tutorial/Article_Array
// @contact.name   Bakhodir Yashin Mansur
// @contact.email  phapp0224mb@gmail.com
// @host      localhost:8080

type DefaultError struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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
	router.PUT("/article/update", UpdateArticle)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

// Show all articles godoc
// @Summary      Show an all article
// @Description  get all articles in memory
// @Tags   Article
// @ID                get-all-article-handler
// @Produce      json
// @Accept        json
// @Success      200  {array}  models.Article
// @Failure      500  {string}  DefaultError
// @Router       /article [GET]
func GetArticleList(ctx *gin.Context) {
	resp := articleStorage.GetList()
	//response body json.
	ctx.JSON(200, resp)
}

// Create Article godoc
// @Summary     Create an Article
// @Description it create article based on on input data
// @ID create-handler
// @Tags   Article
// @Accept       json
// @Produce      json
// @Param        data    body  models.Article true "Article data"
// @Success      200   {object}      ResponseSuccess
// @Failure      400,404  {object}  ResponseError
// @Failure      500  {object}  DefaultError
// @Router       /article [POST]
func CreateArticle(ctx *gin.Context) {
	var article models.Article
	//read request json data.
	err := ctx.BindJSON(&article)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, ResponseError{
			Message: err.Error(),
			Code:    400,
		})
		return
	}
	err = articleStorage.Add(article)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, ResponseError{
			Message: err.Error(),
			Code:    400,
		})
		return
	}
	ctx.JSON(200, ResponseSuccess{
		Message: "Success Created Article",
	})
}

// Search article via title  godoc
// @Summary    search article via title
// @Description it search on base via response title
// @ID search-handler
// @Tags   Article
// @Accept       json
// @Produce      json
// @Param        query    query  string true "Query Title"
// @Success      200   {array}      models.Article
// @Failure      400,404  {object}  ResponseError
// @Failure      500  {object}  DefaultError
// @Router       /article/search [GET]
func SearchArticle(ctx *gin.Context) {
	//read router query e.x. {{base_url}}article/search?query=Lorem
	searchText := ctx.Query("query")
	if searchText != "" {
		searchedList := articleStorage.Search(searchText)
		if len(searchedList) == 0 {
			//response json data gin.H
			ctx.JSON(400, ResponseError{
				Message: "None Result Data", Code: 400,
			})
			return
		} else {
			ctx.JSON(200, searchedList)
			return
		}
	} else {
		ctx.JSON(400, ResponseError{
			Message: "Bad Input",
			Code:    400,
		})
	}
}

// Get Article by Id  godoc
// @Summary   get article by id
// @Description it return Article which you send id
// @ID getarticle-by-id-handler
// @Tags   Article
// @Accept       json
// @Produce      json
// @Param        id   path  int     true "Param ID"
// @Success      200   {object}      models.Article
// @Failure      400,404  {object}  ResponseError
// @Failure      500  {object}  DefaultError
// @Router       /article/id{id} [GET]
func GetArticleByID(ctx *gin.Context) {
	//read router param e.x. {{base_url}}article/id7
	paramID := ctx.Param("id")
	if paramID != "" {
		id, err := strconv.Atoi(paramID)
		if err != nil {
			ctx.JSON(400, ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			log.Println(err)
			return
		}
		getArticle, err := articleStorage.GetByID(id)
		if err != nil {
			ctx.JSON(400, ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			log.Println(err)
			return
		}
		ctx.JSON(200, ResponseSuccess{
			Message: "Successfull",
			Data:    getArticle,
		})
	}
}

// Delete Article by Id  godoc
// @Summary   delete article by id
// @Description it delete Article which you send id of article
// @ID delete-handler
// @Tags   Article
// @Accept       json
// @Produce      json
// @Param        id   path  int     true "Param ID"
// @Success      200   {object}    ResponseSuccess
// @Failure      400,404  {object}  ResponseError
// @Failure      500  {object}  DefaultError
// @Router       /article/id{id} [DELETE]
func DeleteArticle(ctx *gin.Context) {
	paramID := ctx.Param("id")
	if paramID != "" {
		id, err := strconv.Atoi(paramID)
		if err != nil {
			ctx.JSON(400, ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			log.Println(err)
			return
		}
		err = articleStorage.Delete(id)
		if err != nil {
			ctx.JSON(400, ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			return
		}
		ctx.JSON(200, ResponseSuccess{
			Message: "Successfull Deleted",
		})
	}
}

// Update Article godoc
// @Summary     Update an Article
// @Description it update article based on on input data
// @ID update-handler
// @Tags   Article
// @Accept       json
// @Produce      json
// @Param        data    body  models.Article true "Article data"
// @Success      200   {object}      ResponseSuccess
// @Failure      400,404  {object}  ResponseError
// @Failure      500  {object}  DefaultError
// @Router       /article/update [PUT]
func UpdateArticle(ctx *gin.Context) {
	var article models.Article
	err := ctx.BindJSON(&article)
	if err != nil {
		log.Println(err)
		ctx.JSON(422, ResponseError{
			Message: err.Error(),
			Code:    422,
		})
		return
	}
	err = articleStorage.Update(article)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, ResponseError{
			Message: err.Error(),
			Code:    400,
		})
		return
	}
	ctx.JSON(200, ResponseSuccess{
		Message: "Successfull Updated",
	})
}
