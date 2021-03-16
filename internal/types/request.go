package types

import (
	"encoding/xml"
)

type MethodCall struct {
	XMLName    xml.Name `xml:"methodCall"`
	MethodName string   `xml:"methodName"`
	Params     []Param  `xml:"params>param"`

	paramTypes []interface{} `xml:"-"`
}

func (mc *MethodCall) Encode(indent bool) ([]byte, error) {
	if indent {
		return xml.MarshalIndent(mc, "", "\t")
	}

	return xml.Marshal(mc)
}

func (mc *MethodCall) Decode(_ []byte, paramTypes ...interface{}) error {
	mc.paramTypes = paramTypes
	return nil
}
