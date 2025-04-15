# unsignedlint

Go linter which finds potentially unsafe unsigned integer subtractions in Go code.

The linter checks for binary expressions of the form `x - y` where `y`
is an unsigned integer type (uint, uint8, uint16, uint32, uint64, uintptr).
Such operations can lead to unexpected integer wraps (underflow) if `y > x` at runtime.

**Note:** This initial version flags *all* such subtractions. It does not currently
check for safety guards (like `if x >= y { ... x - y ... }`). Adding such checks
requires more complex static analysis.

## Installation

```bash
go install github.com/fingon/unsignedlint@latest
```

## Usage

Run the linter on your package:

```bash
go vet -vettool=$(which unsignedlint) ./...
```

Or directly:

```bash
unsignedlint ./...
```

## Example

Code like this will be flagged:

```go
package main

func main() {
	var a uint = 5
	var b uint = 10
	c := a - b // Potential underflow: unsignedlint will report this line
	println(c)
}
```

## Known issues

Running it stand-alone may sometimes not work (Go bug? Who knows):

```
> unsignedlint ./...
unsignedlint: internal error: package "iter" without types was imported from "..."
```

In this case, `go vet` seems to work, for some reason.
