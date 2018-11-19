// Copyright (c) 2014 Hiram Jerónimo Pérez worg{at}linuxmail[dot]org

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
	"log"
	"reflect"
	"strings"
)

var (
	// ErrDistinctType occurs when trying to merge structs of distinct type
	ErrDistinctType = errors.New(`dst and src must be of the same type`)
	// ErrNoPtr occurs when no struct pointer is sent as destination
	ErrNoPtr = errors.New(`src must be a pointer to a struct`)
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
		sf := vSrc.Field(i)
		df := vDst.Field(i)
		if err := merge(df, sf); err != nil {
			return err
		}
	}

	return nil
}

func merge(dst, src reflect.Value) (err error) {
	if dst.CanSet() && !isZero(src) {
		switch dst.Kind() {
		case reflect.Int, reflect.Int64, reflect.Float32, reflect.Float64, reflect.String, reflect.Bool:
			if isZero(dst) {
				switch dst.Kind() {
				case reflect.Int, reflect.Int64:
					dst.SetInt(src.Int())
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
		case reflect.Map:
			dst.Set(mergeMap(dst, src))
		case reflect.Struct:
			for i := 0; i < src.NumField(); i++ {
				sf := src.Field(i)
				df := dst.Field(i)
				if err := merge(df, sf); err != nil {
					return err
				}
			}
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
	log.Printf("\n\n\nRES: %+v\n\n\n", res.Interface())
	return
}

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

func typesMatch(a, b interface{}) bool {
	return strings.TrimPrefix(reflect.TypeOf(a).String(), "*") == strings.TrimPrefix(reflect.TypeOf(b).String(), "*")
}

func getValue(t interface{}) (rslt reflect.Value) {
	rslt = reflect.ValueOf(t)

	for rslt.Kind() == reflect.Ptr && !rslt.IsNil() {
		rslt = rslt.Elem()
	}

	return
}

func isStructPtr(v interface{}) bool {
	t := reflect.TypeOf(v)
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func isZero(v reflect.Value) bool {
	if !v.CanSet() {
		return false
	}

	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		t := true
		for i := 0; i < v.Len(); i++ {
			t = t && isZero(v.Index(i))
		}
		return t
	case reflect.Struct:
		t := true
		for i := 0; i < v.NumField(); i++ {
			t = t && isZero(v.Field(i))
		}
		return t
	}
	// Compare other types directly:
	t := reflect.Zero(v.Type())
	return v.Interface() == t.Interface()
}
