package conv

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errTerminatorNotPresent = errors.New("couldn't parse string, string doesn't conatin terminator '-'")
	errInvalidRow           = errors.New("string doesn't contain key or value")
	errInvalidInt           = errors.New("expected int")
	errMultipleOfThree      = errors.New("number multiple of three")
)


func ParseRow(row string) (key int, value int, err error) {
	if strings.ContainsRune(row, '-') {
		slicedString := strings.Split(row, "-")
		if slicedString[0] == "" || slicedString[1] == "" {
			err = errInvalidRow
			return
		}

		key, err = strconv.Atoi(slicedString[0])
		if err != nil {
			err = errInvalidInt
			return
		}

		value, err = strconv.Atoi(slicedString[1])
		if err != nil {
			err = errInvalidInt
			return
		}
	} else {
		err = errTerminatorNotPresent
		return
	}
	if value%3 == 0 {
		err = errMultipleOfThree
		return
	}
	return
}