package sbparam

import (
	"fmt"
	"gopkg.in/sensorbee/sensorbee.v0/data"
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
		switch field.Type.Kind() {
		case reflect.String:
			var value string
			if dataPathErr != nil {
				value = defValue
			} else if dataValue, err := data.AsString(dv); err == nil {
				value = dataValue
			} else {
				if required {
					return fmt.Errorf("type mismatch, key '%v' is not '%v'",
						keyName, field.Type)
				}
				value = defValue
			}
			ptr := structValue.Addr().Interface().(*string)
			*ptr = value
		}
	}
	return nil
}
