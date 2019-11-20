// Copyright (c) 2012-present The upper.io/db authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package sqlbuilder

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"upper.io/db.v2"
)

type scanner struct {
	v db.Unmarshaler
}

func (u scanner) Scan(v interface{}) error {
	return u.v.UnmarshalDB(v)
}

var _ sql.Scanner = scanner{}

//------

type jsonbType struct {
	V interface{}
}

func (j *jsonbType) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("Scan source was not []bytes")
	}

	v := jsonbType{}
	if err := json.Unmarshal(b, &v.V); err != nil {
		return err
	}
	*j = v
	return nil
}

func (j jsonbType) Value() (driver.Value, error) {
	if v, ok := j.V.(json.RawMessage); ok {
		return string(v), nil
	}
	b, err := json.Marshal(j.V)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//------

type stringArray []string

func (a *stringArray) Scan(src interface{}) error {
	if src == nil {
		*a = stringArray{}
		return nil
	}
	b, ok := src.([]byte)
	if !ok {
		return errors.New("Scan source was not []bytes")
	}
	if len(b) == 0 {
		return nil
	}
	s := string(b)[1 : len(b)-1]
	if s == "" {
		return nil
	}
	results := strings.Split(s, ",")
	*a = stringArray(results)
	return nil
}

// Value implements the driver.Valuer interface.
func (a stringArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, 2*N bytes of quotes,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+3*n)
		b[0] = '{'

		b = appendArrayQuotedString(b, a[0])
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = appendArrayQuotedString(b, a[i])
		}

		return append(b, '}'), nil
	}

	return []byte{'{', '}'}, nil
}

func appendArrayQuotedString(b []byte, v string) []byte {
	b = append(b, '"')
	for {
		i := strings.IndexAny(v, `"\`)
		if i < 0 {
			b = append(b, v...)
			break
		}
		if i > 0 {
			b = append(b, v[:i]...)
		}
		b = append(b, '\\', v[i])
		v = v[i+1:]
	}
	return append(b, '"')
}

//------

type int64Array []int64

func (a *int64Array) Scan(src interface{}) error {
	if src == nil {
		*a = int64Array{}
		return nil
	}
	b, ok := src.([]byte)
	if !ok {
		return errors.New("Scan source was not []bytes")
	}
	if len(b) == 0 {
		return nil
	}

	s := string(b)[1 : len(b)-1]
	results := []int64{}
	if s != "" {
		parts := strings.Split(s, ",")
		for _, n := range parts {
			i, err := strconv.ParseInt(n, 10, 64)
			if err != nil {
				return err
			}
			results = append(results, i)
		}
	}
	*a = int64Array(results)
	return nil
}

// Value implements the driver.Valuer interface.
func (a int64Array) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, N bytes of values,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+2*n)
		b[0] = '{'

		b = strconv.AppendInt(b, a[0], 10)
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = strconv.AppendInt(b, a[i], 10)
		}

		return append(b, '}'), nil
	}

	return []byte{'{', '}'}, nil
}
