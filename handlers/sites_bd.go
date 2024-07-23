package handlers

import (
	"adminBDvidj/models"
	"database/sql"
)

func getSitesFromDB(db *sql.DB) ([]models.Site, error) {
	rows, err := db.Query(`SELECT ID, Name, Working, Debugging FROM Site`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sites []models.Site
	for rows.Next() {
		var site models.Site
		err := rows.Scan(&site.ID, &site.Name, &site.Working, &site.Debugging)
		if err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}
	return sites, nil
}
