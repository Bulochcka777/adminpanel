package models

type SiteUpdate struct {
	SiteID    int    `json:"siteID"`
	Field     string `json:"field"` // Может быть "working" или "debugging"
	IsChecked bool   `json:"isChecked"`
}
