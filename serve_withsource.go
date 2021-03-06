//go:build withsource
// +build withsource

//构建时打包整个资源
package main

import (
	"fmt"
	"net"
	"net/http"
)

func serve() (string, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	go func() {
		defer ln.Close()
		http.Serve(ln, http.FileServer(FS))
	}()
	return fmt.Sprintf("http://%s", ln.Addr()), nil
}
