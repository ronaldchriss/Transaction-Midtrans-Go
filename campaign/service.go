package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type service struct {
	reprository Reprository
}

func NewService(reprository Reprository) *service {
	return &service{reprository}
}

func (s *service) GetCampaigns(UserID int) ([]Campaign, error) {
	if UserID != 0 {
		campaigns, err := s.reprository.FindByUserID(UserID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	} else {
		campaigns, err := s.reprository.FindAll()
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.reprository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Title = input.Title
	campaign.Desc = input.Desc
	campaign.ShortDesc = input.ShortDesc
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	compact := fmt.Sprintf("%s %d", input.Title, input.User.ID)
	campaign.Slug = slug.Make(compact)

	newCamapaign, err := s.reprository.Save(campaign)
	if err != nil {
		return newCamapaign, err
	}

	return newCamapaign, nil
}
