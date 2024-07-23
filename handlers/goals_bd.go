package handlers

import (
	"adminBDvidj/models"
	"database/sql"
)

func getGoals(db *sql.DB) ([]models.Goal, error) {
	rows, err := db.Query(`SELECT ID, Name, About FROM Goal`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var goals []models.Goal
	for rows.Next() {
		var goal models.Goal
		err := rows.Scan(&goal.ID, &goal.Name, &goal.About)
		if err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}
	return goals, nil
}
