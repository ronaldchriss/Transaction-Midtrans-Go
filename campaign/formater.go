package campaign

import "strings"

type CampaignFormatter struct {
	ID            int    `json: "id"`
	UserID        int    `json: "user_id"`
	Title         string `json : "title"`
	ImageUrl      string `json : "image_url"`
	ShortDesc     string `json : "short_desc"`
	GoalAmount    int    `json: "goal_amount"`
	CurrentAmount int    `json: "current_amount"`
	Slug          string `json: "slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	CampaignFormatter := CampaignFormatter{}
	CampaignFormatter.ID = campaign.ID
	CampaignFormatter.UserID = campaign.UserID
	CampaignFormatter.Title = campaign.Title
	CampaignFormatter.ShortDesc = campaign.ShortDesc
	CampaignFormatter.GoalAmount = campaign.GoalAmount
	CampaignFormatter.CurrentAmount = campaign.CurrentAmount
	CampaignFormatter.Slug = campaign.Slug
	CampaignFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		CampaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return CampaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFotmatter struct {
	ID            int                    `json: "id"`
	Title         string                 `json : "title"`
	ImageUrl      string                 `json : "image_url"`
	ShortDesc     string                 `json : "short_desc"`
	Desc          string                 `json : "desc"`
	GoalAmount    int                    `json: "goal_amount"`
	CurrentAmount int                    `json: "current_amount"`
	Slug          string                 `json: "slug"`
	UserID        int                    `json: "user_id"`
	Perks         []string               `json: perks`
	User          UserDetailFormatter    `json: user`
	Images        []ImageDetailFormatter `json: images`
}

type UserDetailFormatter struct {
	Name     string `json: "name"`
	ImageUrl string `json:  "image_url"`
}

type ImageDetailFormatter struct {
	ImageUrl string `json: "image_url"`
	Primary  bool   `json:  "is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFotmatter {
	CampaignDetailFotmatter := CampaignDetailFotmatter{}
	CampaignDetailFotmatter.ID = campaign.ID
	CampaignDetailFotmatter.UserID = campaign.UserID
	CampaignDetailFotmatter.Title = campaign.Title
	CampaignDetailFotmatter.ShortDesc = campaign.ShortDesc
	CampaignDetailFotmatter.Desc = campaign.Desc
	CampaignDetailFotmatter.GoalAmount = campaign.GoalAmount
	CampaignDetailFotmatter.CurrentAmount = campaign.CurrentAmount
	CampaignDetailFotmatter.Slug = campaign.Slug
	CampaignDetailFotmatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		CampaignDetailFotmatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	CampaignDetailFotmatter.Perks = perks

	data := campaign.User

	UserDetailFormatter := UserDetailFormatter{}
	UserDetailFormatter.Name = data.Name
	UserDetailFormatter.ImageUrl = data.AvatarFileName

	CampaignDetailFotmatter.User = UserDetailFormatter

	images := []ImageDetailFormatter{}

	for _, image := range campaign.CampaignImages {
		ImageDetailFormatter := ImageDetailFormatter{}
		ImageDetailFormatter.ImageUrl = image.FileName
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		ImageDetailFormatter.Primary = isPrimary

		images = append(images, ImageDetailFormatter)
	}

	CampaignDetailFotmatter.Images = images

	return CampaignDetailFotmatter

}
