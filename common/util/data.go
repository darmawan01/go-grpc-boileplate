package util

import "strconv"

func DefaultValueString(defaultValue, data string) string {
	if data == "" {
		return defaultValue
	}
	return data
}

func DefaultValueInt(defaultValue int, data string) int {
	if data == "" {
		return defaultValue
	}

	dataInt, err := strconv.Atoi(data)
	if err != nil {
		return defaultValue
	}

	return dataInt
}

func DefaultValueBool(defaultValue bool, data string) bool {
	if data == "" {
		return defaultValue
	}

	dataBool, err := strconv.ParseBool(data)
	if err != nil {
		return defaultValue
	}

	return dataBool
}
