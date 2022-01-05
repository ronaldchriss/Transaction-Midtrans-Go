package transaction

import (
	"bwa_go/campaign"
	"bwa_go/payment"
	"errors"

	"github.com/google/uuid"
)

type service struct {
	repository         Repository
	CampaignRepository campaign.Reprository
	paymentService     payment.Service
}

type Service interface {
	GetTransaction(input InputGetTransaction) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
	CreateTrans(input InputCreateTrans) (Transaction, error)
}

func NewService(repository Repository, CampaignRepository campaign.Reprository, paymentService payment.Service) *service {
	return &service{repository, CampaignRepository, paymentService}
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

func (s *service) GetByUserID(userID int) ([]Transaction, error) {
	trans, err := s.repository.GetByUserId(userID)
	if err != nil {
		return trans, err
	}
	return trans, nil
}

func (s *service) CreateTrans(input InputCreateTrans) (Transaction, error) {
	Transaction := Transaction{}
	Transaction.Amount = input.Amount
	Transaction.CampaignID = input.CampaignID
	Transaction.UserID = input.User.ID
	Transaction.Status = "pending"
	Transaction.Code = uuid.NewString()

	trans, err := s.repository.SaveTrans(Transaction)
	if err != nil {
		return trans, err
	}

	payment := payment.Transaction{
		ID:     trans.ID,
		Code:   trans.Code,
		Amount: trans.Amount,
	}

	URL, err := s.paymentService.GetPaymentURL(payment, input.User)
	if err != nil {
		return trans, err
	}

	trans.PaymentURL = URL
	trans, err = s.repository.Update(trans)
	if err != nil {
		return trans, err
	}

	return trans, nil
}
