// Code generated by "esc -include=.*\.tmpl -o=tmpl.go -pkg=templates ./"; DO NOT EDIT.

package templates

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/testify/call.tmpl": {
		name:    "call.tmpl",
		local:   "testify/call.tmpl",
		size:    242,
		modtime: 1615224234,
		compressed: `
H4sIAAAAAAAC/0SOwUrEQAyG7z5FKD0olDyA4AP0IqKi5zBNa6AdJRN3WULefZnS3Tn98M833x/3iWfJ
DF2ide0i3M9iP4DvnFhOrLWRGfKvAY7lw/Q/WYQZunOe6uuNBIxoLb7SxhGP7kp5YehlgJ5XeH4BfCOl
jY21HPZeIga4/2173ypWbzBD0qU0/a6oizuNn5c/xrF8kQpNkiIQG7vH05EP1wAAAP//TC+HD/IAAAA=
`,
	},

	"/testify/function.tmpl": {
		name:    "function.tmpl",
		local:   "testify/function.tmpl",
		size:    4564,
		modtime: 1616701246,
		compressed: `
H4sIAAAAAAAC/8xY3WvjRhB/lv6KQYRDOnRLnw33kEt9d4F8EYteoRSj2CNHRF6lu+MeQd3/vcyuvixL
jnNtoXlItF8zv/n67Wyqao1ZLhGCVGtUlJcygA/G+ESinQkphqrKMxC3sni5R9opqedKlcqYqiLcPhcp
IQSrtCgCEDyJhUZjUKmqQrk2xq+qD5BnUCoIZUkgFrsHQk06AnGnckmX8nlH2qoGAIgh25JYPPNSFgZ9
LVvUOt2gVRQwsG4pt0LsSmQ1olxbke3IGN/vbMY/dmkRGOMsFXMeWkNfgWlMVtvVmNeXutWbgFdOkQMx
/IBx+8phT3u2kysbxtrpZxnMPvIh3+clqCqRoKabdIvGhATvGVcuNyKJoPK9Ok5iP8q+5/E2O2JxyB9a
3OD3MOCFpZ0IIt/zrE910t/NekMiqN2cOIVJDOgExZCDECKXhCpLV1iZCB7KsoAKlIUB/QCFRDG0aGLA
CIzPwOtIWBO+5/TIRqww/xOtAbVhqVyDuNQLUrsVgficY7HWvE4vzwicwRq0W6x8z7PHVCo32N/rNdJs
bC/1p1Tnq+TlGZtF4dx7jfRYriGX1JxxGPcHjXJn63+ivvFrq9mW597m2iOcHj1Jk4h7o8EnZw+H+C5V
6RYJVefeVG32zOtZd3jCqrRTnL01ATE2+AvO1+u7dPWUbhDqSrMrNrCpMuZ9WyFn4iuma1Si3m+MqE0U
v6TFDjuyai1vMumIybaWObV/+71njky3yOblclMfPkxD79VE9DyXhfY3D5u8qP8OAzP8Hve/51nn86+R
M70g3KPeFdTm2bdU0v/B/9PE5LX3VEsTvHbezH7eydUgljb3+ChHbAbB4vz67mq+TOaLZHlxvpgHcePk
sQCeEEEXwpmLYdUeGqnk46V8UMwz+Cluj7WlORgZt6XOl1mTOP8uDpkXPSQtnxwwyqyjlDA6GfsU87ye
4jOb40PynKCXYyb3yGcGIzn+rs3x6cJ4c/b3pq1sM+rjATTxBcmirw+/5kITT3r1SHXNYHC1H0hhucb3
slI5f4Wl6jqfuhG61DfpE66jyBhYxkDELNo0Mi5YbnvVq/lGxh4RsAuKAgtjnBCijjhI3O9kSCS4wOO6
ARm2OqOi2u8wGuchb0WqYH2bcluunrgHuiglqbIoUIXELZC3xgwV8EbxOZe5fgwj3+k7ciG0VNLNDuv0
9KvfefUHGJttso3ddbl6OpGpQ7Y02tMv5r/ezS+SMBLX8+Tr7c/L5HZ5cX51VScYh6ZmJTFglkgk+RY1
77DsebB+pIdxVjdsM6bDeTjaz1w3GqkVN8ozOBMXpXQ3fTm4BQZh4/kmvPbetJBGeWBPaI0vbGSMk3TP
1rjb2bGnF01Q8j+BFJ3AyH1/1cCn6FZc6m8qp/Y23evyZh/h3cMLoRafdlmGqjKnXQd8FzuCGT5Ro8P5
W4k2oyNokXXPLOX6nwDOMtv1DJ+3PH0qfcLgpy+t996eFDnVlo25EkZ++vrcW9fq4vc8keiaO35otjEQ
C9vBMv/13qV6484eJNi4XgvvLBvz+sSZt+LtP4yLXKLbOoF5CuQRG96K50tJ7mn+NgDjbdBY49tepMaA
iaD/7w3P+Pb/Gu7I3wEAAP//hRqVltQRAAA=
`,
	},

	"/testify/header.tmpl": {
		name:    "header.tmpl",
		local:   "testify/header.tmpl",
		size:    146,
		modtime: 1616699566,
		compressed: `
H4sIAAAAAAAC/0TMsQ7CMAwE0N1fYXWCgfwEEwvqwA9Y5GgrlLRKvVn374hW0M2+072IjNdUod0Iy2gd
KRHN6gBN17kUVF/JiLQVqFkvpCz2fNsAjUgPrN7vLykylWVurqdDuW3JjtytgPzOevPxR5Jy/l+fAAAA
//+wiJl6kgAAAA==
`,
	},

	"/testify/inline.tmpl": {
		name:    "inline.tmpl",
		local:   "testify/inline.tmpl",
		size:    50,
		modtime: 1615224234,
		compressed: `
H4sIAAAAAAAC/6quTklNy8xLVVDKzMvJzEtVqq1VqK4uSc0tyEksSVVQSk7MyVFS0AOLpual1NZyAQIA
AP//jg0rBDIAAAA=
`,
	},

	"/testify/inputs.tmpl": {
		name:    "inputs.tmpl",
		local:   "testify/inputs.tmpl",
		size:    153,
		modtime: 1615224234,
		compressed: `
H4sIAAAAAAAC/0yNMQrDQAwE+7xCGJdBDwjkAekCeYGCdeYKK0EnV4v+Hu7iwpXEsDsLLFqqKU3Vvnu0
KROYC93uxP2thewTxK/9HdqiZUawyaZXAtSWIzMXfnq1eAxJhy626uDismmo/7via2Ng0D5x8pzP5RcA
AP//CGcNOJkAAAA=
`,
	},

	"/testify/message.tmpl": {
		name:    "message.tmpl",
		local:   "testify/message.tmpl",
		size:    202,
		modtime: 1615224234,
		compressed: `
H4sIAAAAAAAC/zyN4WqDQBCE//sUiyi0oPsAhT5A/xRpS/9f4mgW9GLuTkNY9t2DB/HXDDPDN6o9BvGg
ckaMbkRJrVmhKgP5ayL+XU8JMUWz+sakCt+bqd4lXYh/cIZsCHvCf48F/O+mFWZ8DPnbzTB7y0Tugvj0
5Zd1B6oG50dQJQ1VmOjjk7hzwc1ICLmXgSoxa16/9XZws7wXqi1lWzwDAAD///ksisLKAAAA
`,
	},

	"/testify/results.tmpl": {
		name:    "results.tmpl",
		local:   "testify/results.tmpl",
		size:    169,
		modtime: 1615224234,
		compressed: `
H4sIAAAAAAAC/1yNPQrDMAyF957iETyGHKDQsXTvDQqRiyDY8OxMQncvSk0LmfTzfdIzWyVrEUyUtm+9
Te4w46u8BUlnJNlwvWF5frG7mWYkdZ9hJmWNzaN2LNGMWXMc9J2l3cnKkIUcHIdQ+Xt6liPw7x718gkA
AP//UHrX+akAAAA=
`,
	},

	"/": {
		name:  "/",
		local: `./`,
		isDir: true,
	},

	"/testify": {
		name:  "testify",
		local: `testify`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"./": {},

	"testify": {
		_escData["/testify/call.tmpl"],
		_escData["/testify/function.tmpl"],
		_escData["/testify/header.tmpl"],
		_escData["/testify/inline.tmpl"],
		_escData["/testify/inputs.tmpl"],
		_escData["/testify/message.tmpl"],
		_escData["/testify/results.tmpl"],
	},
}
