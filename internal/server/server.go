package server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Run launches the server
func Run(addr string, tlsEnabled bool, certFile string, keyFile string) error {
	handler := http.HandlerFunc(handle)
	log.Println(fmt.Sprintf("Listening on %s", addr))

	if tlsEnabled {
		// http2 server with tls
		srv := &http.Server{
			Addr:    addr,
			Handler: handler,
		}

		http2.ConfigureServer(srv, nil)
		return srv.ListenAndServeTLS(certFile, keyFile)
	}

	// http2 server without tls
	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(handler, h2s),
	}
	return h1s.ListenAndServe()
}

func handle(w http.ResponseWriter, r *http.Request) {
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
	log.Println("Connection received")
	io.Copy(flushWriter{w}, r.Body)
}

// flushWriter is io.Writer implementation that tries to flush after each writing
type flushWriter struct {
	w io.Writer
}

// Write implements io.Writer.Write
func (fw flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if f, ok := fw.w.(http.Flusher); ok {
		f.Flush()
	}
	log.Printf("%s", p)
	return
}
