package sbparam

import (
	"fmt"
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"strconv"
)

type paramSet struct {
	key      string
	required bool
	defValue string
}

func decodeString(dv data.Value, ps paramSet) (string, error) {
	var value string
	if dv == nil {
		value = ps.defValue
	} else if dataValue, err := data.AsString(dv); err == nil {
		value = dataValue
	} else {
		if ps.required {
			return "", fmt.Errorf("type mismatch, key '%v' is not 'string'",
				ps.key)
		}
		value = ps.defValue
	}
	return value, nil
}

// decodeInt return int value get from "dv" or default value, always return
// as int64. SensorBee supports only int64 type and this function also
// supports only int64.
func decodeInt(dv data.Value, ps paramSet) (int64, error) {
	var value int64
	if dv == nil {
		defInt, err := strconv.ParseInt(ps.defValue, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("default value is not 'int': %v", err)
		}
		value = defInt
	} else if dataValue, err := data.AsInt(dv); err == nil {
		value = dataValue
	} else {
		if ps.required {
			return 0, fmt.Errorf("type mismatch, key '%v' is not 'int'",
				ps.key)
		}
		defInt, err := strconv.ParseInt(ps.defValue, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("default value is not 'int': %v", err)
		}
		value = defInt
	}
	return value, nil
}
