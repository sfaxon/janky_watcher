package main

import "fmt"
import "net/http"
// import "strings"
// import "io/ioutil"
import "code.google.com/p/go-html-transform/h5"

func main() {
  resp, err := http.Get("http://build.marshill.info/marshill")
  if err != nil {
    fmt.Printf("error: %s", err)
  }
  defer resp.Body.Close()

  p := h5.NewParser(resp.Body)

	p.Parse()
  getFirstLI := func(n *h5.Node) {
    if "li" == n.Data() {
      fmt.Printf("%s : %s\n", n.Data(), n.Attr)
      fmt.Printf(" n.attr: %s\n", n.Attr[0].Value)
      if "good" == n.Attr[0].Value {
        fmt.Printf("   FOUND GOOD")
      }
    }
  }
	getFirstUL := func(n *h5.Node) {
    if "ul" == n.Data() {
      fmt.Printf("%s : %T\n", n.Data(), n.Attr)
      n.Walk(getFirstLI)
    }
	}
	p.Top.Walk(getFirstUL)
  
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
