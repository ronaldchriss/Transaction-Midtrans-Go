package handler

import (
	"bwa_go/auth"
	"bwa_go/helper"
	"bwa_go/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	AuthService auth.Service
}

func NewUserHandler(userService user.Service, AuthService auth.Service) *userHandler {
	return &userHandler{userService, AuthService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.JsonResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.AuthService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.JsonResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)
	response := helper.JsonResponse("Account Has Been Register", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMassage := gin.H{"errors": err.Error()}

		response := helper.JsonResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.AuthService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.JsonResponse("Login Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)
	response := helper.JsonResponse("Successfuly Login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmail(c *gin.Context) {
	var input user.CheckEmail

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Email Has Been Registered", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	EmailAvailable, err := h.userService.CheckEmail(input)
	if err != nil {
		errorMassage := gin.H{"errors": "Server Error"}

		response := helper.JsonResponse("Email Check Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": EmailAvailable,
	}

	metamessage := "Email Has Been Registered"

	if EmailAvailable {
		metamessage = "Email Available"
	}

	response := helper.JsonResponse(metamessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.JsonResponse("Failed to upload avatar images", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := c.MustGet("codeUser").(user.User)
	userID := code.ID

	pathImages := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, pathImages)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.JsonResponse("Failed to upload avatar images", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.userService.SaveAvatar(userID, pathImages)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.JsonResponse("Failed to upload avatar images", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.JsonResponse("Success to upload avatar images", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FetchUser(c *gin.Context) {
	User := c.MustGet("codeUser").(user.User)

	format := user.FormatUser(User, "")

	response := helper.JsonResponse("Success to Fetch Data User", http.StatusOK, "success", format)

	c.JSON(http.StatusOK, response)

}
