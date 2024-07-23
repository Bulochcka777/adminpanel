package models

type TableData struct {
	ID     int    `json:"id"`
	Site   string `json:"site"`
	Widget string `json:"widget"`
	Goal   string `json:"goal"`
	Dvalue string `json:"dvalue"`
}

type Sitef struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Widget struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Goalf struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ResponseData struct {
	Sitesf    []Sitef     `json:"sites"`
	Widgets   []Widget    `json:"widgets"`
	Goalsf    []Goalf     `json:"goals"`
	TableData []TableData `json:"tableData"`
}
