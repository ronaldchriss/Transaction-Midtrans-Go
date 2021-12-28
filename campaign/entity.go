package campaign

import (
	"bwa_go/user"
	"time"
)

type Campaign struct {
	ID             int
	UserID         int
	Title          string
	ShortDesc      string
	Desc           string
	GoalAmount     int
	CurrentAmount  int
	Perks          string
	Slug           string
	BackerCount    int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CampaignImages []CampaignImage
	User           user.User
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
