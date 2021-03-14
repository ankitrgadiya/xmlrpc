package xmlrpc

import "encoding/xml"

const tagName = "xmlrpc"

var (
	memberStart   = xml.StartElement{Name: xml.Name{Local: "member"}}
	memberStop    = xml.EndElement{Name: xml.Name{Local: "member"}}
	structStart   = xml.StartElement{Name: xml.Name{Local: "struct"}}
	structStop    = xml.EndElement{Name: xml.Name{Local: "struct"}}
	arrayStart    = xml.StartElement{Name: xml.Name{Local: "array"}}
	arrayStop     = xml.EndElement{Name: xml.Name{Local: "array"}}
	dataStart     = xml.StartElement{Name: xml.Name{Local: "data"}}
	dataStop      = xml.EndElement{Name: xml.Name{Local: "data"}}
	valueStart    = xml.StartElement{Name: xml.Name{Local: "value"}}
	valueStop     = xml.EndElement{Name: xml.Name{Local: "value"}}
	nameStart     = xml.StartElement{Name: xml.Name{Local: "name"}}
	intStart      = xml.StartElement{Name: xml.Name{Local: "i4"}}
	booleanStart  = xml.StartElement{Name: xml.Name{Local: "boolean"}}
	stringStart   = xml.StartElement{Name: xml.Name{Local: "string"}}
	doubleStart   = xml.StartElement{Name: xml.Name{Local: "double"}}
	dateTimeStart = xml.StartElement{Name: xml.Name{Local: "dateTime.iso8601"}}
	base64Start   = xml.StartElement{Name: xml.Name{Local: "base64"}}
)
