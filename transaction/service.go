package transaction

import (
	"errors"
	"go-crowdfunding/campaign"
)

type service struct {
	repository Repository
	// tambahkan campaignRepository untuk akses authorization
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsDetailInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsDetailInput) ([]Transaction, error) {
	// flow authorization
	// get campaign
	// check campaign.userid != user_id_yang_melakukan_request
	// karena akses service hanya ke repository, perlu ditambahkan campaignRepository pada service structnya
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
