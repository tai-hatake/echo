package engine

import (
	"io"
	"time"

	"github.com/labstack/gommon/log"
)

type (
	// Engine defines an interface for HTTP server.
	Engine interface {
		SetHandler(Handler)
		SetLogger(*log.Logger)
		Start()
	}

	// Request defines an interface for HTTP request.
	Request interface {
		TLS() bool
		Scheme() string
		Host() string
		URI() string
		URL() URL
		Header() Header
		// Proto() string
		// ProtoMajor() int
		// ProtoMinor() int
		RemoteAddress() string
		Method() string
		Body() io.ReadCloser
		FormValue(string) string
		Object() interface{}
	}

	// Response defines an interface for HTTP response.
	Response interface {
		Header() Header
		WriteHeader(int)
		Write(b []byte) (int, error)
		Status() int
		Size() int64
		Committed() bool
		SetWriter(io.Writer)
		Writer() io.Writer
		Object() interface{}
	}

	// Header defines an interface for HTTP header.
	Header interface {
		Add(string, string)
		Del(string)
		Get(string) string
		Set(string, string)
		Object() interface{}
	}

	// URL defines an interface for HTTP request url.
	URL interface {
		SetPath(string)
		Path() string
		QueryValue(string) string
		Object() interface{}
	}

	// Config defines engine configuration.
	Config struct {
		Address      string
		TLSCertfile  string
		TLSKeyfile   string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	Handler interface {
		ServeHTTP(Request, Response)
	}

	HandlerFunc func(Request, Response)
)

func (h HandlerFunc) ServeHTTP(req Request, res Response) {
	h(req, res)
}
