package handlers

import (
	"adminBDvidj/models"
	"database/sql"
	"fmt"
	"log"
)

func getUsers(db *sql.DB) ([]models.Users, error) {
	rows, err := db.Query(`SELECT ID, Unique_ID, Reg_flag, Reg_id, Referrer, Device_type, Browser, Os FROM Users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.Users
	for rows.Next() {
		var user models.Users
		err := rows.Scan(&user.ID, &user.Unique_ID, &user.Reg_flag, &user.Reg_id, &user.Referrer, &user.Device_type, &user.Browser, &user.Os)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func deleteEmptyRows(db *sql.DB) error {
	query := `
		DELETE FROM Users
		WHERE Unique_ID = '' OR Reg_flag = FALSE OR Reg_id = '' OR Referrer = '' OR Device_type = '' OR Browser = '' OR Os = ''
	`
	result, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error executing delete query: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %v", err)
	}
	log.Printf("Deleted %d rows with empty fields\n", rowsAffected)
	return nil
}
