package handler

import (
	"bwa_go/campaign"
	"bwa_go/helper"
	"bwa_go/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *CampaignHandler {
	return &CampaignHandler{service}
}

func (h *CampaignHandler) GetCampaign(c *gin.Context) {
	c.Query("user_id")
	UserID, err := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(UserID)
	if err != nil {
		response := helper.JsonResponse("Error to Get Campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("List of Campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) GetDetail(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Detail Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("Detail Campaigns", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Create Campaign Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	code := c.MustGet("codeUser").(user.User)

	input.User = code

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.JsonResponse("Create Campaign Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaign(newCampaign)
	response := helper.JsonResponse("Success to Create Campaign", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.JsonResponse("Error to Update Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Error to Update Campaign", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	code := c.MustGet("codeUser").(user.User)
	input.User = code

	updateCampaign, err := h.service.Update(inputID, input)
	if err != nil {
		response := helper.JsonResponse("Error to Update Campaign", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := campaign.FormatCampaign(updateCampaign)
	response := helper.JsonResponse("Success to Update Campaign", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
