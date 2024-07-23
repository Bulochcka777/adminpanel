package handlers

import (
	"adminBDvidj/models"
	"database/sql"
)

func getWidgetsFromDB(db *sql.DB) ([]models.WidgetA, error) {
	rows, err := db.Query(`SELECT ID, Name FROM Widgets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var widgetsa []models.WidgetA
	for rows.Next() {
		var widgeta models.WidgetA
		err := rows.Scan(&widgeta.ID, &widgeta.Name)
		if err != nil {
			return nil, err
		}
		widgetsa = append(widgetsa, widgeta)
	}
	return widgetsa, nil
}
