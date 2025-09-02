package runlength

import (
	"bytes"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Run("compress", func(tt *testing.T) {
		src := []byte{5, 1, 3, 2, 1, 3, 1, 4, 4, 5}
		b, err := NewDecoder().Decode(bytes.NewReader(src))
		if err != nil {
			tt.Errorf("no error:%+v", err)
		}
		tt.Logf("\n%v\n%v", src, b)

		if bytes.Equal(b, []byte{1, 1, 1, 1, 1, 2, 2, 2, 3, 4, 5, 5, 5, 5}) != true {
			tt.Errorf("actual=%v", b)
		}
	})
	t.Run("redandant", func(tt *testing.T) {
		src := []byte{1, 1, 1, 2, 1, 3, 1, 4, 1, 5}
		b, err := NewDecoder().Decode(bytes.NewReader(src))
		if err != nil {
			tt.Errorf("no error:%+v", err)
		}
		tt.Logf("\n%v\n%v", src, b)

		if bytes.Equal(b, []byte{1, 2, 3, 4, 5}) != true {
			tt.Errorf("actual=%v", b)
		}
	})
	t.Run("zero/250", func(tt *testing.T) {
		src := bytes.Repeat([]byte{0}, 250)
		out := bytes.NewBuffer(nil)
		if err := NewEncoder(out).Encode(src); err != nil {
			tt.Errorf("no error:%+v", err)
		}
		b, err := NewDecoder().Decode(bytes.NewReader(out.Bytes()))
		if err != nil {
			tt.Errorf("no error:%+v", err)
		}
		if bytes.Equal(b, src) != true {
			tt.Errorf("actual=%v", b)
		}
	})
}
