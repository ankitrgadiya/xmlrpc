package xmlrpc

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"reflect"
	"time"
)

func encodeData(e *xml.Encoder, data interface{}) (err error) {
	if err = e.EncodeToken(valueStart); err != nil {
		return err
	}

	switch v := data.(type) {
	case string:
		if err = e.EncodeElement(v, stringStart); err != nil {
			return err
		}
	case int:
		if err = e.EncodeElement(v, intStart); err != nil {
			return err
		}
	case bool:
		if err = e.EncodeElement(v, booleanStart); err != nil {
			return err
		}
	case float64:
		if err = e.EncodeElement(v, doubleStart); err != nil {
			return err
		}
	case time.Time:
		value := v.Format(time.RFC3339)
		if err = e.EncodeElement(value, dateTimeStart); err != nil {
			return err
		}
	case []byte:
		value := base64.StdEncoding.EncodeToString(v)
		if err = e.EncodeElement(value, base64Start); err != nil {
			return err
		}
	default:
		if err = encodeCompoundType(e, reflect.ValueOf(v)); err != nil {
			return err
		}
	}

	if err = e.EncodeToken(valueStop); err != nil {
		return err
	}

	return nil
}

func encodeCompoundType(e *xml.Encoder, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		return encodeCompoundType(e, v.Elem())
	case reflect.Array, reflect.Slice:
		return encodeArray(e, v)
	case reflect.Struct:
		return encodeStruct(e, v)
	default:
		return errors.New("not supported")
	}
}

func encodeArray(e *xml.Encoder, v reflect.Value) (err error) {
	if err = e.EncodeToken(arrayStart); err != nil {
		return err
	}

	if err = e.EncodeToken(dataStart); err != nil {
		return err
	}

	length := v.Len()
	for i := 0; i < length; i++ {
		if err = encodeData(e, v.Index(i)); err != nil {
			return err
		}
	}

	if err = e.EncodeToken(dataStop); err != nil {
		return err
	}

	if err = e.EncodeToken(arrayStop); err != nil {
		return err
	}

	return nil
}

func encodeStruct(e *xml.Encoder, v reflect.Value) (err error) {
	if err = e.EncodeToken(structStart); err != nil {
		return err
	}

	t := v.Type()
	numFields := v.NumField()
	for i := 0; i < numFields; i++ {
		if err = encodeStructField(e, t.Field(i), v.Field(i)); err != nil {
			return err
		}
	}

	if err = e.EncodeToken(structStop); err != nil {
		return err
	}

	return nil
}

func encodeStructField(e *xml.Encoder, d reflect.StructField, v reflect.Value) (err error) {
	if err = e.EncodeToken(memberStart); err != nil {
		return err
	}

	if err = e.EncodeElement(parseName(d), nameStart); err != nil {
		return err
	}

	if err = encodeData(e, v.Interface()); err != nil {
		return err
	}

	if err = e.EncodeToken(memberStop); err != nil {
		return err
	}

	return nil
}

func parseName(f reflect.StructField) string {
	name := f.Tag.Get(tagName)
	if name == "" {
		name = f.Name
	}

	return name
}
