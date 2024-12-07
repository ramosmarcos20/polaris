package api

import (
	"polaris/internal/controllers"
	"polaris/internal/middlewares"
	"polaris/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, userService *services.UserService) {
	// Instanciamos el controlador para usuarios
	userController := controllers.NewUserController(userService)

	// Rutas públicas (sin autenticación requerida)
	publicRoutes := router.Group("/api")
	{
		// Ruta de login (sin autenticación)
		publicRoutes.POST("/login", func(c *gin.Context) {
			controllers.Login(c, userService)
		})
	}

	// Rutas protegidas (requieren autenticación)
	protectedRoutes := router.Group("/api")
	// Middleware de autenticación para rutas protegidas
	protectedRoutes.Use(middlewares.AuthMiddleware())
	{
		users := protectedRoutes.Group("/users")
		{
			users.GET("/", userController.GetAllUsers)
			users.GET("/profile", userController.GetUserProfile)
		}
	}
}
