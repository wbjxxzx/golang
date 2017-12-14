package stringutil

import "strconv"

type any interface{}

func Str(a any) string {
	switch v := a.(type) {
	case int:
		return strconv.Itoa(v)
	case int32:
		return string(v)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', 4, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', 4, 64)
	case byte:
		return strconv.Itoa(int(v))
	case bool:
		return strconv.FormatBool(v)
	case string:
		return v
	default:
		return "unknown type"
	}
}
