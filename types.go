package rk7client

import (
	"encoding/xml"
	"fmt"
)

// RK7Query — корневой элемент XML-запроса к серверу.
type RK7Query struct {
	XMLName    xml.Name     `xml:"RK7Query"`
	RK7Command []RK7Command `xml:"RK7Command"`
}

// RK7Command описывает одну команду внутри запроса.
type RK7Command struct {
	CMD            string   `xml:"CMD,attr"`
	RefName        string   `xml:"RefName,attr,omitempty"`
	OnlyActive     string   `xml:"OnlyActive,attr,omitempty"`
	WithChildItems string   `xml:"WithChildItems,attr,omitempty"`
	WithMacroProp  string   `xml:"WithMacroProp,attr,omitempty"`
	PropMask       string   `xml:"PropMask,attr,omitempty"`
	RefItemIdent   string   `xml:"RefItemIdent,attr,omitempty"`
	Station        *Station `xml:"Station,omitempty"`
}

// Station задаёт кассовую станцию для команд, работающих в контексте конкретной станции.
type Station struct {
	Code string `xml:"Code,attr"`
}

// RK7QueryResult — корневой элемент XML-ответа сервера.
// Status="Ok" означает успех на уровне транспорта; детальный статус команды — в CommandResult.
type RK7QueryResult struct {
	ServerVersion   string         `xml:"ServerVersion,attr"`
	XmlVersion      string         `xml:"XmlVersion,attr"`
	NetName         string         `xml:"NetName,attr"`
	Status          string         `xml:"Status,attr"`
	Processed       string         `xml:"Processed,attr"`
	ArrivalDateTime string         `xml:"ArrivalDateTime,attr"`
	CommandResult   *CommandResult `xml:"CommandResult"`
}

// CommandResult содержит результат выполнения конкретной команды.
// ErrorText непуст при ошибке выполнения команды на стороне сервера.
type CommandResult struct {
	CMD          string        `xml:"CMD,attr"`
	Status       string        `xml:"Status,attr"`
	ErrorText    string        `xml:"ErrorText,attr"`
	DateTime     string        `xml:"DateTime,attr"`
	WorkTime     string        `xml:"WorkTime,attr"`
	SystemInfo   *SystemInfo   `xml:"SystemInfo"`
	RK7Reference *RK7Reference `xml:"RK7Reference"`
}

// SystemInfo содержит информацию о состоянии кассового сервера.
type SystemInfo struct {
	SystemTime string `xml:"SystemTime,attr"`
	ReqSysVer  string `xml:"ReqSysVer,attr"`
	ProcessID  string `xml:"ProcessID,attr"`
	ShiftDate  string `xml:"ShiftDate,attr"`
	RestCode   string `xml:"RestCode,attr"`
	BusinessPeriod struct {
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

// RK7Reference содержит версию данных и список элементов справочника.
type RK7Reference struct {
	DataVersion    string `xml:"DataVersion,attr"`
	TotalItemCount string `xml:"TotalItemCount,attr"`
	Items          Items  `xml:"Items"`
}

// Items — обёртка вокруг списка элементов справочника.
type Items struct {
	Item []Item `xml:"Item"`
}

// Item — элемент справочника (блюдо, категория и т.д.).
// Динамические атрибуты (например, PRICETYPES-3) доступны через Attributes и метод GetPrice.
type Item struct {
	Ident           string      `xml:"Ident,attr"`
	GUIDString      string      `xml:"GUIDString,attr"`
	Code            string      `xml:"Code,attr"`
	Name            string      `xml:"Name,attr"`
	Status          string      `xml:"Status,attr"`
	Parent          string      `xml:"Parent,attr"`
	MainParentIdent string      `xml:"MainParentIdent,attr"`
	Modscheme       string      `xml:"ModiScheme,attr"`
	CategPath       string      `xml:"CategPath,attr"`
	Attributes      []xml.Attr  `xml:",any,attr"`
	Selectors       []TSelector `xml:"RIChildItems>TSelector"`
}

// GetPrice возвращает цену элемента для указанного типа цены (PRICETYPES-N).
// Возвращает пустую строку, если цена для данного типа не задана.
func (i Item) GetPrice(priceType int) string {
	key := fmt.Sprintf("PRICETYPES-%d", priceType)
	for _, a := range i.Attributes {
		if a.Name.Local == key {
			return a.Value
		}
	}
	return ""
}

// TSelector — селектор внутри группы селекторов (SelectorGroup).
type TSelector struct {
	ItemIdent string            `xml:"ItemIdent,attr"`
	Ident     string            `xml:"Ident,attr"`
	Name      string            `xml:"Name,attr"`
	Status    string            `xml:"Status,attr"`
	Children  []TSelectorDetail `xml:"RIChildItems>TSelectorDetail"`
}

// TSelectorDetail — элемент внутри селектора: ссылка на блюдо (toDish) или вложенный селектор (toSelector).
type TSelectorDetail struct {
	ItemIdent   string `xml:"ItemIdent,attr"`
	TradeObject string `xml:"TradeObject,attr"`
	ObjectSifr  string `xml:"ObjectSifr,attr"`
	OrderNum    string `xml:"OrderNum,attr"`
}
