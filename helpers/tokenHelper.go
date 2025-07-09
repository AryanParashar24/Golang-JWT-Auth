package helper

import{
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/AryanParashar24/jwt-project/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for the primitive object ID to create a new user
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongo.db.org/mongo-driver/mongo/options" // for the mongo driver to connect to the database
}

type SignedDetails struct{
	Email string
	First_name string
	Last_name string
	Uid string
	User_type string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, lastName string, firstName string, userType string, uid string)(signedToekn string, signedRefinedToken string, err error){  	
	// here int he below class we can see that the email firstName,etc are the fields and methods from this function class while the Email First_name are from the above struct
	claims:= &SignedDetails{
		Email: email,
		First_name: firstName,	
		Last_name: lastName,
		Uid: uid,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 72).Unix(),
		},
	}
	refreshClaims:= &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 72).Unix(),
		},
	}

	//function of jwt which manages the creation of encrypted tokens
	token, err:= jwt.NewWithClaims(jwt.SigningMethodH256, claims).SignedString(SECRET_KEY) 
	refreshToken, err:= jwt.NewWithClaims(jwt.SigningMethodH256, refreshClaims).SignedString([]byte(SECRET_KEY)) // will create a new token and refresh token and then the third field we dont know yet so will leave _ to generate tooken by iusing GenerateAllTokens function from the helper package
	if err != nil{
		log.Panic(err)
		return 
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string)(claims *SignedDetails, msg string){ // here the claims are all of the type SignedDEtails as defined above in the struct SignedDetails in this file
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{},
	func(token *jwt.Token) (interface{}, error) { // here we are parsing the signedToken and then passing the SignedDetails struct to it
		return []byte(SECRET_KEY), nil // here we are returning the secret key which is used to sign the token
	},
	) // here we are parsing the signedToken and then passing the SignedDetails struct to it
		if err!= nil{
			msg=err.Error()
			return
		}
			
		claims, ok:= token.Claims.(*SignedDetails) // here we are type asserting the token.Claims to the SignedDetails struct so that we can access the fields of the struct
		if !ok{
			msg = fmt.Sprintf("The token is invalid")
			msg = err.Error()
			return 
		}

		if claims.ExpiresAt < time.Now().Local().Unix(){ // here we are checking if the token has expired or not
			msg = fmt.Sprintf("The token is expired")
			msg = err.Error()
			return
		}
		return claims.msg
}

if claims
 // here we will update the tokens of the login credentials snd the user everytime we are logging in and when the update has been done then its been updated and then appended to the earlier one and the tokens are changed everytime for the changed data
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string){
	var ctx, cancel = contect.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(ubdateObj, bson.E{"token", signedToken}) // here we are appending the token to the updateObj which is a primitive.D type and then we are adding the token to it
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken}) // here

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339)) // here we are parsing the time in the RFC3339 format and then we are updating the updated_at field in the database
	updateObj = append(updateObj, bson.E{"updated_at", Updated_at}) // here we are appending the updated_at field to the updateObj

	upsert:= true	// upsrt is a boolean value which is used to update the document if it exists or else insert a new document if it does not exist
	filter := bson.M{"_id": userId} // here we are creating a filter to find the document in the database with the given userId
	opts := options.UpdateOptions{
		Upsert: &upsert, // here we are setting the upsert option to true so that it will update the document if it exists or insert a new document if it does not exist
	}

	_, err:= userCollectionUpdateOne(	// here we'll update the info or the data to the database
		ctx,	// to update the user's data
		filter, // here we are passing the filter to find the document in the database
		bson.D{{"$set", updateObj}}, // here we are setting the updateObj to the document in the database
		&opts, // here we are passing the options to the updateOne function
	)

	defer cancel()

	if err != nil{
		log.Panic(err) // if there is an error while updating the document then we will log the error and panic
		return
	}
}
