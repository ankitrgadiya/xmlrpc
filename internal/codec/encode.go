package codec

import (
	"encoding/base64"
	"encoding/xml"
	"reflect"
	"time"
)

func EncodeData(e *xml.Encoder, data interface{}) (err error) {
	if err = e.EncodeToken(ValueStart); err != nil {
		return err
	}

	switch v := data.(type) {
	case string:
		if err = e.EncodeElement(v, StringStart); err != nil {
			return err
		}
	case int:
		if err = e.EncodeElement(v, IntStart); err != nil {
			return err
		}
	case bool:
		if err = e.EncodeElement(v, BooleanStart); err != nil {
			return err
		}
	case float64:
		if err = e.EncodeElement(v, DoubleStart); err != nil {
			return err
		}
	case time.Time:
		value := v.Format(time.RFC3339)
		if err = e.EncodeElement(value, DateTimeStart); err != nil {
			return err
		}
	case []byte:
		value := base64.StdEncoding.EncodeToString(v)
		if err = e.EncodeElement(value, Base64Start); err != nil {
			return err
		}
	default:
		if err = encodeCompoundType(e, reflect.ValueOf(v)); err != nil {
			return err
		}
	}

	if err = e.EncodeToken(ValueStop); err != nil {
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
		return &xml.UnsupportedTypeError{Type: v.Type()}
	}
}

func encodeArray(e *xml.Encoder, v reflect.Value) (err error) {
	if err = e.EncodeToken(ArrayStart); err != nil {
		return err
	}

	if err = e.EncodeToken(DataStart); err != nil {
		return err
	}

	length := v.Len()
	for i := 0; i < length; i++ {
		if err = EncodeData(e, v.Index(i)); err != nil {
			return err
		}
	}

	if err = e.EncodeToken(DataStop); err != nil {
		return err
	}

	if err = e.EncodeToken(ArrayStop); err != nil {
		return err
	}

	return nil
}

func encodeStruct(e *xml.Encoder, v reflect.Value) (err error) {
	if err = e.EncodeToken(StructStart); err != nil {
		return err
	}

	t := v.Type()
	numFields := v.NumField()
	for i := 0; i < numFields; i++ {
		if err = encodeStructField(e, t.Field(i), v.Field(i)); err != nil {
			return err
		}
	}

	if err = e.EncodeToken(StructStop); err != nil {
		return err
	}

	return nil
}

func encodeStructField(e *xml.Encoder, d reflect.StructField, v reflect.Value) (err error) {
	if err = e.EncodeToken(MemberStart); err != nil {
		return err
	}

	if err = e.EncodeElement(parseName(d), NameStart); err != nil {
		return err
	}

	if err = EncodeData(e, v.Interface()); err != nil {
		return err
	}

	if err = e.EncodeToken(MemberStop); err != nil {
		return err
	}

	return nil
}
