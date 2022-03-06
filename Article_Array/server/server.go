package main

import (
	_ "article/docs"
	"article/models"
	"article/storage"
	"article/storage/postgres"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// var articleStorage storage.ArticleStorage

// @title           Article API Docs
// @version         1.1
// @description     This is a sample article server celler server.
// @termsOfService  http://github.com/Golang_Tutorial/Article_Array
// @contact.name   Bakhodir Yashin Mansur
// @contact.email  phapp0224mb@gmail.com
// @host      localhost:8080

var articleRepo storage.ArticleRepository

func main() {
	// initialize database
	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		4000,
		"postgres",
		"0224",
		"article",
	)
	var err error

	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		fmt.Println(err)
	}
	articleRepo = postgres.NewArticleRepository(db)

	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	//routers names
	router.GET("/articles", GetArticleList)
	router.POST("/article", CreateArticle)
	router.GET("/article/search", SearchArticle)
	router.GET("/article/id:id", GetArticleByID)
	router.DELETE("/article/id:id", DeleteArticle)
	router.PUT("/article/update", UpdateArticle)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

// Show all articles godoc
// @Summary     Show an all article
// @Description  get all articles in memory
// @ID getlist-handler
// @Tags Article
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseSuccess
// @Failure 400,404 {object} models.DefaultError
// @Failure 500,503 {object} models.DefaultError
// @Router /articles [GET]
func GetArticleList(ctx *gin.Context) {
	offsetStr := ctx.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		ctx.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		fmt.Println(err)
	}

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.DefaultError{
			Message: err.Error(),
		})
		fmt.Println(err)
	}

	resp, err := articleRepo.GetList(models.Query{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.DefaultError{
			Message: err.Error(),
		})
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "successfull response",
		Data:    resp,
	})
}

// Create Article godoc
// @Summary     Create an Article
// @Description it create article based on on input data
// @ID create-handler
// @Tags   Article
// @Accept       json
// @Produce      json
// @Param        data    body  models.Article true "Article data"
// @Success      200   {object}      models.ResponseSuccess
// @Failure      400,404  {object}  models.ResponseError
// @Failure      500  {object}  models.DefaultError
// @Router       /article [POST]
func CreateArticle(ctx *gin.Context) {
	var article models.Article
	//read request json data.
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Message: err.Error(),
			Code:    400,
		})
		return
	}
	err = articleRepo.Create(article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Message: err.Error(),
			Code:    400,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
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
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
// @Success      200   {array}      models.Article
// @Failure      400,404  {object}  models.ResponseError
// @Failure      500  {object}  models.DefaultError
// @Router       /article/search [GET]
func SearchArticle(ctx *gin.Context) {
	offsetStr := ctx.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		ctx.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		fmt.Println(err)
	}

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.DefaultError{
			Message: err.Error(),
		})
		fmt.Println(err)
	}
	//read router query e.x. {{base_url}}article/search?query=Lorem
	searchText := ctx.Query("search")
	if searchText != "" {
		resp, err := articleRepo.Search(models.Query{Offset: offset, Limit: limit, Search: searchText})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.DefaultError{
				Message: err.Error(),
			})
			fmt.Println(err)
		}
		if len(resp) == 0 {
			ctx.JSON(http.StatusBadRequest, models.ResponseError{
				Message: "None Result Data", Code: 400,
			})
			return
		} else {
			ctx.JSON(200, models.ResponseSuccess{Message: "Successfull data", Data: resp})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
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
// @Failure      400,404  {object}  models.ResponseError
// @Failure      500  {object}  models.DefaultError
// @Router       /article/id{id} [GET]
func GetArticleByID(ctx *gin.Context) {
	//read router param e.x. {{base_url}}article/id7
	paramID := ctx.Param("id")
	if paramID != "" {
		id, err := strconv.Atoi(paramID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			log.Println(err)
			return
		}
		resp, err := articleRepo.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			return
		}
		ctx.JSON(http.StatusOK, models.ResponseSuccess{
			Message: "Success ",
			Data:    resp,
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
// @Success      200   {object}    models.ResponseSuccess
// @Failure      400,404  {object}  models.ResponseError
// @Failure      500  {object}  models.DefaultError
// @Router       /article/id{id} [DELETE]
func DeleteArticle(ctx *gin.Context) {
	paramID := ctx.Param("id")
	if paramID != "" {
		id, err := strconv.Atoi(paramID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			return
		}
		repo, err := articleRepo.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			return
		}
		if repo == 0 {
			ctx.JSON(http.StatusBadRequest, models.ResponseError{Message: "found id not found"})
			return
		}
		ctx.JSON(http.StatusOK, models.ResponseSuccess{
			Message: "Success Deleted Article",
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
// @Success      200   {object}      models.ResponseSuccess
// @Failure      400,404  {object}  models.ResponseError
// @Failure      500  {object}  models.DefaultError
// @Router       /article/update [PUT]
func UpdateArticle(ctx *gin.Context) {
	var article models.Article
	//read request json data.
	err := ctx.BindJSON(&article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Message: err.Error(),
			Code:    400,
		})
		return
	}
	repo, err := articleRepo.Update(article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Message: err.Error(),
			Code:    400,
		})
		return
	}
	if repo == 0 {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Message: "Not Found Id",
			Code:    400,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Success Updated Article",
	})
}
