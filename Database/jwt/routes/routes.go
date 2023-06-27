package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(incoming Routes *gin.Engine){
	incomingRoutes.POST("users/signup",controller.Signup)
	incomingRoutes.POST("users/login",controller.Login)
	
}

func UserRoutes(incoming Routes *gin.Engine){
	incomingRoutes.Use(middleware.Autentication)
	incomingRoutes.GET("/user",GetUser())
	incomingRoutes.GET("/user/:user_id",GetUser())
	
}