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
c:> sleep 0.1
c:> sleep 0.1s
c:> sleep 100ms
c:> sleep 100000us
c:> sleep 100000000ns
~~~

Will sleep for 100ms (see [ParseDuration](http://golang.org/pkg/time/#ParseDuration) for time formats).

