# UE Reader

UEReader is a simple Unreal Engine binary reader.

## Installation

```bash
go add github.com/gounreal/uereader
```

## Usage

```go
package main

import "github.com/gounreal/uereader"

func main() {
    var r io.ReadSeeker = ...

    ar := uereader.NewReader("main-reader", r)

    someBytes := ar.Bytes(16)
    aBoolean := ar.Bool()

    aSlice := uereader.ReadSlice[*MyType](ar, MyReaderFunc)

    if err := ar.Err(); err != nil {
        handleError(err)
    }
}
```
