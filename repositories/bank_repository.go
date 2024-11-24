package repositories

import (
	"backend-test/config"
	"backend-test/models"
	"backend-test/utils"
)

func GetBankByID(id string) (*models.Bank, error) {
	var banks []models.Bank
	err := utils.ReadJSON(config.MerchantFile, &banks)
	if err != nil {
		return nil, err
	}

	for _, b := range banks {
		if b.Id == id {
			return &b, nil
		}
	}
	return nil, nil
}
