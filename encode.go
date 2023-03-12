package runlength

import (
	"io"

	"github.com/pkg/errors"
)

type Encoder struct {
	out io.Writer
}

func (e *Encoder) Encode(data []byte) error {
	currentValue := data[0]
	currentLength := byte(1)

	for i := 1; i < len(data); i += 1 {
		out := true
		if currentValue == data[i] {
			out = false
			currentLength += 1
			if 255 <= currentLength {
				out = true
			}
		}

		if out {
			if _, err := e.out.Write([]byte{currentLength, currentValue}); err != nil {
				return errors.Wrapf(err, "failed to write data:%d %v", currentLength, currentValue)
			}
			currentLength = 1
			currentValue = data[i]
		}
	}
	if _, err := e.out.Write([]byte{currentLength, currentValue}); err != nil {
		return errors.Wrapf(err, "failed to write data:%d %v", currentLength, currentValue)
	}
	return nil
}

func NewEncoder(out io.Writer) *Encoder {
	return &Encoder{out}
}
