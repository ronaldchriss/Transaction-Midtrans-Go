package campaign

import (
	"gorm.io/gorm"
)

type Reprository interface {
	Save(campaign Campaign) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByUserID(UserID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkNonPrimary(campaignID int) (bool, error)
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

func (r *reprository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *reprository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}
	return campaignImage, nil
}

func (r *reprository) MarkNonPrimary(campaignID int) (bool, error) {
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
