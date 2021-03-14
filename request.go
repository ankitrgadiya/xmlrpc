package xmlrpc

import "encoding/xml"

func EncodeRequest(r *Request, indent bool) ([]byte, error) {
	if indent {
		return xml.MarshalIndent(r, "", "\t")
	}

	return xml.Marshal(r)
}

type Request struct {
	XMLName    xml.Name `xml:"methodCall"`
	MethodName string   `xml:"methodName"`
	Params     []Param  `xml:"params>param"`
}

type Param struct {
	XMLName xml.Name `xml:"param"`
	Value   Value    `xml:"value"`
}

type Value struct {
	Data interface{}
}

func (v *Value) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	return encodeData(e, v.Data)
}
