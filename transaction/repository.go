package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignID(CampaignID int) ([]Transaction, error)
	GetByUserId(UserID int) ([]Transaction, error)
	SaveTrans(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	GetByID(Code int) (Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(CampaignID int) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", CampaignID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetByID(Code int) (Transaction, error) {
	var transaction Transaction

	err := r.db.Where("code = ?", Code).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetByUserId(UserID int) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", UserID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) SaveTrans(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
