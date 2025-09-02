package runlength

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
)

type Decoder struct {
}

func (d *Decoder) Decode(r io.Reader) ([]byte, error) {
	out := bytes.NewBuffer(nil)
	buf := make([]byte, 2)
	for {
		if _, err := r.Read(buf[0:2]); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, errors.WithStack(err)
		}

		length := buf[0]
		values := make([]byte, length)
		for i := byte(0); i < length; i += 1 {
			values[i] = buf[1]
		}
		if _, err := out.Write(values); err != nil {
			return nil, errors.Wrapf(err, "failed to decoded value")
		}
	}
	return out.Bytes(), nil
}

func NewDecoder() *Decoder {
	return &Decoder{}
}
