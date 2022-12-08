package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(campaignID int) ([]Transaction, error) {
	var transactions []Transaction

	// fungsi Order("id desc") akan mengurutkan berdasarkan id transaction dari urutan yg paling besar karena berdasarkan created_at paling terbaru
	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserID(userID int) ([]Transaction, error) {
	var transaction []Transaction
	// campaignImages tidak punya relasi langsung dengan Transaction
	// Flownya dari transaction -> relasi ke campaign -> relasi ke campaignImages
	// Caranya Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1") -> load campaign dan dari campaign akan load campaignImages yg berelasi dengan campaign
	// batasi dari campaignImages yg diload hanya is_primary = 1
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
