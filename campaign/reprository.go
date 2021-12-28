package campaign

import (
	"gorm.io/gorm"
)

type Reprository interface {
	Save(campaign Campaign) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByUserID(UserID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
}

type reprository struct {
	db *gorm.DB
}

func NewReprository(db *gorm.DB) *reprository {
	return &reprository{db}
}

func (r *reprository) FindAll() ([]Campaign, error) {
	var Campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&Campaigns).Error
	if err != nil {
		return Campaigns, err
	}
	return Campaigns, nil
}

func (r *reprository) FindByUserID(UserID int) ([]Campaign, error) {
	var Campaigns []Campaign
	err := r.db.Where("user_id = ?", UserID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&Campaigns).Error
	if err != nil {
		return Campaigns, err
	}
	return Campaigns, nil
}

func (r *reprository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *reprository) FindByID(ID int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Preload("User").Preload("CampaignImages").Where("id=?", ID).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
