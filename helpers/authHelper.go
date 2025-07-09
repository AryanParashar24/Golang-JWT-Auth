package helper
import(
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, userType string) (err error) {
	userType:= GetString("user_type")
	if userType != role{
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err	// so here the error is been returned which is retuned below in the MatchUserTypeToUid
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType:= c.GetString("user_Type") // As we defined earlier that the UserType will be a string in the model dir thatswhy we'll extract user_type from the postman or the database been storing that info 
	uid:= c.getString("uid")
	err=nil

	if userType == "USER" && uid != userId { // if the userType is USER and the uid is not equal to the userId then we will return an error
		err = errors.New("unauthorized access")	//errors.New because we are using the error package and is a function that creates a new error with the given message
		return err
	}

	CheckUserType(c, userType) // this will check if the userType is valid or not and if it is not then it will return an error
	return err
}