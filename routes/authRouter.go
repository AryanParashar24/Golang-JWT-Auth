package routes

import(
	controller "github.com/AryanParashar24/jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incoming Routes *gin.Engine){ // 
	incomingRoutes.POST("users/signup", controller.Signup()) // here the users/signup will be a method which'll be refering to the function signup in the conrtoller fodler of our app
	incomingRoutes.POST("users/login", controller.login()) //
}