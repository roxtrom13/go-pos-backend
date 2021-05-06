package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roier/POS/controller"
	"github.com/roier/POS/database"
	"github.com/roier/POS/repository"
	"github.com/roier/POS/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = database.SetDatabase()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer database.CloseDatabase(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run(":8080")
}
