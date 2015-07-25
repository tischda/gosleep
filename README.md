# sleep [![Build status](https://ci.appveyor.com/api/projects/status/7en2vn95tptff7n7?svg=true)](https://ci.appveyor.com/project/tischda/sleep)

Utility written in [Go](https://www.golang.org) to sleep for X seconds.

### Install

There are no dependencies.

~~~
go get github.com/tischda/sleep
~~~

### Usage

~~~
Usage: sleep seconds
  -version=false: print version and exit
~~~

Example:

~~~
c:> sleep 0.5
~~~

Will sleep for 500ms (sleeping time is not accurate below 14ms).
