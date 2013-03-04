package main

import (
	"flag"
	"fmt"
	"jankywatcher/jankypack"
	"log"
	"net/http"
)

// fetch.go --port :1987
var addr = flag.String("port", ":1718", "http service address")

func main() {
	flag.Parse()
	filename := "sitelist.txt"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		siteList := jankypack.ReadConfigFile(filename)
		for i := 0; i < len(siteList); i++ {
			siteList[i].WasLastBuildGood()
		}
		fmt.Fprintf(w, "%s\n", jankypack.WostCaseBuild(siteList))
	})
	fmt.Printf("listening on port %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
