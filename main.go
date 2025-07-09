package main 

import(
	"os"
	"github.com/gin-gonic/gin"		// gin is helpful in creating a router for us and our project which might have been dnoe by other packages like fibre, etc but its a lot cleaner. 
	routes "jwt-projrect/routes"	// we are having two files int he routes folder amongst which the authRoute folder will be accessed by routes keyword
)

func main(){
	port=os.Getenv("PORT") {// we will have an env file and if it has any field as PORY then it will be imported to the object port in our main.go file
	if port == ""{
		port="8080"
	} 

	router := gin.New() //created a new router instance by using gin 
	router.Use(gin.Logger())	// will help in logging all the requests that will hit our server
	// router.Use(gin.Recovery()) // will help in recovering from any panics that may occur during the execution of our program
	
	router := gin.New()
	router.Use(gin.Logger()) // will help in logging all the requests that will hit our server

	routes.AuthRoutes(router) // this will import the authRouter file and will use the AuthRoutes function from that file to handle all the routes that are related to authentication
	routes.UserRoutes(router) // this will import the userRouter file and will use the UserRoutes function from that file to handle all the routes that are related to user management

	router.GET("/api-1", func(c *gin.Context) { // this is a test route to check if our server is running or not
		// as we might have seen in other projects we have to write w and r in order to define the req and response from the server and to the other microservices but here only c for gin.Context is reponsible to get access to the res & req
	})

	router.GET("/api-2", func(c *gin.Context) {
	c.JSON(200, gin.H{"success granted for api1"}) // this will return a json response with the status code 200 and a message that our server is running)
		// as in the basic and the core langauge we n=might have also set the different Headers for the status codes and the messages or the reponse that must have given by the server to the client but here
		//  gin.H is a shorthand for gin.H{"key": "value"} which is a map[string]interface{} that is used to create a JSON response easily

	})
	router.Run(":" + port) // to start the server with the port mentioned in the env variable or else the default port that has been assigned
}