package transaction

import (
	"bwa_go/campaign"
	"errors"
)

type service struct {
	repository         Repository
	CampaignRepository campaign.Reprository
}

type Service interface {
	GetTransaction(input InputGetTransaction) ([]Transaction, error)
}

func NewService(repository Repository, CampaignRepository campaign.Reprository) *service {
	return &service{repository, CampaignRepository}
}

func (s *service) GetTransaction(input InputGetTransaction) ([]Transaction, error) {

	campaign, err := s.CampaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("You're not an owner this campaign")
	}

	transaction, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
