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

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
