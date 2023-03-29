package controllers
import (
	// "net/http"
	"github.com/gin-gonic/gin"
	//"errors"
	"example/Backend-Online-Stor-App/models"
	"example/Backend-Online-Stor-App/initializers"
	
)
func ProductCreate(c *gin.Context){
	var body struct {
		Name string 
		Description string 
		Quantity int 
		Price int 
		Seller string 
		Category string 
	}
	//{Name: "Samsung Galaxy",Description:"Hape Flagship",Quantity:10,Price:5000000,Seller:"Samsung Corp",Category:"Electronic"}
	//{ "name" : "Indomie","description" :"Mie terbaik se Indonesia","quantity" : 100,"price" : 2000,"seller" : "PT Indofood","category" : "food"}
	c.Bind(&body)
	product := models.Product{Name: body.Name,Description:body.Description,Quantity:body.Quantity,Price:body.Price,Seller:body.Seller,Category:body.Category}
	result := initializers.DB.Create(&product)

	if result.Error != nil{
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post":product,
	})
}

func ProductsIndex(c *gin.Context){
	var products []models.Product
	initializers.DB.Find(&products)
	
	c.JSON(200, gin.H{
		"products":products,
	})
}

func ProductsShow(c *gin.Context){
	//Get id for url
	id := c.Param("id")
	var product models.Product
	initializers.DB.First(&product,id)
	
	c.JSON(200, gin.H{
		"products":product,
	})
}

func ProductsUpdate(c *gin.Context){
	id := c.Param("id")
	var body struct {
		Name string 
		Description string 
		Quantity int 
		Price int 
		Seller string 
		Category string 
	}
	c.Bind(&body)

	var product models.Product
	initializers.DB.First(&product,id)

	initializers.DB.Model(&product).Updates(models.Product{
		Name: body.Name,Description:body.Description,Quantity:body.Quantity,Price:body.Price,Seller:body.Seller,Category:body.Category,
	})
	c.JSON(200, gin.H{
		"products":product,
	})
}