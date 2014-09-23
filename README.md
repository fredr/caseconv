caseconv
========

A very simple case converter in go


### Installation
    go get github.com/fredr/caseconv
    
### Test
    go test -v

### Use
    package main
    
    import (
      "fmt"
      "github.com/fredr/caseconv"
    )
    
    func main() {
      fmt.Println(caseconv.CamelToSnake("HelloWorld"))
    }
