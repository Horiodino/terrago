package usercontrollers

import (
	"context"
	"log"
	"net/http"
	"time"

	db "github.com/Horiodino/terrago/Database/jwt/db"
	"github.com/Horiodino/terrago/Database/jwt/helpers"
	models "github.com/Horiodino/terrago/Database/jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/tools/go/analysis/passes/defers"
)

// this valida
var validata = validator.New()
func Signup() gin.HandlerFunc{

	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return	
		}

		validatorErr := validate.Struct(user)
		if validatorErr != nil{
			 c.JSON(http.StatusBadRequest, gin.H{"err": validatorErr.Error()})
			 return
		}

		count , err := UserCollection.CountDocuments(ctx,bson.M{"email":user.Email})
		defer cancel()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		}

		// i know you are confused about this part
		// this is the part where we check if the email already exists
		// count is the number of documents that match the query
		// if the email not there then it will return 0 which 
		// if it returns any number greater than 0 then it means the email already exists
		// bevause the countdocuments function returns the number of documents that match the query
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		}
	}

}

func Login() {

}

func PasswordReset() {

}

func Password() {

}

func VerifyPassword() {

}
func GetUsers() {

}

var UserCollection *mongo.Collection = db.Opencollection(db.MongoClient, "user")

func GetUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		userID := c.Params("user_id")

		if err := helpers.MatchUser(c, userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User

		err := UserCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
		defer.cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}

}
