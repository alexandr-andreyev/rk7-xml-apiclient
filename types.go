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
	RegisteredOnly string         `xml:"registeredOnly,attr,omitempty"`
	PROPFILTERS    []PROPFILTER   `xml:"PROPFILTERS>PROPFILTER,omitempty"`
}

type PROPFILTER struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

type Station struct {
	Code string `xml:"Code,attr"`
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
	DataVersion    string `xml:"DataVersion,attr,omitempty"`
	ClassName      string `xml:"ClassName,attr,omitempty"`
	TotalItemCount string `xml:"TotalItemCount,attr,omitempty"`
	Count          string `xml:"Count,attr,omitempty"`
	Items          struct {
		Item []struct {
			Ident string `xml:"Ident,attr"`
			Code  string `xml:"Code,attr"`
			Name  string `xml:"Name,attr"`
			Attrs []xml.Attr
		} `xml:"Item"`
	} `xml:"Items"`
}

// WaiterList req
type Waiters struct {
	Waiter []struct {
		ID   string `xml:"ID,attr"`
		Code string `xml:"Code,attr"`
	} `xml:"waiter"`
}

type CommandResult struct {
	CMD       string `xml:"CMD,attr"`
	Status    string `xml:"Status,attr"`
	ErrorText string `xml:"ErrorText,attr"`
	DateTime  string `xml:"DateTime,attr"`
	WorkTime  string `xml:"WorkTime,attr"`
	Data      []interface{}
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
