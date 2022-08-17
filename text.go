package text

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Unmarshal(txt string, v interface{}) error {
	switch v := v.(type) {
	case *string:
		*v = txt
	case *bool:
		b, err := strconv.ParseBool(txt)
		if err != nil {
			return err
		}
		*v = b
	case *int:
		i, err := strconv.ParseInt(txt, 10, strconv.IntSize)
		if err != nil {
			return err
		}
		*v = int(i)
	case *int8:
		i, err := strconv.ParseInt(txt, 10, 8)
		if err != nil {
			return err
		}
		*v = int8(i)
	case *int16:
		i, err := strconv.ParseInt(txt, 10, 16)
		if err != nil {
			return err
		}
		*v = int16(i)
	case *int32:
		i, err := strconv.ParseInt(txt, 10, 32)
		if err != nil {
			return err
		}
		*v = int32(i)
	case *int64:
		i, err := strconv.ParseInt(txt, 10, 64)
		if err != nil {
			return err
		}
		*v = i
	case *uint:
		i, err := strconv.ParseUint(txt, 10, strconv.IntSize)
		if err != nil {
			return err
		}
		*v = uint(i)
	case *uint8:
		i, err := strconv.ParseUint(txt, 10, 8)
		if err != nil {
			return err
		}
		*v = uint8(i)
	case *uint16:
		i, err := strconv.ParseUint(txt, 10, 16)
		if err != nil {
			return err
		}
		*v = uint16(i)
	case *uint32:
		i, err := strconv.ParseUint(txt, 10, 32)
		if err != nil {
			return err
		}
		*v = uint32(i)
	case *uint64:
		i, err := strconv.ParseUint(txt, 10, 64)
		if err != nil {
			return err
		}
		*v = i
	case *float32:
		f, err := strconv.ParseFloat(txt, 32)
		if err != nil {
			return err
		}
		*v = float32(f)
	case *float64:
		f, err := strconv.ParseFloat(txt, 64)
		if err != nil {
			return err
		}
		*v = f
	case *complex64:
		c, err := strconv.ParseComplex(txt, 64)
		if err != nil {
			return err
		}
		*v = complex64(c)
	case *complex128:
		c, err := strconv.ParseComplex(txt, 128)
		if err != nil {
			return err
		}
		*v = c
	case *time.Duration:
		d, err := time.ParseDuration(txt)
		if err != nil {
			return err
		}
		*v = d
	case encoding.TextUnmarshaler:
		return v.UnmarshalText([]byte(txt))
	default:
		t := reflect.TypeOf(v)
		if t.Kind() != reflect.Ptr {
			return errors.New("v must be a pointer")
		}
		switch et := t.Elem(); et.Kind() {
		case reflect.Slice:
			a := strings.Split(txt, ",")
			s := reflect.MakeSlice(et, len(a), len(a))
			for i := range a {
				if err := Unmarshal(a[i], s.Index(i).Addr().Interface()); err != nil {
					return err
				}
			}
			reflect.ValueOf(v).Elem().Set(s)
		case reflect.Map:
			a := strings.Split(txt, ",")
			m := reflect.MakeMap(et)
			kt := et.Key()
			vt := et.Elem()
			for i := range a {
				pair := strings.Split(a[i], ":")
				if len(pair) != 2 {
					return fmt.Errorf("invalid map: %s", txt)
				}
				kv := reflect.New(kt)
				if err := Unmarshal(pair[0], kv.Interface()); err != nil {
					return err
				}
				vv := reflect.New(vt)
				if err := Unmarshal(pair[1], vv.Interface()); err != nil {
					return err
				}
				m.SetMapIndex(kv.Elem(), vv.Elem())
			}
			reflect.ValueOf(v).Elem().Set(m)
		default:
			return fmt.Errorf("unsupported type: %s", t)
		}
	}
	return nil
}
