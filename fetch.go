package main

import "fmt"
import "net/http"
// import "strings"
// import "io/ioutil"
import "code.google.com/p/go-html-transform/h5"

func main() {
  resp, err := http.Get("http://build.marshill.info")
  if err != nil {
    fmt.Printf("error: %s", err)
  }
  defer resp.Body.Close()

  p := h5.NewParser(resp.Body)

	p.Parse()
	i := 0
	f := func(n *h5.Node) {
    fmt.Printf("%s\n", n.Data())
		i++
	}
	p.Top.Walk(f)
  
  fmt.Printf("hello\n")
  
  // p := h5.NewParserFromString("<html><body><a>foo</a><div>bar</div></body></html>")
  // resp, err := http.Get("http://build.marshill.info")
  // if err != nil {
  //   fmt.Printf("error: %s", err)
  // }
  // defer resp.Body.Close()
  // body, err := ioutil.ReadAll(resp.Body)
  // 
  //   p := h5.NewParser(resp.Body)
  //   parse_err := p.Parse()
  //   if parse_err != nil {
  //     panic(parse_err)
  //   }
  //   
  //   // tree := p.Tree()
  //   
  //   p.Top.Walk(nodeSomething)
  // 
  //   // tree.Walk(func(n *Node) {
  //   //    // do something with the node
  //   // })
  //   
  //   // fmt.Printf("%s", body)
  // fmt.Printf("hello, world\n")
}
