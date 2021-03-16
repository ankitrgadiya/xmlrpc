package codec

import (
	"encoding/base64"
	"encoding/xml"
	"reflect"
	"time"

	"argc.in/xmlrpc/internal/errors"
)

func DecodeData(d *xml.Decoder, paramType interface{}) (interface{}, error) {
	start, err := skipUntilStartToken(d, nil)
	if err != nil {
		return nil, err
	}

	if start == nil {
		return nil, nil
	}

	switch start.Name.Local {
	case IntStart.Name.Local:
		if _, ok := paramType.(*int); !ok {
			return nil, errors.NewInvalidParamError("not of type int")
		}

		return decodePrimitives(d, start, paramType)
	case BooleanStart.Name.Local:
		if _, ok := paramType.(*bool); !ok {
			return nil, errors.NewInvalidParamError("not of type bool")
		}

		return decodePrimitives(d, start, paramType)
	case StringStart.Name.Local:
		if _, ok := paramType.(*string); !ok {
			return nil, errors.NewInvalidParamError("not of type string")
		}

		return decodePrimitives(d, start, paramType)
	case DoubleStart.Name.Local:
		if _, ok := paramType.(*float64); !ok {
			return nil, errors.NewInvalidParamError("not of type float64")
		}

		return decodePrimitives(d, start, paramType)
	case Base64Start.Name.Local:
		_, ok := paramType.(*[]byte)
		if !ok {
			return nil, errors.NewInvalidParamError("not of type []byte")
		}

		return decodeBase64(d, start, paramType)
	case DateTimeStart.Name.Local:
		_, ok := paramType.(*time.Time)
		if !ok {
			return nil, errors.NewInvalidParamError("not of type *time.Time")
		}

		return decodeDateTime(d, start, paramType)
	case ArrayStart.Name.Local:
		return decodeArray(d, paramType)
	case StructStart.Name.Local:
		return decodeStruct(d, paramType)
	default:
		return nil, errors.NewInvalidStartTokenError(start)
	}
}

func decodePrimitives(d *xml.Decoder, start *xml.StartElement, paramType interface{}) (interface{}, error) {
	v := reflect.ValueOf(paramType)
	if v.Kind() != reflect.Ptr {
		return nil, errors.ErrNotAPointer
	}

	// Parse the XML data in a new instance of Parameter Type.
	zero := reflect.New(v.Elem().Type())
	if err := d.DecodeElement(zero.Interface(), start); err != nil {
		return nil, err
	}

	return zero.Elem().Interface(), nil
}

func decodeBase64(d *xml.Decoder, start *xml.StartElement, _ interface{}) (interface{}, error) {
	var s string
	if err := d.DecodeElement(&s, start); err != nil {
		return nil, err
	}

	value, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func decodeDateTime(d *xml.Decoder, start *xml.StartElement, _ interface{}) (interface{}, error) {
	var s string
	if err := d.DecodeElement(&s, start); err != nil {
		return nil, err
	}

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func decodeArray(d *xml.Decoder, paramType interface{}) (interface{}, error) {
	t := reflect.TypeOf(paramType)
	if t.Kind() != reflect.Ptr {
		return nil, errors.ErrNotAPointer
	}

	t = t.Elem()
	if t.Kind() != reflect.Array && t.Kind() != reflect.Slice {
		return nil, errors.NewInvalidParamError("not of type pointer to slice/array")
	}

	t = t.Elem()
	slice := reflect.MakeSlice(reflect.SliceOf(t), 0, 0)

	for {
		_, err := skipUntilStartToken(d, &ValueStart)
		if err != nil {
			return nil, err
		}

		value, err := DecodeData(d, reflect.New(t).Interface())
		if err != nil {
			return nil, err
		}

		if value == nil {
			break
		}

		slice = reflect.Append(slice, reflect.ValueOf(value))
	}

	return slice.Interface(), nil
}

func decodeStruct(d *xml.Decoder, paramType interface{}) (interface{}, error) {
	t := reflect.TypeOf(paramType)
	if t.Kind() != reflect.Ptr {
		return nil, errors.ErrNotAPointer
	}

	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return nil, errors.NewInvalidParamError("not of type pointer to struct")
	}

	// Allocate new struct of the given type.
	s := reflect.New(t)

	for {
		start, err := skipUntilStartToken(d, &NameStart)
		if err != nil {
			return nil, err
		}

		if start == nil {
			break
		}

		var name string

		err = d.DecodeElement(&name, start)
		if err != nil {
			return nil, err
		}

		if err = decodeStructField(d, s, name); err != nil {
			return nil, err
		}
	}

	return s.Interface(), nil
}

func decodeStructField(d *xml.Decoder, s reflect.Value, fieldName string) error {
	field, err := structFieldByName(fieldName, s)
	if err != nil {
		return err
	}

	if field.Kind() == reflect.Ptr {
		return errors.NewInvalidParamError("struct field is a pointer")
	}

	addr := field
	if field.CanAddr() {
		addr = field.Addr()
	}

	_, err = skipUntilStartToken(d, &ValueStart)
	if err != nil {
		return err
	}

	value, err := DecodeData(d, addr.Interface())
	if err != nil {
		return err
	}

	if value == nil {
		return nil
	}

	field.Set(reflect.ValueOf(value))

	return nil
}
