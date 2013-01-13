package main

import "fmt"
import "net/http"
import "bufio"
import "io"
import "os"
import "errors"
// import "strings"
// import "io/ioutil"
import "code.google.com/p/go-html-transform/h5"


type Build struct {
  status string // building, good, janky
}

func parseWasLastBuildGood(p *h5.Parser) Build {
  returning := Build{status: "janky"} // default to failed state
  count := 0
  setReturningTrueIfFistLIisGood:= func(n *h5.Node) {
    if "li" == n.Data() && 0 == count {
      switch n.Attr[0].Value {
      case "good":
        returning.status = "good"
      case "building":
        returning.status = "building"
      }
      count++
    }
  }
	p.Top.Walk(setReturningTrueIfFistLIisGood)
  
  return returning
}

func wasLastBuildGoodOn(url string) Build {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Printf("error: %s", err)
  }
  defer resp.Body.Close()

  p := h5.NewParser(resp.Body)

	p.Parse()
  
  return parseWasLastBuildGood(p)
}

func ReadConfigFile(filename string) []string {
  returning := []string{}
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
    returning = append(returning, s)
    line, isPrefix, err = r.ReadLine()
  }
  if isPrefix {
    fmt.Println(errors.New("buffer size to small"))
    return returning
  }
  if err != io.EOF {
    fmt.Println(err)
    return returning
  }
  return returning
}

func main() {
  
  filename := "sitelist.txt"
  siteList := ReadConfigFile(filename)

  // fmt.Printf("%T %s\n", siteList, siteList)

  for i := 0; i < len(siteList); i++ {
    fmt.Printf("%s %s\n", siteList[i], wasLastBuildGoodOn(siteList[i]))
  }

}
