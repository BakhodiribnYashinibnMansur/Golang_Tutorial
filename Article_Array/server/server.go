package main

import (
	_ "article/docs"
	"article/models"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
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

var db *sqlx.DB

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
	db, err = sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("database", db)
	// db, err = sql.Open("postgres", psqlConnString)
	// if err != nil {
	// 	fmt.Println(err)
	// }

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
	rows, err := db.Query(
		`SELECT
		ar.id,
		ar.title,
		ar.body,
		au.firstname,
		au.lastname,
		ar.created_at FROM article AS ar JOIN author AS au ON ar.author_id = au.id WHERE  ar.deleted_at IS  NULL `,
	)
	if err != nil {
		ctx.JSON(400, ResponseError{
			Message: err.Error(),
			Code:    400,
		})
		return
	}

	var arr []models.Article
	for rows.Next() {
		a := models.Article{}
		err = rows.Scan(
			&a.ID,
			&a.Title,
			&a.Body,
			&a.Author.Firstname,
			&a.Author.Lastname,
			&a.CreatedAt)
		arr = append(arr, a)
		if err != nil {
			log.Panic(err)
		}
	}
	defer rows.Close()
	ctx.JSON(200, ResponseSuccess{
		Message: "successfull response",
		Data:    arr,
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
	rows, err := db.Query(`INSERT INTO article(
		title,
		body,
		author_id)
VALUES ($1,$2,$3) RETURNING id;`, article.Title, article.Body, article.ID)
	if err != nil {
		ctx.JSON(422, ResponseError{
			Message: "Invalid ID",
			Code:    422,
		})
		return
	}
	var created_id int
	for rows.Next() {
		err = rows.Scan(
			&created_id,
		)
		if err != nil {
			ctx.JSON(400, ResponseError{
				Message: err.Error(),
				Code:    400,
			})
		}
	}
	if created_id == 0 {
		ctx.JSON(400, ResponseError{
			Message: err.Error(),
			Code:    400,
		})
	} else {
		ctx.JSON(200, ResponseSuccess{
			Message: "Success Created Article",
		})
	}

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
	fmt.Println(searchText)
	if searchText != "" {
		rows, err := db.Query(`
		SELECT
			ar.id,
		ar.title,
		ar.body,
		au.firstname,
		au.lastname,
		ar.created_at FROM article AS ar JOIN author AS au ON ar.author_id = au.id WHERE  ar.deleted_at IS  NULL AND ar.title LIKE  $1
		`, searchText)
		if err != nil {
			ctx.JSON(400, ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			return
		}
		var arr []models.Article
		for rows.Next() {
			a := models.Article{}
			err = rows.Scan(
				&a.ID,
				&a.Title,
				&a.Body,
				&a.Author.Firstname,
				&a.Author.Lastname,
				&a.CreatedAt)
			fmt.Println(arr)
			arr = append(arr, a)
			if err != nil {
				log.Panic(err)
			}
		}
		fmt.Println(arr)
		if len(arr) == 0 {
			//response json data gin.H
			ctx.JSON(400, ResponseError{
				Message: "None Result Data", Code: 400,
			})
			return
		} else {
			ctx.JSON(200, ResponseSuccess{Message: "Successfull data", Data: arr})
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
		rows, err := db.Query(`SELECT
			ar.id,
		ar.title,
		ar.body,
		au.firstname,
		au.lastname,
		ar.created_at FROM article AS ar JOIN author AS au ON ar.id = $1 WHERE  ar.deleted_at IS NULL `, id)
		// getArticle, err := articleStorage.GetByID(id)
		if err != nil {
			ctx.JSON(400, ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			log.Println(err)
			return
		}
		defer rows.Close()
		var article models.Article
		for rows.Next() {
			err = rows.Scan(
				&article.ID,
				&article.Title,
				&article.Body,
				&article.Author.Firstname,
				&article.Author.Lastname,
				&article.CreatedAt)
			if err != nil {
				log.Panic(err)
			}
		}
		if article.ID == 0 && article.CreatedAt == nil {
			ctx.JSON(424, ResponseError{Message: "invalid article id "})
		} else {
			ctx.JSON(200, ResponseSuccess{
				Message: "Successfull",
				Data:    article,
			})
		}
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
		tm := time.Now()
		rows, err := db.Exec(` UPDATE article SET  deleted_at = $1 WHERE id = $2  RETURNING id`, tm, id)
		if err != nil {
			ctx.JSON(400, ResponseError{
				Message: err.Error(),
				Code:    400,
			})
			return
		}
		updated, err := rows.RowsAffected()
		if err != nil {
			ctx.JSON(422, ResponseError{
				Message: err.Error(),
				Code:    422,
			})
			return
		}
		if updated == 0 {
			ctx.JSON(404, ResponseError{
				Message: "Not Found Id",
				Code:    404,
			})
		} else {
			ctx.JSON(200, ResponseSuccess{
				Message: "Success Updated Article",
			})
		}
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
	//read request json data.
	err := ctx.BindJSON(&article)
	if err != nil {
		ctx.JSON(400, ResponseError{
			Message: err.Error(),
			Code:    400,
		})
		return
	}
	rows, err := db.Exec(`UPDATE article SET title=$1, body=$2  WHERE  id=$3 RETURNING id;`, article.Title, article.Body, article.ID)
	if err != nil {
		ctx.JSON(422, ResponseError{
			Message: err.Error(),
			Code:    422,
		})
		return
	}
	updated, err := rows.RowsAffected()
	if err != nil {
		ctx.JSON(422, ResponseError{
			Message: err.Error(),
			Code:    422,
		})
		return
	}
	if updated == 0 {
		ctx.JSON(404, ResponseError{
			Message: "Not Found Id",
			Code:    404,
		})
	} else {
		ctx.JSON(200, ResponseSuccess{
			Message: "Success Updated Article",
		})
	}

}
