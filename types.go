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
	OnlyActrive    onlyactive     `xml:"OnlyActive,attr,omitempty"`
	WithChildItems withChildItems `xml:"WithChildItems,attr,omitempty"`
	WithMacroProp  withMacroProp  `xml:"WithMacroProp,attr,omitempty"`
	PropMask       string         `xml:"PropMask,attr,omitempty"`
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
	//XMLName         xml.Name        `xml:"RK7QueryResult"`
	ServerVersion   string          `xml:"ServerVersion,attr"`
	XmlVersion      string          `xml:"XmlVersion,attr"`
	NetName         string          `xml:"NetName,attr"`
	Status          string          `xml:"Status,attr"`
	Processed       string          `xml:"Processed,attr"`
	ArrivalDateTime string          `xml:"ArrivalDateTime,attr"`
	CommandResult   []CommandResult `xml:"CommandResult"`
}

type RK7Reference struct {
	DataVersion    string    `xml:"DataVersion,attr,omitempty"`
	ClassName      string    `xml:"ClassName,attr,omitempty"`
	TotalItemCount string    `xml:"TotalItemCount,attr,omitempty"`
	Count          string    `xml:"Count,attr,omitempty"`
	Items          []RK7Item `xml:"Items>Item,omitempty"`
}

type RK7Item struct {
	Ident      string `xml:"Ident,attr,omitempty"`
	Code       string `xml:"Code,attr,omitempty"`
	Name       string `xml:"Name,attr,omitempty"`
	GuidString string `xml:"GUIDString,attr,omitempty"`
	Attrs      map[string]string
}

func (rkitem *RK7Item) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	fmt.Println("start name>", start.Name.Local)
	startAttrs := start.Attr
	attrMap := make(map[string]string, len(startAttrs))
	for _, at := range startAttrs {
		fmt.Println("Start atr", at.Name.Local, at.Value)
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

// WaiterList req
type Waiters struct {
	Waiter []struct {
		ID   string `xml:"ID,attr"`
		Code string `xml:"Code,attr"`
	} `xml:"waiter"`
}

type Visit struct {
	Visit                string `xml:"visit,attr"`
	Guid                 string `xml:"guid,attr"`
	Finished             string `xml:"finished,attr"`
	PersistentComment    string `xml:"persistentComment,attr"`
	NonPersistentComment string `xml:"nonPersistentComment,attr"`
	Orders               struct {
		Order struct {
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
	}
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
	result := make(map[string]string, 0)
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

		switch tok.(type) {
		case xml.StartElement:
			nextStart := tok.(xml.StartElement)
			attrs := nextStart.Attr

			if nextStart.Name.Local == "RK7Reference" {
				ref := RK7Reference{}
				d.DecodeElement(&ref, &nextStart)

				cr.Data = append(cr.Data, ref)
				break
			} else if nextStart.Name.Local == "Waiters" {
				waiters := Waiters{}
				d.DecodeElement(&waiters, &nextStart)
				cr.Data = append(cr.Data, waiters)
				break
			} else if nextStart.Name.Local == "Visit" {
				visits := Visit{}
				d.DecodeElement(&visits, &nextStart)
				cr.Data = append(cr.Data, visits)
				break
			} else {

				for _, attr := range attrs {
					//fmt.Printf("attr %s = %s\n", attr.Name.Local, attr.Value)
					attrName := fmt.Sprintf("%s_%s", nextStart.Name.Local, attr.Name.Local)
					result[attrName] = attr.Value
				}
			}
		}

	}
	if len(result) > 0 {
		cr.Data = append(cr.Data, result)
	}
	return nil
}
