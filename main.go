package main

import (
	// "net/http"
	"github.com/gin-gonic/gin"
	"example/Backend-Online-Stor-App/initializers"
	"example/Backend-Online-Stor-App/controllers"
	//"errors"
	
)


func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	router := gin.Default()
	router.POST("/products",controllers.PostsCreate)
	router.PUT("/products/:id",controllers.PostsUpdate)
	router.GET("/products",controllers.PostsIndex)
	router.GET("/products/:id",controllers.PostsShow)
	router.Run()
}