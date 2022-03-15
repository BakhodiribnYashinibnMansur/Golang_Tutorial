package api

import (
	"article/api/handlers"
	"article/configs"
	"article/docs"
	_ "article/docs"

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

func SetUpAPI(router *gin.Engine, handler handlers.Handler, cfg configs.Config) {
	// programmaticaly set swagger info
	docs.SwaggerInfo_swagger.Title = cfg.App
	docs.SwaggerInfo_swagger.Version = cfg.Version
	docs.SwaggerInfo_swagger.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo_swagger.Schemes = []string{"http", "https"}

	router.GET("/articles", handler.GetArticleList)
	router.POST("/article", handler.CreateArticle)
	router.GET("/article/search", handler.SearchArticle)
	router.GET("/article/id:id", handler.GetArticleByID)
	router.DELETE("/article/id:id", handler.DeleteArticle)
	router.PUT("/article/update", handler.UpdateArticle)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
