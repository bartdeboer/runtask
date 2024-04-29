package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func convertToType(value string, targetType reflect.Type) (reflect.Value, error) {
	valueOf := reflect.ValueOf(value)

	if valueOf.CanConvert(targetType) {
		return valueOf.Convert(targetType), nil
	}

	var convertedValue reflect.Value
	var err error

	switch targetType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var intValue int64
		intValue, err = strconv.ParseInt(value, 10, targetType.Bits())
		convertedValue = reflect.ValueOf(intValue).Convert(targetType)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var uintValue uint64
		uintValue, err = strconv.ParseUint(value, 10, targetType.Bits())
		convertedValue = reflect.ValueOf(uintValue).Convert(targetType)
	case reflect.Float32, reflect.Float64:
		var floatValue float64
		floatValue, err = strconv.ParseFloat(value, targetType.Bits())
		convertedValue = reflect.ValueOf(floatValue).Convert(targetType)
	case reflect.Bool:
		var boolValue bool
		boolValue, err = strconv.ParseBool(value)
		convertedValue = reflect.ValueOf(boolValue)
	default:
		err = fmt.Errorf("unsupported string conversion to %s", targetType)
	}

	if err != nil {
		return reflect.Value{}, fmt.Errorf("error converting string to %s: %v", targetType, err)
	}

	return convertedValue, nil
}

func callFunc(fn reflect.Value, args []string) ([]reflect.Value, error) {
	var inArgs []reflect.Value
	fnType := fn.Type()
	numArgs := fnType.NumIn()
	isVariadic := fnType.IsVariadic()

	for i, argValue := range args {
		var argType reflect.Type

		// Determine the type of the current argument
		if isVariadic && i >= numArgs-1 {
			argType = fnType.In(numArgs - 1).Elem() // Get the element type of the variadic slice
		} else {
			argType = fnType.In(i)
		}

		// Convert the argument to the required type
		convertedArg, err := convertToType(argValue, argType)
		if err != nil {
			return nil, fmt.Errorf("error converting argument %d for function %v: %v", i, fnType.Name(), err)
		}
		inArgs = append(inArgs, convertedArg)
	}

	return fn.Call(inArgs), nil
}
