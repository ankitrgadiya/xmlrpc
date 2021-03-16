package types

import (
	"encoding/xml"

	"argc.in/xmlrpc/internal/codec"
)

type Param struct {
	XMLName xml.Name `xml:"param"`
	Value   Value    `xml:"value"`
}

type Value struct {
	Data interface{}
}

var _ xml.Marshaler = (*Value)(nil)

func (v *Value) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	return codec.EncodeData(e, v.Data)
}
