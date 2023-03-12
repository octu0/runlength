package runlength

import (
	"io"

	"github.com/pkg/errors"
)

type Decoder struct {
	r io.Reader
}

func (d *Decoder) Decode(out io.Writer) error {
	buf := make([]byte, 2)
	for {
		_, err := d.r.Read(buf[0:2])
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return errors.WithStack(err)
		}

		length := buf[0]
		values := make([]byte, length)
		for i := byte(0); i < length; i += 1 {
			values[i] = buf[1]
		}
		if _, err := out.Write(values); err != nil {
			return errors.Wrapf(err, "failed to decoded value")
		}
	}
	return nil
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r}
}
