package common

import (
	"fmt"
	"strconv"
)

// StringTo converts a string to specify type.
type StringTo string

// Exist returns ture if the string is not empty
func (f StringTo) Exist() bool {
	return string(f) != string(0x1E)
}

// Uint8 returns an 8-bit unsigned integer for a given string or error
func (f StringTo) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

// Int returns an integer for a given string or error
func (f StringTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 0)
	return int(v), err
}

// Int64 returns a 64-bit integer for a given string or error
func (f StringTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	return int64(v), err
}

// MustUint8 returns an 8-bit unsigned integer without erroring
func (f StringTo) MustUint8() uint8 {
	v, _ := f.Uint8()
	return v
}

// MustInt returns an integer without erroring
func (f StringTo) MustInt() int {
	v, _ := f.Int()
	return v
}

// MustInt64 returns a 64-bit integer without erroring
func (f StringTo) MustInt64() int64 {
	v, _ := f.Int64()
	return v
}

// String returns the string if not empty
func (f StringTo) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}

// ToString converts any type to string.
func ToString(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

type argInt []int

func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}
