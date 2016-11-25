Merger
=====

A go package to merge structs of the same type, filling zero values with non zero values from another.

Merger is useful if you want to provide a default settings struct for something, or if you are working with database results and want to 'update' just some fields from a struct


## Usage

```
Merge(&dst, src)
````

Being `&dst` a pointer to the destination struct and `src` the origin struct

Example

``` 
package main

import (
	"fmt"
	"github.com/worg/merger"
)

type Person struct {
	ID      int
	Name    string
	Address string
}

var (
	defaultPerson = Person{
		Name:    `Homer`,
		Address: `742 Evergreen Terrace`,
	}
)

func main() {
	p := Person{Name: `John`}

	if err := merger.Merge(&p, defaultPerson); err != nil {
		// handle err
	}

    fmt.Printf("p is: %+v", p)
    // prints : p is {Id: 0, Name: John, Address: 742 Evergreen Terrace}
}

```


## Notes

Mergo won't merge non zero nor unexported fields , that is it will respect any value already present in dst.

## License

Copyright (c) 2014 Hiram Jerónimo Pérez worg{at}linuxmail[dot]org

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.