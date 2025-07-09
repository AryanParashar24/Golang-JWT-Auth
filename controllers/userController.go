package controllers

import(
	"context"
	"log"
	"fmt"
	"strconv"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helper "github.com/AryanParashar24/jwt-project/helper"
	"github.com/AryanParashar24/jwt-project/models"
	"github.com/AryanParashar24/jwt-project/database"
	"golang.org/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"	// for the update_at and created_at fields 
	"go.mongo.db.org/mongo-driver/bson/primitive" // for the primitive object ID to create a new user
	"go.mongodb.org/mongo-driver/mongo"	// for the mongo driver to connect to the database
)
// gin is just building a layer on top of our appliction shorting up the work of http and mux 
var userCollection = database.OpenCollection(database.Client, "users")
var validator = validator.New()


func HashPassword(password string) string{
	bycrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panis(err)
	}
	return string(bytes) // will return the hashed password as a string
}
func verifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword)) // will compare the provided password with the user password
	check := true // will check if the passwords match or not
	msg:= "" // will store the error message if the passwords do not match
	if err != nil {
		msg= fmt.Sprintf("email of password is incorrect") // if the passwords do not match then it will return false and the error message
		check = false
	}
	return true, "" // if the passwords match then it will return true and an empty string
}
func SignUp() gin.HandleFun(){	 		// now here int he signup function we'll have our validations, timestamps and create our tokens and refresh tokens 
	return func(c *gin.Context){
		var ctx, cancel:= context.WithTimeout(context.Background(), 100*time.Second)	// will handle the context background of our server with a timeout of 100 seconds
		var user model.User // because it will create a new user after getting logged in
		
		if err:= c.BindJSON(&user); err != nil {	// will bind the JSON data sent from the postman to the user model
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // will return a JSON response with the error code we defined int he error package thatswhy we are importing it in the way of 
			// error: err.Error
			return 
		}
		validation Err:= valideate.Struct(user)	// its the validation error which will be checking the validate field being defined int eh model directory of our project with each field for our
		//  database to validate if everything's right
		 if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()}) // will return a JSON response with the error code we defined int he error package thatswhy we are importing it
			//  in the way of error: err.Error
			return
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email}) // will count the number of clients in the database have been signed up 
		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H("error": "error occured while checking for the email"))	// so in other cases we use count to count the documents but here count is use to
			//  validate as if the email is already been registered then it means the count(error) will be more then zero(not null) thus we will handle it by the error handling as mentioned below
			//  by sending back the function as error occured while checking for the email
		}

		password:= HashPassword(*user.Password) // will hash the password using the HashPassword function defined above in the code
		user.Password = &password // will store the hashed password in the user model
		count, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone}) // will count the number of clients in the database have been signed up 
		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H("error": "error occured while checking for the phone"))   // same as email address
		}

		if count >0{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user with this email or phone already exists"}) 
		}

		user.Created_at, _= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339)) // will parse the current time in the RFC3339 format and store it in the created_at field of the user model
		user.Updated_at, _= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID() // will create a new object ID for the user model
		user.User_id = user.ID.Hex() // will convert the object ID to a hexadecimal string and store it in the user_id field of the user model
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *user.User_id) /* will generate a new token and refresh token and then the third field we dont know yet so will leave _ to generate tooken by iusing GenerateAllTokens function from the helper package
		// we willl have to create a function named GenerateAllTokens in the helper package to hangle these mathods.

		It will extract  user.email, first name last name user type and user id(that has been used in the above field to hex )*/
		user.Token = &token // will store the token in the user model
		user.Refresh_token = &refreshToken

		// now we will have to pass all these info to our database that will be inserted for one user or entry at a time to our database along with the context & user info which will be passed along with an insertionnumber
		resultInsertionNumber, InsertErr := userCollection.InsertOne(ctx, user)
		if InsertErr != nil{
			msg:= fmt.Sprintf("User items wasnt created")	// Spintf will format the string with the given arguments and return it as a string
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg}) // will return a JSON response with the error code we defined int he error package thatswhy we are importing it
			return 
		}
		defer cancel()
		c.JSON(http.StatusOk, resultInsertNumber)
	}
}
func Login() gin.HandleFun{
	return func(c *gin.Context){
		var ctx, cancel := context.WithTimeout(context.Background)(100*time.Second)
		var user models.User
		var foundUser models.User

		if err:= c.BindJSON(&user); err!= nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}
		err:= userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser) // will find the user in the database with the email which is been stored in the database as the email para and then decode it as
		//  for our golang application as the data is in JSON and model layer is the middle layer for it
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"}) // will return a JSON response with the status code 500 and an error message
			return 
		}
		passwordIsValid, msg := VerifyPassword(*user.password, *founderUser.Password)	/* this func VerifyPassword will verify the email id and password been entered by the user
		So, the email adddr that has been entered by the user will be checked by the user.Collection.FindOne function and if it is found in the database then it will be alligible to send their password, 
		Now the foundUser's password will be called out from the dtabase and will be checked if the password that has beeen entered by the user matches with the password being set for that found accound email id
		and if it does then it will be logged-in*/
		defer cancel()
		if passwordIsValid != true{
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return 
		}

		if foundUser.Email == nul{
			c.Json(http.StatusInternalServerError, gin.H{"error": "user not found"}) // will return a JSON response with the status code 500 and an error message
			return
		}
		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.User_type, *foundUser.User_id)	// we r sending all the information of struct to generate tokens which is in our helper
		helper.UpdateAllToken(token, refreshToken, foundUser.User_id) // will update the token and refresh token in the database for the user
		err = userColelction.FindOne(ctx, bson.M{"user_id":foundUser.User_id}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // will return a JSON response with the status code 500 and an error message
			return 
		}
		c.JSON(http.StatusOk, foundUser)
	}
}

func GuestUsers() gin.HandleFun{	// IT CAN ONLY BE USED BY this function is for extracting the info and details about the users registered
	return func(c *gin.Context){
		helper.CheckUser(c, "ADMIN"); err!= nil{
			c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
			return
		}	
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPage, err := strconv.Atoi(c.Query("page")) // it use to count/record the number of Pages and then convert it to an integer
		if err != nil || recordPage <1 {
			recordPage = 10 // if the page number is not valid then it will be set to 10
		}
		page, errl := strconv.Atoic(c.Query("page"))
		if err != nil {
			page =1 // if the page number is not valid then it will be set to 1'
		}

		startIndex := (page - 1) * recordPage // will calculate the start index of the page
		startIndex, err = strconv.Atoi(c.Query("start_index")) // will convert the start index to an integer

		//here it'll make the data more readable and observable by grouping the data and then pushing it to the database
		matchStage := bson.D{{"$match", bson.D{{}}}} // will match the stage of the pipeline to get the data from the database
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}} // will see  the grouped data othwerwise if we dont push then it will just be pushed wont wont be as observale as pushed

		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", bson.D{{"_id", null}}}, // refering to the id field of which we wants to group things with 
				{"total_count", 1}, // will show the total count of the users in the database
				{"user_items", bson.D{
					{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}}, // will slice the data to get the required number of users in the page}}	
			}}
		}
		// now we'll call our aggregate function because of which we built our match stage
		result, err:= userColelction.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage
		})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listening user items"})
		}
		var allUsers [bson.M]	// we creates a variable allUsers whcih is a slice of bson.M type which will hold the data of all the users in the database
		if err = result.All(ctx, &allUsers); err != nil {	// will
			log.Fatal(err)
		}
		c.JSON(http.StatusOk, allusers[0])
	}
}

// here we know that the detials and the info of the users can aonly be accessed by the admin and not by any other user thus we'll have to authenticate and route it in that way
func GetUser() gin.HandleFun(){ 	// gin gives access to handler functions
	return func(c *gin.Context){	// calling out the function defined in the gin package
		userId:= c.Param("user_id") // with the help of c we can access params sent from the postman using the api to the golang functions so we can access it using user id

		helper.MatchUserTypeToUid(c, userid); err!= nil{	// will use helper to check if the user is adnin or not
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err:= userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user) // will find the user in the database with the user id which is been stored in the database as the user_id param,
		//  and decode it as in the start we formed the model file because golang is unabel to understand JSON and thus moedl layer will act as a layer between the JSON and the golang to get it
		//  converted back to the server and then to the database
		defer cancel() // will cancel the context after the operation is done
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // will return a JSON response with the status code 500 and an error message
			return
		}
		c.JSON(http.StatusOK, user) // will return a JSON response with the status code 200 and the user details
		return 
	}
}