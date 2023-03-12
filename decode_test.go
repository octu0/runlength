package runlength

import (
	"bytes"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Run("compress", func(tt *testing.T) {
		out := bytes.NewBuffer(nil)

		src := []byte{5, 1, 3, 2, 1, 3, 1, 4, 4, 5}
		r := bytes.NewReader(src)
		dec := NewDecoder(r)
		if err := dec.Decode(out); err != nil {
			tt.Errorf("no error:%+v", err)
		}

		b := out.Bytes()
		tt.Logf("\n%v\n%v", src, b)

		if bytes.Equal(b, []byte{1, 1, 1, 1, 1, 2, 2, 2, 3, 4, 5, 5, 5, 5}) != true {
			tt.Errorf("actual=%v", b)
		}
	})
	t.Run("redandant", func(tt *testing.T) {
		out := bytes.NewBuffer(nil)

		src := []byte{1, 1, 1, 2, 1, 3, 1, 4, 1, 5}
		r := bytes.NewReader(src)
		dec := NewDecoder(r)
		if err := dec.Decode(out); err != nil {
			tt.Errorf("no error:%+v", err)
		}

		b := out.Bytes()
		tt.Logf("\n%v\n%v", src, b)

		if bytes.Equal(b, []byte{1, 2, 3, 4, 5}) != true {
			tt.Errorf("actual=%v", b)
		}
	})
}
