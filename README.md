caseconv
========

A very simple case converter in go


### Install
    go get github.com/fredr/caseconv
    
### Test
    go test -v

### Use
```go
package main

import (
	"fmt"
	"github.com/fredr/caseconv"
)

func main() {
	fmt.Println(caseconv.CamelToSnake("HelloWorld"))
}
```
