package jwt

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json : "first_name"`
	Last_name     *string            `json : "last_name"`
	Pass          *string            `json : "pass valiate : required min=8"`
	Email         *string            `json :  "email valiate : required email"`
	Token         *string            `json : "token"`
	Creation      time.Time          `json : "creation"`
	Last_login    time.Time          `json : "last_login"`
	Clusterconfig string             `json : "clusterconf"`
	Slackalert    string             `json : "slackalert"`
}
