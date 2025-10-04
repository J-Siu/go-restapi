/*
The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

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
*/

package restapi

import (
	"net/http"
	"net/url"
)

// GitApi http output structure
type Res struct {
	Url    *url.URL     `json:"Url"`    // In.Uri + In.Endpoint
	Header *http.Header `json:"Header"` // Http response header
	Status string       `json:"Status"` // Http response status
	Body   *[]byte      `json:"Body"`
	Output *string      `json:"Output"` // Api response body in string
	Err    string       `json:"Err"`
}

func (t *Res) New() *Res {
	return t
}

// Check response status == 2xx
func (t *Res) Ok() bool {
	return t.Status != "" && t.Status[0] == '2' && t.Err == ""
}
