# gosleep [![Build status](https://ci.appveyor.com/api/projects/status/ymwk8atrb6tiig2y?svg=true)](https://ci.appveyor.com/project/tischda/gosleep)

Utility written in [Go](https://www.golang.org) to sleep for X seconds.

### Install

There are no dependencies.

~~~
go get github.com/tischda/gosleep
~~~

### Usage

~~~
Usage: gosleep [options] duration (default unit = seconds)

OPTIONS:
  -quiet
    	do not print anything
  -version
    	print version and exit
~~~

Examples:

~~~
$ gosleep 3s                                                                                                                                  [21:53]
[░░░░░░░░░░] 100%
~~~

Will sleep for 3s (see [ParseDuration](http://golang.org/pkg/time/#ParseDuration) for time formats).


### Known issues

Accuracy for actual sleeping time is not fantastic, but the delta should be under 20ms.


### Inspiration

* https://github.com/tianon/gosleep
* https://github.com/cheggaaa/pb
* https://github.com/gosuri/uiprogress
