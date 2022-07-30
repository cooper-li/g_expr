package g_expr

import (
	"fmt"
	"strconv"
	"strings"
)

func convToInt64(v interface{}) (int64, error) {
	switch v.(type) {
	case int64:
		return v.(int64), nil
	case string:
		vI64, e := strconv.ParseInt(v.(string), 10, 64)
		return vI64, e
	case int:
		return int64(v.(int)), nil
	case int32:
		return int64(v.(int32)), nil
	}
	return int64(0), fmt.Errorf("data: %v convert to int64 err", v)
}

func convToString(v interface{}) (string, error) {
	switch v.(type) {
	case string:
		return v.(string), nil
	case int:
		return strconv.Itoa(v.(int)), nil
	case int32:
		return strconv.Itoa(int(v.(int32))), nil
	case int64:
		return strconv.Itoa(int(v.(int64))), nil
	}
	return "", fmt.Errorf("data: %v convert to string err", v)
}

func trimQuotes(str string) string {
	return strings.Trim(str, "\"")
}
