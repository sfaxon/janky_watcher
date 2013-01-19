package main

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"flag"
	"log"
	"net/http"
	"net/url"
	"code.google.com/p/go-html-transform/h5"
)

const (
	UNKNOWN  int = 0
	GOOD     int = 1
	BUILDING int = 2
	JANKY    int = 3
)

type Build struct {
	url    string // where are we looking
	status int    // building, good, janky
}

func (b Build) String() string {
	switch {
	case b.status == GOOD:
		return "GOOD"
	case b.status == BUILDING:
		return "BUILDING"
	case b.status == JANKY:
		return "JANKY"
	}
	return "UNKNOWN STATE"
}

// b *Build is pasesd by pointer, so the Build struct is modified
// not sure if this is idomatic or just a mess. leaning toward messy
func (b *Build) wasLastBuildGood() int {
	resp, err := http.Get(b.url)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Printf("fetched: %s\n", b.url)
	defer resp.Body.Close()

	p := h5.NewParser(resp.Body)

	p.Parse()

	b.status = parseWasLastBuildGood(p)
	return b.status
}

func wostCaseBuild(builds []Build) Build {
	returning := Build{}
	for i := 0; i < len(builds); i++ {
		if builds[i].status > returning.status {
			returning = builds[i]
		}
	}
	return returning
}

func parseWasLastBuildGood(p *h5.Parser) int {
	returning := 0 //Build{url: url} // default to unknown state
	count := 0
	setReturningTrueIfFistLIisGood := func(n *h5.Node) {
		if "li" == n.Data() && 0 == count {
			switch n.Attr[0].Value {
			case "good":
				returning = GOOD
			case "building":
				returning = BUILDING
			case "janky":
				returning = JANKY
			}
			count++
		}
	}
	p.Top.Walk(setReturningTrueIfFistLIisGood)

	return returning
}


// Reads the sitelist.txt file and returns an array of Build structs
// The sitelist.txt file is expected to be a line seperated list of arrays
func ReadConfigFile(filename string) []Build {
	returning := []Build{}
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return returning
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		_, pErr := url.Parse(s)
		// only add items that are valid urls
		if nil == pErr {
			returning = append(returning, Build{url: s})
		}
		line, isPrefix, err = r.ReadLine()
	}
	if isPrefix {
		fmt.Println("error: buffer size to small")
		return returning
	}
	if err != io.EOF {
		fmt.Println(err)
		return returning
	}
	return returning
}

// fetch.go --port :1987
var addr = flag.String("port", ":1718", "http service address")

func main() {
	flag.Parse()
	filename := "sitelist.txt"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		siteList := ReadConfigFile(filename)
		for i := 0; i < len(siteList); i++ {
			siteList[i].wasLastBuildGood()
		}
		fmt.Fprintf(w, "%s\n", wostCaseBuild(siteList))
	})
	fmt.Printf("listening on port %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
