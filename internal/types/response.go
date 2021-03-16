package types

import (
	"encoding/xml"

	"argc.in/xmlrpc/internal/codec"
	"argc.in/xmlrpc/internal/errors"
)

type MethodResponse struct {
	XMLName xml.Name `xml:"methodResponse"`
	Fault   *Fault   `xml:"fault,omitempty"`
	Params  []Param  `xml:"params>param,omitempty"`

	// paramType is used to store the user-given datatype for parsing the
	// Parameter.
	paramType interface{} `xml:"-"`
}

func (mr *MethodResponse) Encode(indent bool) ([]byte, error) {
	// MethodResponse must have either Fault or Params. Also, MethodResponse can
	// only have a single Parameter under Params.
	if (mr.Fault != nil && len(mr.Params) != 0) || (len(mr.Params) > 1) {
		return nil, errors.NewInvalidResponseError("both fault and param is present")
	}

	if indent {
		return xml.MarshalIndent(mr, "", "\t")
	}

	return xml.Marshal(mr)
}

func (mr *MethodResponse) Decode(data []byte, paramType interface{}) error {
	if paramType == nil {
		return codec.ErrInvalidParamCount
	}

	mr.paramType = paramType
	return xml.Unmarshal(data, mr)
}

type Fault struct {
	XMLName     xml.Name `xml:"fault"`
	FaultCode   int      `xmlrpc:"faultCode"`
	FaultString string   `xmlrpc:"faultString"`
}

var _ xml.Marshaler = (*Fault)(nil)

// TODO(ankit): This might not be required, validate it.
func (f *Fault) MarshalXML(e *xml.Encoder, _ xml.StartElement) (err error) {
	if err = e.EncodeToken(codec.FaultStart); err != nil {
		return err
	}

	if err = codec.EncodeData(e, f); err != nil {
		return err
	}

	if err = e.EncodeToken(codec.FaultStop); err != nil {
		return err
	}

	return nil
}
