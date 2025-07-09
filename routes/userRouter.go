package routes

import(
	controller "github.com/AryanParashar24/jwt-project/controllers" // importing the controller package from our project
	"github.com/gin-gonic/gin"
	"github.com/AryanParashar24/jwt-project/middleware"	// to ensure if the routes are protected routes in our service
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.USe(middlewear.Authenticate())	// to authenticate if the route is secure
	incomingRoutes.GET("/users", controller.GetUsers()) // this will get all the users from the database by using GetUsers method from the controller folder.
	incomingRoutes.GET("/users/:user_id", controller.GetUser()) // this will get a specific user id and then access data of that particular user
}	/* Now from the main.go file we came to the routers and then from the routers as wel can see above we will now move ahead to the database & controlelrs to ensure the secure connection &
  also get the user data access */




