# `runlength`

a simple run-length algorithm implementation.

### Example Encode

```go
package main

import (
	"bytes"

	"github.com/octu0/runlength"
)

func main() {
	out := bytes.NewBuffer(nil)
	input := []byte{1, 1, 1, 1, 1, 2, 2, 2, 3, 4, 5, 5, 5, 5}

	enc := runlength.NewEncoder(out)
	if err := enc.Encode(input); err != nil {
		panic(err)
	}
	println(out.Bytes()) // => [5 1 3 2 1 3 1 4 4 5]
}
```

### Example Decode

```go
package main

import (
	"bytes"

	"github.com/octu0/runlength"
)

func main() {
	out := bytes.NewBuffer(nil)
	input := []byte{5, 1, 3, 2, 1, 3, 1, 4, 4, 5}

	dec := runlength.NewDecoder(bytes.NewReader(input))
	if err := dec.Decode(output); err != nil {
		panic(err)
	}
	println(out.Bytes()) // => [1 1 1 1 1 2 2 2 3 4 5 5 5 5]
}
```

# License

MIT, see LICENSE file for details.
