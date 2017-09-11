package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	fmt.Println("The proxy has started. Logging into the console:")

	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":3128", proxy))
}
