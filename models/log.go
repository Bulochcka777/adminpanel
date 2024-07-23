package models

type Logs struct {
	Log_ID      int
	Unique_ID   string
	Reg_flag    bool
	Reg_ID      string
	Visit       string
	Site_Name   string
	Domain      string
	Goal_Name   string
	Goal_Dvalue string
	Widget      string
	Time        string
}

type Widget_Logs struct {
	IDA        int
	Unique_IDA string
	Reg_IDA    string
	Site_NameA string
	Goal_NameA string
	WidgetA    string
	TimeA      string
}
