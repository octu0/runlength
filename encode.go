package runlength

import (
	"io"

	"github.com/pkg/errors"
)

type Encoder struct {
	out io.Writer
}

func (e *Encoder) Encode(data []byte) error {
	if len(data) < 1 {
		return nil
	}

	currentValue := data[0]
	currentLength := byte(1)

	for i := 1; i < len(data); i += 1 {
		if data[i] != currentValue || currentLength == 255 {
			if _, err := e.out.Write([]byte{currentLength, currentValue}); err != nil {
				return errors.Wrapf(err, "failed to write data: len=%d, val=%v", currentLength, currentValue)
			}
			currentValue = data[i]
			currentLength = 1
		} else {
			currentLength += 1
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
