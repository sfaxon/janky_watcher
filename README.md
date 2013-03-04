## Janky Watcher

This is a [go](http://golang.org/) app that aggregates the last build status from [janky](https://github.com/github/janky).  The output is a single http response:
GOOD, BUILDING or JANKY.  On it's own this is really boring, but with the
[janky stop light](https://github.com/sfaxon/janky_stop_light) Arduino project it can
light up an LED based on the build status.

## configue

The sitelist.txt file is read every time a request comes in.  It should be a
line separated list of janky output URL's.

command line options are: 
  --port :8080    # the : is currently required

## building

compiling with go is as easy as: 
go build jankyapp/server.go

## including

you should be able to import the parsing library and use in your own server if you would like: 

    import "github.com/sfaxon/janky_watcher"
    func main() {
      http.HandleFunc("/PATH_YOU_WANT", func(w http.ResponseWriter, r *http.Request) {
        siteList := jankypack.ReadConfigFile('CONFIG_FILE_PATH')
        for i := 0; i < len(siteList); i++ {
        	siteList[i].WasLastBuildGood()
        }
        fmt.Fprintf(w, "%s\n", jankypack.WostCaseBuild(siteList))
      })
    }

