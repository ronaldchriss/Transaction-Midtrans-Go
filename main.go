package main

import (
	"bwa_go/auth"
	"bwa_go/campaign"
	"bwa_go/handler"
	"bwa_go/helper"
	"bwa_go/payment"
	"bwa_go/transaction"
	"bwa_go/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "golang:lTwiUyonc10rpuUK@tcp(db-mysql-sgp1-47894-do-user-10563339-0.b.db.ondigitalocean.com:25060)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	UserRepository := user.NewRepository(db)
	CampaignRepository := campaign.NewReprository(db)
	TransactionRepository := transaction.NewRepository(db)

	UserService := user.NewService(UserRepository)
	AuthService := auth.NewService()
	CampaignService := campaign.NewService(CampaignRepository)
	PaymentService := payment.NewService()
	TransService := transaction.NewService(TransactionRepository, CampaignRepository, PaymentService)

	userHandler := handler.NewUserHandler(UserService, AuthService)
	campaignHandler := handler.NewCampaignHandler(CampaignService)
	transHandler := handler.NewTransactionHandler(TransService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images/", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/emailcheckers", userHandler.CheckEmail)
	api.POST("/avatar", authMiddleware(AuthService, UserService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaign)
	api.GET("/campaigns/:id", campaignHandler.GetDetail)
	api.POST("/campaigns", authMiddleware(AuthService, UserService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(AuthService, UserService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(AuthService, UserService), campaignHandler.UploadImage)

	api.GET("/campaigns/:id/transaction", authMiddleware(AuthService, UserService), transHandler.GetTransaction)
	api.GET("/transaction", authMiddleware(AuthService, UserService), transHandler.GetTransactionByUser)
	api.POST("/transaction", authMiddleware(AuthService, UserService), transHandler.CreateTrans)
	api.POST("/transaction/notification", transHandler.GetNotif)
	router.Run()
}

func authMiddleware(AuthService auth.Service, UserService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bareer") {
			response := helper.JsonResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//Bareer Token Header
		tokenString := ""
		codeToken := strings.Split(authHeader, " ")
		if len(codeToken) == 2 {
			tokenString = codeToken[1]
		}

		token, err := AuthService.ValidateToken(tokenString)
		if err != nil {
			response := helper.JsonResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.JsonResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := UserService.GetUserByID(userID)
		if err != nil {
			response := helper.JsonResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("codeUser", user)

	}
}
