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
			return 0, fmt.Errorf("type mismatch, key '%v' is not 'int'", ps.key)
		}
		defInt, err := strconv.ParseInt(ps.defValue, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("default value is not 'int': %v", err)
		}
		value = defInt
	}
	return value, nil
}

func decodeUint(dv data.Value, ps paramSet) (uint64, error) {
	var value uint64
	if dv == nil {
		defUint, err := strconv.ParseUint(ps.defValue, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("default value is not 'uint': %v", err)
		}
		value = defUint
	} else if dataValue, err := data.AsInt(dv); err == nil {
		if dataValue < 0 {
			return 0, fmt.Errorf("value is not 'uint': %d", dataValue)
		}
		value = uint64(dataValue)
	} else {
		if ps.required {
			return 0, fmt.Errorf("type mismatch, key '%v' is not 'int'", ps.key)
		}
		defUint, err := strconv.ParseUint(ps.defValue, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("default value is not 'uint': %v", err)
		}
		value = defUint
	}
	return value, nil
}

func decodeFloat(dv data.Value, ps paramSet) (float64, error) {
	var value float64
	if dv == nil {
		defFloat, err := strconv.ParseFloat(ps.defValue, 64)
		if err != nil {
			return 0.0, fmt.Errorf("default value is not 'float': %v", err)
		}
		value = defFloat
	} else if dataValue, err := data.AsFloat(dv); err == nil {
		value = dataValue
	} else {
		if ps.required {
			return 0.0, fmt.Errorf("type mismatch, key '%v' is not 'float'", ps.key)
		}
		defFloat, err := strconv.ParseFloat(ps.defValue, 64)
		if err != nil {
			return 0.0, fmt.Errorf("default value is not 'float': %v", err)
		}
		value = defFloat
	}
	return value, nil
}
