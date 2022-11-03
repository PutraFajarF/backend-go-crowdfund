package campaign

import "go-crowdfunding/user"

type GetCampaignDetailInput struct {
	// case ini tidak pakai JSON kalo untuk mengirim ID tidak pakai body (dipostman), tapi pakai URI
	// 3 beberapa cara untuk mengirim parameter => json (pakai body), query param, dan yg terakhir seakan2 menyatu dgn URL dengan pakai URI (ex: api/v1/campaign/1)
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type CreateCampaignImageInput struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}
