package repositories

import (
	"backend-test/config"
	"backend-test/models"
	"backend-test/utils"
)

func AddHistory(entry models.History) error {
	var history []models.History
	err := utils.ReadJSON(config.HistoryFile, &history)
	if err != nil {
		return err
	}

	history = append(history, entry)
	return utils.WriteJSON(config.HistoryFile, history)

}
