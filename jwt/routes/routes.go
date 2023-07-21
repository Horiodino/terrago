package routes

import (
	"github.com/gin-gonic/gin"
	helpers "github.com/Horiodino/terrago/Database/jwt/helpers"
	models "github.com/Horiodino/terrago/Database/jwt/models"
)

func AuthRoutes(incoming Routes *gin.Engine){
	incomingRoutes.POST("users/signup",controller.Signup)
	incomingRoutes.POST("users/login",controller.Login)
	
}

func UserRoutes(incoming Routes *gin.Engine){
	incomingRoutes.Use(middleware.Autentication)
	incomingRoutes.GET("/users",GetUsers())
	incomingRoutes.GET("/user/:user_id",GetUser())
	
}