package utils

import "strconv"

func GetInt64FromMap(dm map[string]interface{}, key string, dft int64) int64 {
	data, ok := dm[key]
	if !ok {
		return dft
	}
	switch data.(type) {
	case string:
		num, err := strconv.Atoi(data.(string))
		if err != nil {
			return dft
		} else {
			return int64(num)
		}
	case uint:
		return int64(data.(uint))
	case uint8:
		return int64(data.(uint8))
	case uint16:
		return int64(data.(uint16))
	case uint32:
		return int64(data.(uint32))
	case uint64:
		return int64(data.(uint64))
	case int:
		return int64(data.(int))
	case int8:
		return int64(data.(int8))
	case int16:
		return int64(data.(int16))
	case int32:
		return int64(data.(int32))
	case int64:
		return data.(int64)
	case float32:
		return int64(data.(float32))
	case float64:
		return int64(data.(float64))
	}
	return data.(int64)
}
