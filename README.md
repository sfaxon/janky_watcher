Just an experiment in go.

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
go build jankywatcher.go
