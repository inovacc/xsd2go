package xsd

import (
	"encoding/xml"
	"strings"

	"github.com/iancoleman/strcase"
)

// Enumeration defines single XML enumeration.
type Enumeration struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema enumeration"`
	Value   string   `xml:"value,attr"`
}

// GoName returns the public Go Name of this struct item.
func (e *Enumeration) GoName() string {
	return strcase.ToCamel(strings.ToLower(e.Value))
}

func (*Enumeration) Modifiers() string {
	return "-"
}

func (e *Enumeration) XmlName() string {
	return e.Value
}
