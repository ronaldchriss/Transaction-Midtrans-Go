package campaign

import "time"

type Campaign struct {
	ID            int
	UserID        int
	Title         string
	ShortDesc     string
	Desc          string
	GoalAmount    int
	CurrentAmount int
	Perks         string
	Slug          string
	BackerCount   int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
