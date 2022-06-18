//go:generate go run -tags generate gen.go
package main

import (
	"log"

	"github.com/zserge/lorca"
)

func main() {
	ui, _ := lorca.New("", "", 480, 320)

	addr, err := serve()
	if err != nil {
		log.Fatal(err)
	}
	ui.Load(addr)

	registerFunc()
	// Wait for the browser window to be closed
	<-ui.Done()
}
