package codec

import (
	"encoding/xml"
	"io"
	"reflect"

	"argc.in/xmlrpc/internal/errors"
)

func parseName(f reflect.StructField) string {
	name := f.Tag.Get(TagName)
	if name == "" {
		name = f.Name
	}

	return name
}

func skipUntilStartToken(d xml.TokenReader, hint *xml.StartElement) (*xml.StartElement, error) {
	for {
		t, err := d.Token()
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}

		if errors.Is(err, io.EOF) {
			return nil, nil
		}

		// Juicy Part!!
		start, ok := t.(xml.StartElement)
		if !ok {
			continue
		}

		if hint != nil && hint.Name.Local != start.Name.Local {
			continue
		}

		return &start, nil
	}
}

func structFieldByName(name string, s reflect.Value) (reflect.Value, error) {
	s = s.Elem()
	t := s.Type()
	totalFields := s.NumField()

	for i := 0; i < totalFields; i++ {
		if parseName(t.Field(i)) == name {
			return s.Field(i), nil
		}
	}

	return reflect.Value{}, errors.NewInvalidParamError("field not present")
}
