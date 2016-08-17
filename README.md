# gosleep [![Build status](https://ci.appveyor.com/api/projects/status/7en2vn95tptff7n7?svg=true)](https://ci.appveyor.com/project/tischda/gosleep)

Utility written in [Go](https://www.golang.org) to sleep for X seconds.

### Install

There are no dependencies.

~~~
go get github.com/tischda/gosleep
~~~

### Usage

~~~
Usage: gosleep seconds
  -version=false: print version and exit
~~~

Examples:

~~~
gosleep 0.1
gosleep 0.1s
gosleep 100ms
gosleep 100000us
gosleep 100000000ns
~~~

Will sleep for 100ms (see [ParseDuration](http://golang.org/pkg/time/#ParseDuration) for time formats).

