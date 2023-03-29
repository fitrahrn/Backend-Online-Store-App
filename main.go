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
	router.POST("/products",controllers.ProductCreate)
	router.PUT("/products/:id",controllers.ProductsUpdate)
	router.GET("/products",controllers.ProductsIndex)
	router.GET("/products/:id",controllers.ProductsShow)
	router.POST("/carts",controllers.OrderCreate)
	router.PUT("/carts/:id",controllers.OrdersUpdate)
	router.GET("/carts",controllers.OrdersIndex)
	router.GET("/carts/:id",controllers.OrdersShow)
	router.DELETE("/carts/:id",controllers.OrderDelete)
	router.Run()
}