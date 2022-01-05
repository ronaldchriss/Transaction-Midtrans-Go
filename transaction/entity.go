package transaction

import (
	"bwa_go/campaign"
	"bwa_go/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       user.User
	Campaign   campaign.Campaign
}
