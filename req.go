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

// GitApi http input structure
type Req struct {
	EntryPoint string       `json:"Entrypoint"` // Api base url
	Endpoint   string       `json:"Endpoint"`   // Api endpoint
	Token      string       `json:"Token"`      // Api auth token
	Header     *http.Header `json:"Header"`     // Http request header
	UrlVal     *url.Values  `json:"UrlVal"`     // Api url values
	Data       string       `json:"Data"`       // Json marshaled Info
}

func (t *Req) New(entryPoint string) *Req {
	header := make(http.Header)
	t.Header = &header
	t.EntryPoint = entryPoint
	return t
}

// Setup empty API url values
func (t *Req) UrlValInit() *Req {
	urlVal := make(url.Values)
	t.UrlVal = &urlVal
	return t
}
