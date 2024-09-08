package array

import (
	"fmt"
	"strconv"
	"strings"
)

type DataType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

// Join 将数组使用指定分隔符拼接成字符串
func Join[T DataType](separator string, arr []T) string {
	if len(arr) == 0 || len(separator) == 0 {
		return ""
	}
	var builder strings.Builder
	for index, item := range arr {
		builder.WriteString(fmt.Sprint(item))
		if len(arr)-1 > index {
			builder.WriteString(separator)
		}
	}
	return builder.String()
}

// Split 函数按分隔符将字符串分割成指定类型的切片
func Split[T DataType](separator string, str string) []T {
	if len(str) == 0 || len(separator) == 0 {
		return nil
	}

	parts := strings.Split(str, separator)
	result := make([]T, 0, len(parts))
	for _, part := range parts {
		var value T
		switch any(value).(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			num, err := strconv.ParseFloat(part, 64)
			if err != nil {
				return nil
			}
			switch any(value).(type) {
			case int, int8, int16, int32, int64:
				value = any(int(num)).(T)
			case uint, uint8, uint16, uint32, uint64:
				value = any(uint(num)).(T)
			case float32:
				value = any(float32(num)).(T)
			case float64:
				value = any(num).(T)
			}
		case string:
			value = any(part).(T)
		default:
			return nil
		}
		result = append(result, value)
	}

	return result
}
