package models
		// model sits at the centre of the project 
import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)
	// this struct model will act as the middle layer between the golang progam and our mongodb data base
	// since database understand json but our golang prgram doesnt thatswhy this middle layer is for conversion
type User struct{	// theze are the field that are needed for any data entry within our database int he mongodb which will be further processed bo jwt 
	ID  	 primtive.objectID    `bson:"_id"`
	First_name   *string     `json:"first_name validate: "required, min=2, max=100`  
	Last_name  *string     `json:"last_name validate: "required, min=2, max=100`
	Email 	*string     `json:"email" validate: "required, email, required"`
	Phone *string     `json:"phone" validate: "required"`
	Token *string     `json:"token"`
	User_type *string     `json:"user_type" validate: "required, eq=admin|eq=user"`	// eq is equal to which is either admin or else user as we have seen in golang
	Refresh_token *string     `json:"refresh_token"`
	Created_at     time.Time   `json:"created_at"`
	Updated_at  time.Time   `json:"updated_at"`
	User_id    string  `json:"user_id"`
}

