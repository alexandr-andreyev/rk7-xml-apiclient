package rk7client

import "encoding/xml"

type RK7Query struct {
	XMLName    xml.Name     `xml:"RK7Query"`
	RK7Command []RK7Command `xml:"RK7Command"`
}

type RK7Command struct {
	CMD string `xml:"CMD,attr"`
}

type RK7QueryResult struct {
	//XMLName       xml.Name `xml:"RK7QueryResult"`
	ServerVersion   string          `xml:"ServerVersion,attr"`
	XmlVersion      string          `xml:"XmlVersion,attr"`
	NetName         string          `xml:"NetName,attr"`
	Status          string          `xml:"Status,attr"`
	Processed       string          `xml:"Processed,attr"`
	ArrivalDateTime string          `xml:"ArrivalDateTime,attr"`
	CommandResult   []CommandResult `xml:"CommandResult"`
}

type CommandResult struct {
	CMD        string     `xml:"CMD,attr"`
	Status     string     `xml:"Status,attr"`
	ErrorText  string     `xml:"ErrorText,attr"`
	DateTime   string     `xml:"DateTime,attr"`
	WorkTime   string     `xml:"WorkTime,attr"`
	SystemInfo SystemInfo `xml:"SystemInfo"`
}

type SystemInfo struct {
	SystemTime      string `xml:"SystemTime,attr"`
	ReqSysVer       string `xml:"ReqSysVer,attr"`
	ProcessID       string `xml:"ProcessID,attr"`
	ShiftDate       string `xml:"ShiftDate,attr"`
	RestCode        string `xml:"RestCode,attr"`
	BussinessPeriod struct {
		Id   string `xml:"id,attr"`
		Code string `xml:"code,attr"`
	} `xml:"BusinessPeriod"`
	CashGroup struct {
		Id   string `xml:"id,attr"`
		Code string `xml:"code,attr"`
		Name string `xml:"name,attr"`
	} `xml:"CashGroup"`
	Restaurant struct {
		Id   string `xml:"id,attr"`
		Code string `xml:"code,attr"`
		Name string `xml:"name,attr"`
	} `xml:"Restaurant"`
}
