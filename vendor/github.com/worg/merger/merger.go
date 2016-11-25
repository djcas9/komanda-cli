// Copyright (c) 2015 Hiram Jerónimo Pérez https://worg.xyz

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package merger is an utility to merge structs of the same type
package merger

import (
	"errors"
	"reflect"
	"strings"
)

var (
	// ErrDistinctType occurs when trying to merge structs of distinct type
	ErrDistinctType = errors.New(`dst and src must be of the same type`)
	// ErrNoPtr occurs when no struct pointer is sent as destination
	ErrNoPtr = errors.New(`dst must be a pointer to a struct`)
	// ErrNilArguments occurs on receiving nil as arguments
	ErrNilArguments = errors.New(`no nil values allowed`)
	// ErrUnknown occurs if the type can't be merged
	ErrUnknown = errors.New(`could not merge`)
)

// Merge sets zero values from dst to non zero values of src
// accepts two structs of the same type as arguments
// dst must be a pointer to a struct
func Merge(dst, src interface{}) error {
	if dst == nil || src == nil {
		return ErrNilArguments
	}

	if !isStructPtr(dst) {
		return ErrNoPtr
	}

	if !typesMatch(src, dst) {
		return ErrDistinctType
	}

	vSrc := getValue(src)
	vDst := getValue(dst)

	for i := 0; i < vSrc.NumField(); i++ {
		df := vDst.Field(i)
		sf := vSrc.Field(i)
		if err := merge(df, sf); err != nil {
			return err
		}
	}

	return nil
}

// merge merges two reflect values based upon their kinds
func merge(dst, src reflect.Value) (err error) {
	if dst.CanSet() && !isZero(src) {
		switch dst.Kind() {
		// base types
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
			reflect.Float32, reflect.Float64, reflect.String, reflect.Bool:
			if isZero(dst) {
				switch dst.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					dst.SetInt(src.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
					dst.SetUint(src.Uint())
				case reflect.Float32, reflect.Float64:
					dst.SetFloat(src.Float())
				case reflect.String:
					dst.SetString(src.String())
				case reflect.Bool:
					dst.SetBool(src.Bool())
				}
			}
		case reflect.Slice:
			dst.Set(mergeSlice(dst, src))
		case reflect.Struct:
			// handle structs with IsZero method [ie time.Time]
			if fnZero, ok := dst.Type().MethodByName(`IsZero`); ok {
				res := fnZero.Func.Call([]reflect.Value{dst})
				if len(res) > 0 {
					if v, isok := res[0].Interface().(bool); isok && v {
						dst.Set(src)
					}
				}
			}

			for i := 0; i < src.NumField(); i++ {
				df := dst.Field(i)
				sf := src.Field(i)
				if err := merge(df, sf); err != nil {
					return err
				}
			}
		case reflect.Map:
			dst.Set(mergeMap(dst, src))
		case reflect.Ptr:
			// defer pointers
			if !dst.IsNil() {
				dst = getValue(dst)
			} else {
				dst.Set(src)
				break
			}
			if src.CanAddr() && src.IsNil() {
				src = getValue(src)
				if err := merge(dst, src); err != nil {
					return err
				}
			}
		default:
			return ErrUnknown
		}
	}
	return
}

// mergeSlice merges two slices only if dst slice fields are zero and
// src fields are nonzero
func mergeSlice(dst, src reflect.Value) (res reflect.Value) {
	for i := 0; i < src.Len(); i++ {
		if i >= dst.Len() {
			dst = reflect.Append(dst, src.Index(i))
		}
		if err := merge(dst.Index(i), src.Index(i)); err != nil {
			res = dst
			return
		}
	}

	res = dst
	return
}

// mergeMap traverses a map and merges the nonzero values of
// src into dst
func mergeMap(dst, src reflect.Value) (res reflect.Value) {
	if dst.IsNil() {
		dst = reflect.MakeMap(dst.Type())
	}

	for _, k := range src.MapKeys() {
		vs := src.MapIndex(k)
		vd := dst.MapIndex(k)
		if !vd.IsValid() && isZero(vd) && !isZero(vs) {
			dst.SetMapIndex(k, vs)
		}
	}

	return dst
}

// typesMatch typechecks two interfaces
func typesMatch(a, b interface{}) bool {
	return strings.TrimPrefix(reflect.TypeOf(a).String(), "*") == strings.TrimPrefix(reflect.TypeOf(b).String(), "*")
}

// getValue returns a reflect.Value from an interface
// deferring pointers if needed
func getValue(t interface{}) (rslt reflect.Value) {
	rslt = reflect.ValueOf(t)

	for rslt.Kind() == reflect.Ptr && !rslt.IsNil() {
		rslt = rslt.Elem()
	}

	return
}

// isStructPtr determines if a value is a struct pointer
func isStructPtr(v interface{}) bool {
	t := reflect.TypeOf(v)
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

// isZero is mostly stolen from encoding/json package's isEmptyValue function
// determines if a value has the zero value of its type
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr, reflect.Func:
		return v.IsNil()
	case reflect.Struct:
		zero := reflect.Zero(v.Type()).Interface()
		return reflect.DeepEqual(v.Interface(), zero)
	default:
		if !v.IsValid() {
			return true
		}

		zero := reflect.Zero(v.Type())
		return v.Interface() == zero.Interface()
	}

}
