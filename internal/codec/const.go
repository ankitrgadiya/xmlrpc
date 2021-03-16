package codec

import "encoding/xml"

const TagName = "xmlrpc"

var (
	MemberStart   = xml.StartElement{Name: xml.Name{Local: "member"}}
	MemberStop    = xml.EndElement{Name: xml.Name{Local: "member"}}
	StructStart   = xml.StartElement{Name: xml.Name{Local: "struct"}}
	StructStop    = xml.EndElement{Name: xml.Name{Local: "struct"}}
	ArrayStart    = xml.StartElement{Name: xml.Name{Local: "array"}}
	ArrayStop     = xml.EndElement{Name: xml.Name{Local: "array"}}
	DataStart     = xml.StartElement{Name: xml.Name{Local: "data"}}
	DataStop      = xml.EndElement{Name: xml.Name{Local: "data"}}
	ValueStart    = xml.StartElement{Name: xml.Name{Local: "value"}}
	ValueStop     = xml.EndElement{Name: xml.Name{Local: "value"}}
	FaultStart    = xml.StartElement{Name: xml.Name{Local: "fault"}}
	FaultStop     = xml.EndElement{Name: xml.Name{Local: "fault"}}
	NameStart     = xml.StartElement{Name: xml.Name{Local: "name"}}
	NameStop      = xml.EndElement{Name: xml.Name{Local: "name"}}
	ParamsStart   = xml.StartElement{Name: xml.Name{Local: "params"}}
	ParamsStop    = xml.EndElement{Name: xml.Name{Local: "params"}}
	ParamStart    = xml.StartElement{Name: xml.Name{Local: "param"}}
	ParamStop     = xml.EndElement{Name: xml.Name{Local: "param"}}
	IntStart      = xml.StartElement{Name: xml.Name{Local: "i4"}}
	IntStop       = xml.EndElement{Name: xml.Name{Local: "i4"}}
	BooleanStart  = xml.StartElement{Name: xml.Name{Local: "boolean"}}
	BooleanStop   = xml.EndElement{Name: xml.Name{Local: "boolean"}}
	StringStart   = xml.StartElement{Name: xml.Name{Local: "string"}}
	StringStop    = xml.EndElement{Name: xml.Name{Local: "string"}}
	DoubleStart   = xml.StartElement{Name: xml.Name{Local: "double"}}
	DoubleStop    = xml.EndElement{Name: xml.Name{Local: "double"}}
	DateTimeStart = xml.StartElement{Name: xml.Name{Local: "dateTime.iso8601"}}
	DateTimeStop  = xml.EndElement{Name: xml.Name{Local: "dateTime.iso8601"}}
	Base64Start   = xml.StartElement{Name: xml.Name{Local: "base64"}}
	Base64Stop    = xml.EndElement{Name: xml.Name{Local: "base64"}}
)
