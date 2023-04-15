package internal

import (
	"encoding"
	"reflect"
	"time"

	"github.com/okysetiawan/go-document/errors"
)

func GetHeaderFromAny(in any, tag string) ([]string, error) {
	if in == nil {
		return nil, errors.New("types is nil").WithCode(errors.CodeWriteTypesNil)
	}

	type ListStringer interface {
		ListString() []string
	}

	if ls, ok := in.(ListStringer); ok {
		return ls.ListString(), nil
	}

	var (
		typeOf  = reflect.TypeOf(in)
		valueOf = reflect.ValueOf(in)
	)

	switch typeOf.Kind() {
	case reflect.Struct:
		return getHeaderByStruct(in, tag), nil
	case reflect.Ptr:
		return getHeaderByStruct(valueOf.Elem().Interface(), tag), nil
	case reflect.Slice, reflect.Array:
		if valueOf.Len() == 0 {
			return []string{}, nil
		}
		return getHeaderByStruct(valueOf.Index(0).Interface(), tag), nil
	default:
		return nil, errors.New("types should (struct,ptr,slice)").WithCode(errors.CodeWriteTypesInvalid)
	}
}

func getHeaderByStruct(in any, tag string) []string {
	var (
		headers         []string
		header          string
		typeOf, valueOf = reflect.TypeOf(in), reflect.ValueOf(in)
	)

	for i := 0; i < typeOf.NumField(); i++ {
		field := valueOf.Type().Field(i)
		header = field.Name

		if headerByTag := field.Tag.Get(tag); headerByTag != "" {
			header = headerByTag
		}

		if len(header) != 0 {
			headers = append(headers, header)
		}
	}

	return headers
}

func GetRowsFromAny(in interface{}) ([][]string, error) {
	if in == nil {
		return nil, errors.New("types is nil").WithCode(errors.CodeWriteTypesNil)
	}

	var (
		typeOf  = reflect.TypeOf(in)
		valueOf = reflect.ValueOf(in)
	)
	switch typeOf.Kind() {
	case reflect.Slice, reflect.Array:
		if valueOf.Len() == 0 {
			return [][]string{}, nil
		}
		return getRowsFromSlice(in), nil
	default:
		return nil, errors.New("types should (slice,array)").WithCode(errors.CodeWriteTypesInvalid)
	}
}

func getRowsFromSlice(slice interface{}) [][]string {
	var (
		rows    [][]string
		row     []string
		valueOf = reflect.ValueOf(slice)
	)

	for i := 0; i < valueOf.Len(); i++ {
		in := valueOf.Index(i).Interface()
		typeOfStruct := reflect.TypeOf(in)
		valueOfStruct := reflect.ValueOf(in)

		for j := 0; j < typeOfStruct.NumField(); j++ {
			types := valueOfStruct.Field(j).Interface()

			if t, ok := types.(time.Time); ok {
				format := valueOfStruct.Type().Field(j).Tag.Get("format")
				if format == "" {
					row = append(row, t.String())
					continue
				}
				row = append(row, t.Format(format))
				continue
			}

			if t, ok := types.(encoding.TextMarshaler); ok {
				bytes, _ := t.(encoding.TextMarshaler).MarshalText()
				row = append(row, string(bytes))
				continue
			}

			row = append(row, AnyToString(types))
		}

		rows = append(rows, row)
		row = []string{}
	}

	return rows
}
