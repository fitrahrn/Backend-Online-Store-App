package controllers
import (
	// "net/http"
	"github.com/gin-gonic/gin"
	//"errors"
	"example/Backend-Online-Stor-App/models"
	"example/Backend-Online-Stor-App/initializers"
	//"fmt"
)
func OrderCreate(c *gin.Context){
	var body struct {
		UserID int 
		ProductID int 
		Quantity int
	}
	c.Bind(&body)
	product_id := body.ProductID
	var product models.Product
	initializers.DB.First(&product,product_id)
	if(product_id != 0){
		order := models.Order{UserID: body.UserID, ProductID: body.ProductID, Quantity:body.Quantity, Price:product.Price* body.Quantity, Payed:false}
		result := initializers.DB.Create(&order)

		if result.Error != nil{
			c.Status(400)
			return
		}
		c.JSON(200, gin.H{
			"order":order,
		})
		return
	}
	c.Status(400)
	return
}

func OrdersIndex(c *gin.Context){
	var orders []models.Order
	initializers.DB.Find(&orders)
	
	c.JSON(200, gin.H{
		"orders":orders,
	})
}

func OrdersShow(c *gin.Context){
	//Get id for url
	user_id := c.Param("id")
	var orders []models.Order
	initializers.DB.Table("orders").Where("user_id = ?", user_id).Scan(&orders)
	var products []models.Product
	var product_id []int
	for i:=0;i<len(orders);i++ {
		product_id = append(product_id,orders[i].ProductID)
	}
	initializers.DB.Find(&products,product_id)
	type Body struct {
		Name string `json:"name"`
		Description string `json:"description"`
		Quantity int `json:"quantity_buy"`
		TotalPrice int `json:"total_price"`
		Seller string	`json:"seller"`
		Category string	`json:"category"`
		Payed bool		`json:"payed"`
	}
	var product_body []Body
	for i:=0;i<len(products);i++{
		var body Body
		body.Name = products[i].Name
		body.Description = products[i].Description
		body.Quantity = orders[i].Quantity
		body.TotalPrice = orders[i].Price
		body.Seller = products[i].Seller
		body.Category = products[i].Category
		body.Payed = orders[i].Payed
		product_body = append(product_body,body)
	}
	c.JSON(200, gin.H{
		"products":product_body,
	})
}

func OrdersUpdate(c *gin.Context){
	id := c.Param("id")
	var body struct {
		UserID int 
		ProductID int 
		Quantity int
		Price int 
		Payed bool 
	}
	c.Bind(&body)

	var order models.Order
	initializers.DB.First(&order,id)

	initializers.DB.Model(&order).Updates(models.Order{
		UserID: body.UserID, ProductID: body.ProductID, Quantity:body.Quantity, Price:body.Price, Payed:body.Payed,
	})
	c.JSON(200, gin.H{
		"orders":order,
	})
}

func OrderDelete(c *gin.Context){
	id := c.Param("id")

	initializers.DB.Delete(&models.Order{},id)

	c.Status(200)
}