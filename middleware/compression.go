package nexmiddleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/nex-gen-tech/nex/context" // Assuming the context is now in this package
	"github.com/nex-gen-tech/nex/router"  // Assuming the router is in this package
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}

// Compression checks the request's Accept-Encoding header and, if appropriate, wraps the response writer in a gzip writer.
func Compression() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			encodings := c.Request.Header.Get("Accept-Encoding")

			if strings.Contains(encodings, "gzip") {
				gzipWriter := gzip.NewWriter(c.Response)
				defer gzipWriter.Close()

				c.Response.Header().Set("Content-Encoding", "gzip")
				c.Response = &gzipResponseWriter{
					Writer:         gzipWriter,
					ResponseWriter: c.Response,
				}
			}

			next(c)
		}
	}
}
