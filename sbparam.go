package sbparam

import (
	"fmt"
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"math"
	"reflect"
	"strings"
)

// Unmarshal parses data.Map to given structure.
//
// TODO: write tag spec
func Unmarshal(dat data.Map, v interface{}) error {
	// TODO: recover logging

	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("'%v' type is not supported", rt)
	}

	for i := 0; i < rv.Elem().NumField(); i++ {
		field := rt.Elem().Field(i)
		structValue := rv.Elem().Field(i)

		keyName := field.Name
		required := true
		defValue := "" // need to cast when set to arg `v`
		if tag, ok := field.Tag.Lookup("sbparam"); ok {
			paramTag := strings.Split(tag, ",")
			if len(paramTag) > 0 && paramTag[0] != "" {
				keyName = paramTag[0]
			}
			if len(paramTag) > 1 && paramTag[1] == "omitempty" {
				required = false
			}
			if len(paramTag) > 2 {
				required = false
				defValue = paramTag[2]
			}
		}

		keyPath, err := data.CompilePath(keyName)
		if err != nil {
			return fmt.Errorf("cannot compile '%v' as data.Path", keyName)
		}
		dv, dataPathErr := dat.Get(keyPath)
		if dataPathErr != nil {
			if required {
				return fmt.Errorf("key '%v' is not found in param", keyName)
			}
		}
		ps := paramSet{
			key:      keyName,
			required: required,
			defValue: defValue,
		}
		switch field.Type.Kind() {
		case reflect.String:
			value, err := decodeString(dv, ps)
			if err != nil {
				return err
			}
			ptr := structValue.Addr().Interface().(*string)
			*ptr = value

		case reflect.Int:
			value, err := decodeInt(dv, ps)
			if err != nil {
				return err
			}
			ptr := structValue.Addr().Interface().(*int)
			*ptr = int(value)
		case reflect.Int8:
			value, err := decodeInt(dv, ps)
			if err != nil {
				return err
			}
			if value > math.MaxInt8 || value < math.MinInt8 {
				return fmt.Errorf("'%v'(int8) overflow error: %d", keyName,
					value)
			}
			ptr := structValue.Addr().Interface().(*int8)
			*ptr = int8(value)
		case reflect.Int16:
			value, err := decodeInt(dv, ps)
			if err != nil {
				return err
			}
			if value > math.MaxInt16 || value < math.MinInt16 {
				return fmt.Errorf("'%v'(int16) overflow error: %d", keyName,
					value)
			}
			ptr := structValue.Addr().Interface().(*int16)
			*ptr = int16(value)
		case reflect.Int32:
			value, err := decodeInt(dv, ps)
			if err != nil {
				return err
			}
			if value > math.MaxInt32 || value < math.MinInt32 {
				return fmt.Errorf("'%v'(int32) overflow error: %d", keyName,
					value)
			}
			ptr := structValue.Addr().Interface().(*int32)
			*ptr = int32(value)
		case reflect.Int64:
			value, err := decodeInt(dv, ps)
			if err != nil {
				return err
			}
			ptr := structValue.Addr().Interface().(*int64)
			*ptr = value

		case reflect.Uint:
			value, err := decodeUint(dv, ps)
			if err != nil {
				return err
			}
			ptr := structValue.Addr().Interface().(*uint)
			*ptr = uint(value)
		case reflect.Uint8:
			value, err := decodeUint(dv, ps)
			if err != nil {
				return err
			}
			if value > math.MaxUint8 {
				return fmt.Errorf("'%v'(uint8) overflow error: %d", keyName,
					value)
			}
			ptr := structValue.Addr().Interface().(*uint8)
			*ptr = uint8(value)
		case reflect.Uint16:
			value, err := decodeUint(dv, ps)
			if err != nil {
				return err
			}
			if value > math.MaxUint16 {
				return fmt.Errorf("'%v'(uint16) overflow error: %d", keyName,
					value)
			}
			ptr := structValue.Addr().Interface().(*uint16)
			*ptr = uint16(value)
		case reflect.Uint32:
			value, err := decodeUint(dv, ps)
			if err != nil {
				return err
			}
			if value > math.MaxUint32 {
				return fmt.Errorf("'%v'(uint32) overflow error: %d", keyName,
					value)
			}
			ptr := structValue.Addr().Interface().(*uint32)
			*ptr = uint32(value)
		case reflect.Uint64:
			value, err := decodeUint(dv, ps)
			if err != nil {
				return err
			}
			ptr := structValue.Addr().Interface().(*uint64)
			*ptr = uint64(value)

		case reflect.Float32:
			value, err := decodeFloat(dv, ps)
			if err != nil {
				return err
			}
			if math.Abs(value) > math.MaxFloat32 {
				return fmt.Errorf("'%v'(float32) overflow error: %f", keyName,
					value)
			}
			ptr := structValue.Addr().Interface().(*float32)
			*ptr = float32(value)
		case reflect.Float64:
			value, err := decodeFloat(dv, ps)
			if err != nil {
				return err
			}
			ptr := structValue.Addr().Interface().(*float64)
			*ptr = value

		case reflect.Bool:
			value, err := decodeBool(dv, ps)
			if err != nil {
				return err
			}
			ptr := structValue.Addr().Interface().(*bool)
			*ptr = value

		}
	}
	return nil
}
