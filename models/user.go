package models

type Users struct {
	ID          int
	Unique_ID   string
	Reg_flag    bool
	Reg_id      string
	Referrer    string
	Device_type string
	Browser     string
	Os          string
}
