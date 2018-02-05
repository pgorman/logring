logring
============================================================================

Logring is a Go package that records recently logged messages.

```
func Count() int
    Count reports the total number of messages logged.

func Recent() []string
    Recent returns the most recently logged messages.

func Writer(ringSize int) io.Writer
    Writer returns an io.Writer suitable as an argument to log.SetOutput. It
    saves the the most recent log entries, while also writing to Stderr. The
    ringSize argument sets the maximum number of messages to keep.
```

Usage Example
----------------------------------------------------------------------------

```
package main

import (
        "fmt"
        "log"

        "paulgorman.org/logring"
)

func main() {
        log.SetOutput(logring.Writer(5))
        for i := 0; i < 10; i++ {
                log.Println("Logring test number", i)
        }
        for _, v := range logring.Recent() {
                fmt.Println("Reprinting logring message:", v)
        }
        fmt.Println("Total errors logged:", logring.Count())
}
```

License (2-clause BSD license)
----------------------------------------------------------------------------

Copyright 2018 Paul Gorman

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
