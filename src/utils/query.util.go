package utils

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
)

func DecodeQueryParams(values url.Values, s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		return errors.New("destination must be a pointer")
	}
	v = v.Elem()

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}

		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("query")

		// Skip if there's no query tag
		if tag == "" {
			tag = fieldType.Name
		}

		paramValues, ok := values[tag]
		if !ok || len(paramValues) == 0 {
			continue
		}

		paramValue := paramValues[0]
		switch field.Kind() {
		case reflect.String:
			field.SetString(paramValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if intValue, err := strconv.ParseInt(paramValue, 10, 64); err == nil {
				field.SetInt(intValue)
			}
		case reflect.Float32, reflect.Float64:
			if floatValue, err := strconv.ParseFloat(paramValue, 64); err == nil {
				field.SetFloat(floatValue)
			}
			// Add more cases as necessary for other types
		}
	}

	return nil
}
