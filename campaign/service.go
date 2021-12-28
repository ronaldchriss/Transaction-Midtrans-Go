package campaign

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
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
