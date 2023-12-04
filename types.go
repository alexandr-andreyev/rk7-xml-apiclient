package rk7client

import "encoding/xml"

type RK7Query struct {
	XMLName    xml.Name     `xml:"RK7Query"`
	RK7Command []RK7Command `xml:"RK7Command"`
}

type RK7Command struct {
	CMD            string   `xml:"CMD,attr"`
	RefName        string   `xml:"RefName,attr"`
	OnlyActrive    string   `xml:"OnlyActive,attr,omitempty"`
	WithChildItems string   `xml:"WithChildItems,attr,omitempty"`
	WithMacroProp  string   `xml:"WithMacroProp,attr,omitempty"`
	PropMask       string   `xml:"PropMask,attr,omitempty"`
	Station        *Station `xml:"Station,omitempty"`
}

type Station struct {
	Code string `xml:"Code,attr"`
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
	CMD          string        `xml:"CMD,attr"`
	Status       string        `xml:"Status,attr"`
	ErrorText    string        `xml:"ErrorText,attr"`
	DateTime     string        `xml:"DateTime,attr"`
	WorkTime     string        `xml:"WorkTime,attr"`
	SystemInfo   *SystemInfo   `xml:"SystemInfo"`
	RK7Reference *RK7Reference `xml:"RK7Reference"`
	PriceScale   *PriceScale   `xml:"PriceScale"`
	TradeGroup   *TradeGroup   `xml:"TradeGroup"`
	Dishes       *Dishes       `xml:"Dishes"`
}

type PriceScale struct {
	Id   string `xml:"id,attr"`
	Code string `xml:"code,attr"`
	Name string `xml:"name,attr"`
}

type TradeGroup struct {
	Id   string `xml:"id,attr"`
	Code string `xml:"code,attr"`
	Name string `xml:"name,attr"`
}

type Dishes struct {
	Item struct {
		Ident string `xml:"Ident,attr"`
		Price string `xml:"Price,attr"`
	}
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

type RK7Reference struct {
	DataVersion    string `xml:"DataVersion,attr"`
	TotalItemCount string `xml:"TotalItemCount"`
	Items          Items  `xml:"Items"`
}

type Items struct {
	Item []Item `xml:"Item"`
}

type Item struct {
	Ident           string     `xml:"Ident,attr"`
	GUIDString      string     `xml:"GUIDString,attr"`
	Code            string     `xml:"Code,attr"`
	Name            string     `xml:"Name,attr"`
	Status          string     `xml:"Status,attr"`
	Parent          string     `xml:"Parent,attr"`
	MainParentIdent string     `xml:"MainParentIdent,attr"`
	Price           string     `xml:"PRICETYPES-3,attr"`
	Modscheme       string     `xml:"ModiScheme,attr"`
	CategPath       string     `xml:"CategPath,attr"`
	Attributes      []xml.Attr `xml:",any,attr"`
}
