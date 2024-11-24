package repositories

import (
	"backend-test/config"
	"backend-test/models"
	"backend-test/utils"
)

func GetCustomerByEmail(email string) (*models.Customer, error) {
	var customers []models.Customer
	err := utils.ReadJSON(config.CustomerFile, &customers)
	if err != nil {
		return nil, err
	}

	for _, c := range customers {
		if c.Email == email {
			return &c, nil
		}
	}
	return nil, nil
}
