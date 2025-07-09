package middleware

// now for the authentication we are using the middleware at first to chaeck if the connection or the request is secure or not
import(
	"fmt"
	"net/http"
	helper "github.com/AryanParashar24/jwt-project/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc{
	return func(c.*gin.Context){
		token := c.Request.Header.Get("token")	// this will get the token from the header of the request
		if token == ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorisation Header provided")})	// if the token is not present then it will abort the request and return a 401 status code with an error message
			c.Abort()
			return
		}
		
		claims, err := helper.ValidateToken(clienttoken)	// this will call the token helper and will validate the token
		if err!= nil{	
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()	// this will abort the request if the token is not valid
			return
		}
		
		c.Set("email", claims.Email)	// this will set the email in the context so that it can be used later in the request
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		
		c.Next() // this will call the next handler in the chain
	}
}