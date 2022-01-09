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
	ProcessPayment(input TransactionNotif) error
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

func (s *service) ProcessPayment(input TransactionNotif) error {
	id := input.ID
	transaction, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancel"
	}

	trans, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.CampaignRepository.FindByID(trans.CampaignID)
	if err != nil {
		return err
	}

	if trans.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + trans.Amount

		_, err := s.CampaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil

}
