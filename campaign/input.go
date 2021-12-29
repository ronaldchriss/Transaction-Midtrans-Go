package campaign

import "bwa_go/user"

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding: "required"`
}

type CreateCampaignInput struct {
	Title      string `json:"title" binding:"required"`
	ShortDesc  string `json:"short_desc" binding:"required"`
	Desc       string `json:"desc" binding:"required"`
	GoalAmount int    `json:"goal_amount" binding:"required"`
	Perks      string `json:"perks" binding:"required"`
	User       user.User
}
