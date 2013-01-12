package main

import "fmt"
import "net/http"
// import "strings"
// import "io/ioutil"
import "code.google.com/p/go-html-transform/h5"


func parseWasLastBuildGood(p *h5.Parser) bool {
  returning := false
  count := 0
  setReturningTrueIfFistLIisGood:= func(n *h5.Node) {
    if "li" == n.Data() {
      if "good" == n.Attr[0].Value && 0 == count {
        returning = true
      }
      count++
    }
  }
	p.Top.Walk(setReturningTrueIfFistLIisGood)
  
  return returning
}

func wasLastBuildGoodOn(url string) bool {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Printf("error: %s", err)
  }
  defer resp.Body.Close()

  p := h5.NewParser(resp.Body)

	p.Parse()
  
  return parseWasLastBuildGood(p)
}

func main() {
  
  fmt.Printf("was last build good at: %s\n", wasLastBuildGoodOn("http://build.marshill.info/marshill"))

}
