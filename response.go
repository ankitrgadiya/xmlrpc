package xmlrpc

import "encoding/xml"

func EncodeResponse(r *Response, indent bool) ([]byte, error) {
	if (r.Fault != nil && len(r.Params) != 0) || (len(r.Params) > 1) {
		return nil, ErrInvalidResponse
	}

	if indent {
		return xml.MarshalIndent(r, "", "\t")
	}

	return xml.Marshal(r)
}

type Response struct {
	XMLName xml.Name `xml:"methodResponse"`
	Fault   *Fault   `xml:"fault,omitempty"`
	Params  []Param  `xml:"params>param"`
}

type Fault struct {
	FaultCode   int    `xmlrpc:"faultCode"`
	FaultString string `xmlrpc:"faultString"`
}

func (f *Fault) MarshalXML(e *xml.Encoder, _ xml.StartElement) (err error) {
	if err = e.EncodeToken(faultStart); err != nil {
		return err
	}

	if err = encodeData(e, f); err != nil {
		return err
	}

	if err = e.EncodeToken(faultStop); err != nil {
		return err
	}

	return nil
}
