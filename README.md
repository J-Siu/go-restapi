# go-restapi  [![Paypal donate](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/donate/?business=HZF49NM9D35SJ&no_recurring=0&currency_code=CAD)

A simple Golang REST api library.

### Table Of Content
<!-- TOC -->

- [Features](#features)
  - [api.go](#apigo)
- [Pro](#pro)
- [Doc](#doc)
- [Used By Project](#used-by-project)
- [Repository](#repository)
- [Contributors](#contributors)
- [Change Log](#change-log)
- [License](#license)

<!-- /TOC -->
<!--more-->
### Features

- API action
  - [x] Do
  - [x] Get
  - [x] Del
  - [x] Patch
  - [x] Post
  - [x] Put

#### api.go

```go
func (t *Api) New(property *Property) *Api
func New(property *Property) *Api func (t *Api) Body() *string
func (t *Api) Err() *string
func (t *Api) Ok() bool
func (t *Api) Output() *string
func (t *Api) Do() *Api
func (t *Api) Get() *Api
func (t *Api) Del() *Api
func (t *Api) Patch() *Api
func (t *Api) Post() *Api
func (t *Api) Put() *Api
func (t *Api) SetGet() *Api
func (t *Api) SetDel() *Api
func (t *Api) SetPatch() *Api
func (t *Api) SetPost() *Api
func (t *Api) SetPut() *Api
func (t *Api) ProcessOutput() *Api
func (t *Api) ProcessOutputError() *Api
func (t *Api) ProcessError() *Api
```

### Pro

- Easy to extend
- Small size

### Doc

- https://pkg.go.dev/github.com/J-Siu/go-restapi

### Used By Project

- [go-gitapi](https://github.com/J-Siu/go-gitapi)

### Repository

- [go-restapi](https://github.com/J-Siu/go-restapi)

### Contributors

- [John, Sing Dao, Siu](https://github.com/J-Siu)

### Change Log

- v1.0.0
  - Feature complete
- v1.0.1
  - Update [go-helper/v2](https://github.com/J-Siu/go-helper)
  - Update json tag

### License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
