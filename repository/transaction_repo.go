package repository

import (
	"avito-shop/config"
	"avito-shop/models"
)

func GetTransactionHistory(userID int) ([]models.History, []models.History, error) {
	var received []models.History
	var sent []models.History

	rows, err := config.DB.Raw(`
		SELECT users.login AS user, transactions.amount
		FROM transactions
		JOIN users ON transactions.from_user_id = users.user_id
		WHERE transactions.to_user_id = ?
	`, userID).Rows()
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var history models.History
		if err := rows.Scan(&history.User, &history.Amount); err != nil {
			return nil, nil, err
		}
		received = append(received, history)
	}

	rows, err = config.DB.Raw(`
		SELECT users.login AS user, transactions.amount
		FROM transactions
		JOIN users ON transactions.to_user_id = users.user_id
		WHERE transactions.from_user_id = ?
	`, userID).Rows()
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var history models.History
		if err := rows.Scan(&history.User, &history.Amount); err != nil {
			return nil, nil, err
		}
		sent = append(sent, history)
	}

	return received, sent, nil
}
