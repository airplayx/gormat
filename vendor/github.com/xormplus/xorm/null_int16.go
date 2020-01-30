package xorm

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type NullInt16 struct {
	Int16 int16
	Valid bool
}

func (ni NullInt16) Ptr() *int16 {
	if !ni.Valid {
		return nil
	}
	return &ni.Int16
}

func (ni NullInt16) ValueOrZero() int16 {
	if !ni.Valid {
		return 0
	}
	return ni.Int16
}

func (ni NullInt16) IsNil() bool {
	return !ni.Valid
}

func (ni *NullInt16) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case float64:
		err = json.Unmarshal(data, &ni.Int16)
	case string:
		str := string(x)
		if len(str) == 0 {
			ni.Valid = false
			return nil
		}
		var i int64
		i, err = strconv.ParseInt(str, 10, 16)
		if err == nil {
			ni.Int16 = int16(i)
		}
	case map[string]interface{}:
		err = json.Unmarshal(data, &ni)
	case nil:
		ni.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type xorm.NullInt16", reflect.TypeOf(v).Name())
	}
	ni.Valid = err == nil
	return err
}

func (ni *NullInt16) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		ni.Valid = false
		return nil
	}
	var err error
	var i int64
	i, err = strconv.ParseInt(string(text), 10, 16)
	if err == nil {
		ni.Int16 = int16(i)
	}
	ni.Valid = err == nil
	return err
}

func (ni NullInt16) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(int64(ni.Int16), 10)), nil
}

func (ni NullInt16) MarshalText() ([]byte, error) {
	if !ni.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatInt(int64(ni.Int16), 10)), nil
}
