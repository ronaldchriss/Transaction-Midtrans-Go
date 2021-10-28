package campaign

import (
	"gorm.io/gorm"
)

type Reprository interface {
	Save(campaign Campaign) (Campaign, error)
}

type reprository struct {
	db *gorm.DB
}

func NewReprository(db *gorm.DB) *reprository {
	return &reprository{db}
}

func (r *reprository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
