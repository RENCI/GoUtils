package Convert

import (
	"errors"
	"strconv"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(i string) (int, error) {
	return strconv.Atoi(i)
}

func ByteToString(i byte) string { return string(i) }

func StringToByte(i string) (byte, error) {
	if len(i) > 0 {
		return byte(i[0]), nil
	} else {
		return 0, errors.New("empty string")
	}
}

func BoolToString(i bool) string {
	return strconv.FormatBool(i)
}

func StringToBool(i string) (bool, error) { return strconv.ParseBool(i) }

func Float32ToString(i float32) string {
	return strconv.FormatFloat(float64(i), 'E', -1, 64)
}

func StringToFloat32(i string) (float32, error) {
	value, err := strconv.ParseFloat(i, 32)
	if err != nil {
		return 0, err
	}
	return float32(value), nil
}

func Float64ToString(i float64) string {
	return strconv.FormatFloat(i, 'E', -1, 64)
}

func StringToFloat64(i string) (float64, error) {
	value, err := strconv.ParseFloat(i, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}
