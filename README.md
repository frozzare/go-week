# go-week

Get the current week number.

View the [docs](http://godoc.org/github.com/frozzare/go-week).

## Installation

```
$ go get github.com/frozzare/go-week
```

## CLI

[vecka-go](https://github.com/kollegorna/vecka-go) by [kollegorna](https://github.com/kollegorna) is a great cli tool for getting the week! Or if you like the node version of [vecka](https://github.com/frozzare/vecka).

## Example

```go
package main

import (
    "fmt"
    "github.com/frozzare/go-week"
)

func main() {

    w := week.Now()
    
    fmt.Println(w)
}
```

## License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
