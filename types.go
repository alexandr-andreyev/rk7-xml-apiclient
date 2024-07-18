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
	CMD            string   `xml:"CMD,attr"`
	RefName        string   `xml:"RefName,attr,omitempty"`
	OnlyActrive    string   `xml:"OnlyActive,attr,omitempty"`
	WithChildItems string   `xml:"WithChildItems,attr,omitempty"`
	WithMacroProp  string   `xml:"WithMacroProp,attr,omitempty"`
	PropMask       string   `xml:"PropMask,attr,omitempty"`
	Station        *Station `xml:"Station,omitempty"`
	RegisteredOnly string   `xml:"registeredOnly,attr,omitempty"`
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

type CommandResult struct {
	CMD       string `xml:"CMD,attr"`
	Status    string `xml:"Status,attr"`
	ErrorText string `xml:"ErrorText,attr"`
	DateTime  string `xml:"DateTime,attr"`
	WorkTime  string `xml:"WorkTime,attr"`
	Data      map[string]string
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

			fmt.Println(nextStart.Name.Local)
			for _, attr := range attrs {
				fmt.Printf("attr %s = %s\n", attr.Name.Local, attr.Value)
				attrName := fmt.Sprintf("%s_%s", nextStart.Name.Local, attr.Name.Local)
				result[attrName] = attr.Value
			}
		}
		cr.Data = result
	}
	return nil
}
