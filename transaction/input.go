package transaction

import "bwa_go/user"

type InputGetTransaction struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type InputCreateTrans struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User       user.User 
}
