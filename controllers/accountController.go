package controllers
import (
	// "net/http"
	"github.com/gin-gonic/gin"
	//"errors"
	"example/Backend-Online-Stor-App/models"
	"example/Backend-Online-Stor-App/initializers"
	"fmt"
)

func AccountCreate(c *gin.Context){
	var body struct {
		Name string 
		Email string 
		Password string
	}
	c.Bind(&body)
	account := models.Account{Name: body.Name,Email:body.Email,Password:body.Password}
	result := initializers.DB.Create(&account)

	if result.Error != nil{
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Account ":account,
	})
}

func AccountLogin(c *gin.Context){
	//Get id for url
	var body struct {
		Email string 
		Password string
	}
	c.Bind(&body)
	if(body.Email != ""){
		fmt.Println(body.Email)
		var account models.Account
		initializers.DB.First(&account,"email = ?",body.Email)
		if(account.Password == body.Password){
			c.JSON(200, gin.H{
				"message": "Welcome",
				"name" : account.Name,
			})
			return
		}
		c.JSON(404,gin.H{
			"message" : "Email or Password are incorrect",
		})
		return	
	}
	c.JSON(404,gin.H{
		"message" : "Please insert your Email and Password",
	})
	return	
		
	
}