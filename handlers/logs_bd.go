package handlers

import (
	"adminBDvidj/models"
	"database/sql"
	"fmt"
	"time"
)

func getLogs(db *sql.DB) ([]models.Logs, error) {
	query := `
		SELECT
			logs.id,
			users.unique_id,
			users.reg_flag,
			users.reg_id,
			logs.visit,
			site.name,
			alias.domain,
			goal.name,
			widgets_goal.dvalue,
			widgets.name,
			logs.time
		FROM
			logs
		JOIN
			users ON logs.user_id = users.id
		JOIN
			widgets_goal ON logs.id_widgets_goal = widgets_goal.id
		JOIN
			site ON widgets_goal.id_site = site.id
		JOIN
			alias ON site.id = alias.site_id
		JOIN
			goal ON widgets_goal.id_goal = goal.id
		JOIN
			widgets ON widgets_goal.id_widgets = widgets.id
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()
	var logs []models.Logs
	for rows.Next() {
		var log models.Logs
		err := rows.Scan(&log.Log_ID, &log.Unique_ID, &log.Reg_flag, &log.Reg_ID, &log.Visit, &log.Site_Name, &log.Domain, &log.Goal_Name, &log.Goal_Dvalue, &log.Widget, &log.Time)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		logs = append(logs, log)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}
	return logs, nil
}

func getFilteredLogs(db *sql.DB, siteName string, startTime, endTime time.Time) ([]models.Logs, error) {
	query := `
		SELECT
			logs.ID AS Log_ID,
			users.Unique_ID,
			users.Reg_flag,
			users.Reg_id AS Reg_ID,
			logs.Visit,
			site.Name AS Site_Name,
			alias.Domain,
			goal.Name AS Goal_Name,
			widgets_goal.Dvalue AS Goal_Dvalue,
			widgets.Name AS Widget,
			logs.Time
		FROM
			Logs logs
		JOIN
			Users users ON logs.User_ID = users.ID
		JOIN
			Widgets_goal widgets_goal ON logs.ID_widgets_goal = widgets_goal.ID
		JOIN
			Site site ON widgets_goal.ID_site = site.ID
		JOIN
			Alias alias ON site.ID = alias.Site_ID
		JOIN
			Goal goal ON widgets_goal.ID_goal = goal.ID
		JOIN
			Widgets widgets ON widgets_goal.ID_widgets = widgets.ID
		WHERE
			($1 = '' OR site.name = $1) AND ($2 = '' OR logs.time >= $2) AND ($3 = '' OR logs.time <= $3)
	`
	rows, err := db.Query(query, siteName, startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()
	var logs []models.Logs
	for rows.Next() {
		var log models.Logs
		if err := rows.Scan(&log.Log_ID, &log.Unique_ID, &log.Reg_flag, &log.Reg_ID, &log.Visit, &log.Site_Name, &log.Domain, &log.Goal_Name, &log.Goal_Dvalue, &log.Widget, &log.Time); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return logs, nil
}

func getLogsAnalit(db *sql.DB) ([]models.Widget_Logs, error) {
	query := `
		SELECT
			Widget_Logs.id,
			users.unique_id,
			users.reg_id,
			site.name,
			goal.name,
			widgets.name,
			Widget_Logs.time
		FROM
			Widget_Logs
		JOIN
			users ON Widget_Logs.user_id = users.id
		JOIN
			site ON Widget_Logs.site_id = site.id
		JOIN
			goal ON Widget_Logs.goal_id = goal.id
		JOIN
			widgets ON Widget_Logs.widget_id = widgets.id
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var widget_Logs []models.Widget_Logs
	for rows.Next() {
		var widget_Log models.Widget_Logs
		err := rows.Scan(&widget_Log.IDA, &widget_Log.Unique_IDA, &widget_Log.Reg_IDA, &widget_Log.Site_NameA, &widget_Log.Goal_NameA, &widget_Log.WidgetA, &widget_Log.TimeA)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		widget_Logs = append(widget_Logs, widget_Log)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return widget_Logs, nil
}

func getFilteredLogsAnalit(db *sql.DB, siteName string, startTime, endTime time.Time) ([]models.Widget_Logs, error) {
	query := `
		SELECT
			Widget_Logs.id,
			users.unique_id,
			users.reg_id,
			site.name,
			goal.name,
			widgets.name,
			Widget_Logs.time
		FROM
			Widget_Logs
		JOIN
			users ON Widget_Logs.user_id = users.id
		JOIN
			site ON Widget_Logs.site_id = site.id
		JOIN
			goal ON Widget_Logs.goal_id = goal.id
		JOIN
			widgets ON Widget_Logs.widget_id = widgets.id
		WHERE
			($1 = '' OR site.name = $1)
			AND ($2 = '0001-01-01 00:00:00+00:00' OR Widget_Logs.time >= $2)
			AND ($3 = '0001-01-01 00:00:00+00:00' OR Widget_Logs.time <= $3)
	`
	rows, err := db.Query(query, siteName, startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var widget_Logs []models.Widget_Logs
	for rows.Next() {
		var widget_Log models.Widget_Logs
		if err := rows.Scan(&widget_Log.IDA, &widget_Log.Unique_IDA, &widget_Log.Reg_IDA, &widget_Log.Site_NameA, &widget_Log.Goal_NameA, &widget_Log.WidgetA, &widget_Log.TimeA); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		widget_Logs = append(widget_Logs, widget_Log)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return widget_Logs, nil
}
