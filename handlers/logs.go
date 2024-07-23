package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FilterLogs(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		siteName := c.Query("site")
		startTimeStr := c.Query("start_time")
		endTimeStr := c.Query("end_time")
		startTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid start time format: %v", err)
			return
		}
		endTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid end time format: %v", err)
			return
		}
		logs, err := getFilteredLogs(db, siteName, startTime, endTime)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get filtered logs: %v", err)
			return
		}
		c.JSON(http.StatusOK, logs)
	}
}

func FilterLogsA(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		siteName := c.Query("siteA")
		startTimeStr := c.Query("start_timeA")
		endTimeStr := c.Query("end_timeA")
		var startTime, endTime time.Time
		var err error
		if startTimeStr != "" {
			startTime, err = time.Parse(time.RFC3339, startTimeStr)
			if err != nil {
				c.String(http.StatusBadRequest, "Invalid start time format: %v", err)
				return
			}
		}
		if endTimeStr != "" {
			endTime, err = time.Parse(time.RFC3339, endTimeStr)
			if err != nil {
				c.String(http.StatusBadRequest, "Invalid end time format: %v", err)
				return
			}
		}
		logs, err := getFilteredLogsAnalit(db, siteName, startTime, endTime)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get filtered logs: %v", err)
			return
		}
		c.JSON(http.StatusOK, logs)
	}
}
