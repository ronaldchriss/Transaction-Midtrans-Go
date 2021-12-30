package campaign

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	Update(inputID GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error)
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

	id := uuid.New()

	compact := fmt.Sprintf("%s %s", input.Title, id.String())
	campaign.Slug = slug.Make(compact)

	newCamapaign, err := s.reprository.Save(campaign)
	if err != nil {
		return newCamapaign, err
	}

	return newCamapaign, nil
}

func (s *service) Update(inputID GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error) {
	campaign, err := s.reprository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != input.User.ID {
		return campaign, errors.New("You're not an owner this campaign")
	}

	campaign.Title = input.Title
	campaign.Desc = input.Desc
	campaign.ShortDesc = input.ShortDesc
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount

	updateCampaign, err := s.reprository.Update(campaign)
	if err != nil {
		return updateCampaign, err
	}

	return updateCampaign, nil
}
