package main
import (
	"example/Backend-Online-Stor-App/initializers"
	"example/Backend-Online-Stor-App/models"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main(){
	initializers.DB.AutoMigrate(&models.Product{})
	initializers.DB.AutoMigrate(&models.Order{})
}