package convert

import (
	"encoding/json"
	"strconv"
)

// This files is dedicated to all converter functions for primitives types. For non-primitives types, make a new
// files inside package convert

// ToFloat64 Convert any value to float64
func ToFloat64(v interface{}) float64 {
	result := float64(0)
	switch v.(type) {
	case string:
		result, _ = strconv.ParseFloat(v.(string), 64)
	case int:
		result = float64(v.(int))
	case int64:
		result = float64(v.(int64))
	case float64:
		result = float64(v.(float64))
	case uint8:
		result, _ = strconv.ParseFloat(string(v.(uint8)), 64)
	default:
		result = float64(0)
	}

	return result
}

// ToInt Convert any value to int
func ToInt(v interface{}) int {
	result := int(0)
	switch v.(type) {
	case string:
		result, _ = strconv.Atoi(v.(string))
	case int:
		result = int(v.(int))
	case int64:
		result = int(v.(int64))
	case float64:
		result = int(v.(float64))
	case uint8:
		result, _ = strconv.Atoi(string(v.(uint8)))
	case []uint8:
		result, _ = strconv.Atoi(string(v.([]uint8)))
	default:
		result = int(0)
	}

	return result
}

// ToInt64 Convert any value to int64
func ToInt64(v interface{}) int64 {
	result := int64(0)
	switch v.(type) {
	case string:
		result, _ = strconv.ParseInt(v.(string), 10, 64)
	case int:
		result = int64(v.(int))
	case int64:
		result = int64(v.(int64))
	case float64:
		result = int64(v.(float64))
	case uint8:
		result, _ = strconv.ParseInt(string(v.(uint8)), 10, 64)
	case []uint8:
		result, _ = strconv.ParseInt(string(v.([]uint8)), 10, 64)
	default:
		result = int64(0)
	}

	return result
}

// ToString Convert any value to string
func ToString(v interface{}) string {
	result := ""
	if v == nil {
		return ""
	}
	switch v.(type) {
	case string:
		result = v.(string)
	case int:
		result = strconv.Itoa(v.(int))
	case int64:
		result = strconv.FormatInt(v.(int64), 10)
	case bool:
		result = strconv.FormatBool(v.(bool))
	case float64:
		result = strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case uint8:
		result = string(v.(uint8))
	case []uint8:
		result = string(v.([]uint8))
	default:
		resultJSON, err := json.Marshal(v)
		if err == nil {
			result = string(resultJSON)
		} else {
			// failed to convert to string
			return ""
		}
	}

	return result
}
