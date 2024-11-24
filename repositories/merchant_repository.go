package repositories

import (
	"backend-test/config"
	"backend-test/models"
	"backend-test/utils"
)

func GetMerchantById(merchantId string) (*models.Merchant, error) {
	var merchants []models.Merchant
	err := utils.ReadJSON(config.MerchantFile, &merchants)
	if err != nil {
		return nil, err
	}

	for _, m := range merchants {
		if m.Id == merchantId {
			return &m, nil
		}
	}
	return nil, nil
}

func GetMerchantByEmail(email string) (*models.Merchant, error) {
	var merchants []models.Merchant
	err := utils.ReadJSON(config.MerchantFile, &merchants)
	if err != nil {
		return nil, err
	}

	for _, m := range merchants {
		if m.Email == email {
			return &m, nil
		}
	}
	return nil, nil
}
