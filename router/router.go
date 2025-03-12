package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sehha/config"
	"sehha/controller"
	"sehha/logger"
	"sehha/middleware"
	"sehha/repository"
	"sehha/service"
)

var (
	db      *gorm.DB = config.Connection()
	getYaml          = config.ReadYml{}
	port             = getYaml.GetConf().Port

	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userService    service.UserService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(userService)

	//agency
	agencyRepository repository.AgencyRepository = repository.NewAgencyRepository(db)
	agencyService    service.AgencyService       = service.NewAgencyService(agencyRepository, userRepository)
	agencyController controller.AgencyController = controller.NewAgencyController(agencyService)
)

func Router() {
	r := gin.Default()
	r.Use(enableCors)
	r.Use(logger.LoggerFile)

	r.GET("/", func(c *gin.Context) {
		logger.Info("test log", "log router")
		c.String(http.StatusOK, "welcome sehha service")
	})

	r.POST("/login", authController.Login)
	r.POST("/register", authController.RegisterUser)

	agencyRoutes := r.Group("/agency")
	{
		agencyRoutes.POST("", agencyController.Save)
	}

	r.GET("/test", middleware.AuthMiddleware(), func(context *gin.Context) {
		logger.Info("Test", "test data")
	})

	err := r.Run(":" + port)
	if err != nil {
		return
	}
}

func enableCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusOK)
		return
	}
	c.Next()
}
