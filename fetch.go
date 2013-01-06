package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
	resp, err := http.Get("http://build.marshill.info")
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	fmt.Printf("%s", body)
	fmt.Printf("hello, world\n")
}
