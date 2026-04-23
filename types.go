package rk7client

import (
	"encoding/xml"
	"fmt"
	"io"
)

type RK7Query struct {
	XMLName    xml.Name     `xml:"RK7Query"`
	RK7Command []RK7Command `xml:"RK7Command"`
}

type RK7Command struct {
	CMD            rk7cmd         `xml:"CMD,attr"`
	RefName        rk7ref         `xml:"RefName,attr,omitempty"`
	CardCode       string         `xml:"CardCode,attr,omitempty"`
	OnlyActrive    onlyactive     `xml:"OnlyActive,attr,omitempty"`
	WithChildItems withChildItems `xml:"WithChildItems,attr,omitempty"`
	WithMacroProp  withMacroProp  `xml:"WithMacroProp,attr,omitempty"`
	PropMask       string         `xml:"PropMask,attr,omitempty"`
	RefItemIdent   string         `xml:"RefItemIdent,attr,omitempty"`
	Station        *Station       `xml:"Station,omitempty"`
	Waiter         *Waiter        `xml:"Waiter,omitempty"`
	RegisteredOnly registeredOnly `xml:"registeredOnly,attr,omitempty"`
	OnlyOpened     onlyOpened     `xml:"onlyOpened,attr,omitempty"`
	PROPFILTERS    []PROPFILTER   `xml:"PROPFILTERS>PROPFILTER,omitempty"`
}

type PROPFILTER struct {
	Name  string `xml:"Name,attr,omitempty"`
	Value string `xml:"Value,attr,omitempty"`
}

type Station struct {
	ID   string `xml:"id,attr,omitempty"`
	Code string `xml:"code,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

type Table struct {
	ID   string `xml:"id,attr,omitempty"`
	Code string `xml:"code,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

type OrderCategory struct {
	ID   string `xml:"id,attr,omitempty"`
	Code string `xml:"code,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

type OrderType struct {
	ID   string `xml:"id,attr,omitempty"`
	Code string `xml:"code,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

type Waiter struct {
	ID   string `xml:"id,attr,omitempty"`
	Code string `xml:"code,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

type RK7QueryResult struct {
	ServerVersion   string          `xml:"ServerVersion,attr"`
	XmlVersion      string          `xml:"XmlVersion,attr"`
	NetName         string          `xml:"NetName,attr"`
	Status          string          `xml:"Status,attr"`
	Processed       string          `xml:"Processed,attr"`
	ArrivalDateTime string          `xml:"ArrivalDateTime,attr"`
	CommandResult   []CommandResult `xml:"CommandResult"`
}

type RK7RefList struct {
	XMLName      xml.Name       `xml:"RK7RefList"`
	Count        string         `xml:"Count,attr"`
	RK7Reference []RK7Reference `xml:"RK7Reference"`
}

type RK7Reference struct {
	DataVersion    string `xml:"DataVersion,attr,omitempty"`
	RefName        string `xml:"RefName,attr,omitempty"`
	ClassName      string `xml:"ClassName,attr,omitempty"`
	TotalItemCount string `xml:"TotalItemCount,attr,omitempty"`
	Count          string `xml:"Count,attr,omitempty"`
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
func (i Item) GetPrice(priceType int) string {
	key := fmt.Sprintf("PRICETYPES-%d", priceType)
	for _, a := range i.Attributes {
		if a.Name.Local == key {
			return a.Value
		}
	}
	return ""
}

// TSelector — селектор внутри группы селекторов.
type TSelector struct {
	ItemIdent string            `xml:"ItemIdent,attr"`
	Ident     string            `xml:"Ident,attr"`
	Name      string            `xml:"Name,attr"`
	Status    string            `xml:"Status,attr"`
	Children  []TSelectorDetail `xml:"RIChildItems>TSelectorDetail"`
}

// TSelectorDetail — элемент внутри селектора: ссылка на блюдо или вложенный селектор.
type TSelectorDetail struct {
	ItemIdent   string `xml:"ItemIdent,attr"`
	TradeObject string `xml:"TradeObject,attr"`
	ObjectSifr  string `xml:"ObjectSifr,attr"`
	OrderNum    string `xml:"OrderNum,attr"`
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

// RK7Item — упрощённый элемент справочника с динамическими атрибутами в Attrs.
type RK7Item struct {
	Ident      string `xml:"Ident,attr,omitempty"`
	Code       string `xml:"Code,attr,omitempty"`
	Name       string `xml:"Name,attr,omitempty"`
	GuidString string `xml:"GUIDString,attr,omitempty"`
	Attrs      map[string]string
}

func (rkitem *RK7Item) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	attrMap := make(map[string]string, len(start.Attr))
	for _, at := range start.Attr {
		switch at.Name.Local {
		case "Code":
			rkitem.Code = at.Value
		case "Ident":
			rkitem.Ident = at.Value
		case "Name":
			rkitem.Name = at.Value
		case "GUIDString":
			rkitem.GuidString = at.Value
		default:
			attrMap[at.Name.Local] = at.Value
		}
	}
	rkitem.Attrs = attrMap
	return d.Skip()
}

type Waiters struct {
	Waiter []struct {
		ID   string `xml:"ID,attr"`
		Code string `xml:"Code,attr"`
	} `xml:"waiter"`
}

type Visit struct {
	Visit                string  `xml:"visit,attr"`
	Guid                 string  `xml:"guid,attr"`
	Finished             string  `xml:"finished,attr"`
	PersistentComment    string  `xml:"persistentComment,attr"`
	NonPersistentComment string  `xml:"nonPersistentComment,attr"`
	Guests               []Guest `xml:"Guests>Guest"`
	Orders               []Order `xml:"Orders>Order"`
}

type Guest struct {
	GuestLabel string `xml:"guestLabel,attr"`
	CardCode   string `xml:"cardCode,attr"`
	ClientID   string `xml:"clientID,attr"`
	AddressID  string `xml:"addressID,attr"`
	Interface  struct {
		ID   string `xml:"id,attr"`
		Code string `xml:"code,attr"`
		Name string `xml:"name,attr"`
	} `xml:"Interface"`
}

type Order struct {
	Visit         string        `xml:"visit,attr"`
	OrderIdent    string        `xml:"orderIdent,attr"`
	Guid          string        `xml:"guid,attr"`
	OrderName     string        `xml:"orderName,attr"`
	Version       string        `xml:"version,attr"`
	OrderSum      string        `xml:"orderSum,attr"`
	UnpaidSum     string        `xml:"unpaidSum,attr"`
	DiscountSum   string        `xml:"discountSum,attr"`
	TotalPieces   string        `xml:"totalPieces,attr"`
	Paid          string        `xml:"paid,attr"`
	Finished      string        `xml:"finished,attr"`
	OpenTime      string        `xml:"openTime,attr"`
	Waiter        Waiter        `xml:"Waiter"`
	OrderCategory OrderCategory `xml:"OrderCategory"`
	OrderType     OrderType     `xml:"OrderType"`
	Table         Table         `xml:"Table"`
	Station       Station       `xml:"Station"`
}

type CommandResult struct {
	CMD         string `xml:"CMD,attr"`
	Status      string `xml:"Status,attr"`
	ErrorText   string `xml:"ErrorText,attr"`
	DateTime    string `xml:"DateTime,attr"`
	WorkTime    string `xml:"WorkTime,attr"`
	Lastversion string `xml:"lastversion,attr"`
	Data        []interface{}
}

func (cr *CommandResult) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	result := make(map[string]string)
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "CMD":
			cr.CMD = attr.Value
		case "Status":
			cr.Status = attr.Value
		case "ErrorText":
			cr.ErrorText = attr.Value
		case "DateTime":
			cr.DateTime = attr.Value
		case "WorkTime":
			cr.WorkTime = attr.Value
		}
	}

	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		se, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}

		switch se.Name.Local {
		case "RK7Reference":
			ref := RK7Reference{}
			d.DecodeElement(&ref, &se)
			cr.Data = append(cr.Data, ref)
		case "SystemInfo":
			si := SystemInfo{}
			d.DecodeElement(&si, &se)
			cr.Data = append(cr.Data, si)
		case "Waiters":
			waiters := Waiters{}
			d.DecodeElement(&waiters, &se)
			cr.Data = append(cr.Data, waiters)
		case "Visit":
			visit := Visit{}
			d.DecodeElement(&visit, &se)
			cr.Data = append(cr.Data, visit)
		case "RK7RefList":
			refList := RK7RefList{}
			d.DecodeElement(&refList, &se)
			cr.Data = append(cr.Data, refList)
		default:
			for _, attr := range se.Attr {
				result[fmt.Sprintf("%s_%s", se.Name.Local, attr.Name.Local)] = attr.Value
			}
		}
	}

	if len(result) > 0 {
		cr.Data = append(cr.Data, result)
	}
	return nil
}
