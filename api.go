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

// A simple Golang REST api library.
package restapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/J-Siu/go-helper/v2/ezlog"
)

// Api
type Api struct {
	*Property
	log *ezlog.EzLog
	Req *Req `json:"Req,omitempty"`
	Res *Res `json:"Res,omitempty"`
}

// Setup a *Api
func (t *Api) New(property *Property) *Api {
	t.Property = property
	t.Req = new(Req).New(t.EntryPoint)
	t.Res = new(Res).New()

	t.log = ezlog.New()
	if t.Debug {
		t.log.SetLogLevel(ezlog.DEBUG)
	}
	return t
}

// Setup a *Api
func New(property *Property) *Api {
	return new(Api).New(property)
}

// Return Res.Body
func (t *Api) Body() *string {
	s := string(*t.Res.Body)
	return &s
}

// Return Res.Err
func (t *Api) Err() *string {
	return &t.Res.Err
}

// Return Res.Ok()
func (t *Api) Ok() bool {
	return t.Res.Ok()
}

// Return Res.Output
func (t *Api) Output() *string {
	return t.Res.Output
}

// Execute http request using info in Api.Req. Then put response info in Api.Res.
//
//	Api.Info, if not nil, will be
//			- auto marshal for send other than "GET"
//			- auto unmarshal from http response body
func (t *Api) Do() *Api {
	// Prepare Api Data
	if t.Method != http.MethodGet && t.Info != nil {
		j, _ := json.Marshal(&t.Info)
		t.Req.Data = string(j)
	}
	// Prepare url
	t.Res.Url, _ = url.Parse(t.Req.EntryPoint)
	t.Res.Url.Path = path.Join(t.Res.Url.Path, t.Req.Endpoint)
	if t.Req.UrlVal != nil {
		t.Res.Url.RawQuery = t.Req.UrlVal.Encode()
	}
	// Prepare request
	dataBufferP := bytes.NewBufferString(t.Req.Data)
	req, err := http.NewRequest(
		t.Method,
		t.Res.Url.String(),
		dataBufferP)
	if err != nil {
		t.Res.Err = err.Error()
	}
	// Set request headers
	req.Header = *t.Req.Header
	// Request
	// - Configure transport
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: t.SkipVerify},
	}

	client := &http.Client{Transport: transport}
	res, err := client.Do(req)
	if err != nil {
		t.Res.Err = err.Error()
	} else {
		// Response
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Res.Err = err.Error()
		}
		res.Body.Close()
		// Fill in self.Out
		t.Res.Body = &body
		t.Res.Header = &res.Header
		t.Res.Status = res.Status
	}

	// Unmarshal
	if t.Res.Err == "" {
		t.ProcessOutput()
	} else {
		t.ProcessError()
	}

	t.log.Debug().Nn("api").Mn(&t).Nn("api.Out.Body (decoded)").M(t.Res.Body).Out()
	return t
}

// Api Get action wrapper
func (t *Api) Get() *Api {
	t.SetGet()
	return t.Do()
}

// Api Del action wrapper
func (t *Api) Del() *Api {
	t.SetDel()
	return t.Do()
}

// Api Patch action wrapper
func (t *Api) Patch() *Api {
	t.SetPatch()
	return t.Do()
}

// Api Post action wrapper
func (t *Api) Post() *Api {
	t.SetPost()
	return t.Do()
}

// Api Put action wrapper
func (t *Api) Put() *Api {
	t.Method = http.MethodPut
	return t.Do()
}

func (t *Api) SetGet() *Api {
	t.Method = http.MethodGet
	return t
}

// Api set http Del
func (t *Api) SetDel() *Api {
	t.Method = http.MethodDelete
	return t
}

// Api set http Patch
func (t *Api) SetPatch() *Api {
	t.Method = http.MethodPatch
	return t
}

// Api set http Post
func (t *Api) SetPost() *Api {
	t.Method = http.MethodPost
	return t
}

// Api set http Put
func (t *Api) SetPut() *Api {
	t.Method = http.MethodPut
	return t
}

// Print HTTP Body into string pointer
func (t *Api) ProcessOutput() *Api {
	// Check API error
	t.ProcessOutputError()
	if t.Res.Ok() {
		// Unmarshal
		if t.Info == nil {
			// Return the whole JSON
			tmpStr := string(*t.Res.Body)
			t.Res.Output = &tmpStr
		} else {
			err := json.Unmarshal(*t.Res.Body, t.Info)
			if t.Res.Ok() && err == nil && t.Info != nil {
				// Use Info string func
				t.Res.Output = t.Info.StringP()
			}
		}
	}
	return t
}

// Print HTTP Body into string pointer
func (t *Api) ProcessOutputError() *Api {
	var e Err
	err := json.Unmarshal(*t.Res.Body, &e)
	if err == nil {
		// Use Info string func
		if e.Message != "" {
			t.Res.Err = e.String()
		}
	}
	return t
}

// Print HTTP Body Err into string pointer
func (t *Api) ProcessError() *Api {
	t.log.Debug().Nn("api.Out.Body").M(t.Res.Body).Out()
	// Unmarshal
	var output string
	if t.Res.Body != nil {
		output += string(*t.Res.Body)
	}
	output += t.Res.Err
	t.Res.Output = &output
	return t
}
