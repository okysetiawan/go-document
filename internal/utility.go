package internal

import (
	"fmt"
	"reflect"
	"strconv"
)

func SliceAnyToString(slice []any) []string {
	var (
		res []string
	)

	for _, each := range slice {
		res = append(res, AnyToString(each))
	}

	return res
}

func AnyToString(toString any) string {
	if s, ok := toString.(fmt.Stringer); ok {
		return s.String()
	}

	var (
		to = reflect.TypeOf(toString)
		vo = reflect.ValueOf(toString)
	)

	switch to.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(vo.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(vo.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(vo.Float(), 'f', 5, 64)
	case reflect.String:
		return vo.String()
	default:
		return fmt.Sprintf("%v", vo)
	}
}

func EqualsAnyRune(actual rune, expects []rune) bool {
	for i := range expects {
		if expects[i] == actual {
			return true
		}
	}

	return false
}
