package runlength

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	t.Run("same data", func(tt *testing.T) {
		out := bytes.NewBuffer(nil)

		enc := NewEncoder(out)
		if err := enc.Encode([]byte{1, 1, 1, 1, 1, 2, 2, 2, 3, 4, 5, 5, 5, 5}); err != nil {
			tt.Errorf("no error:%+v", err)
		}
		b := out.Bytes()
		tt.Logf("%v", b)

		if (b[0] == 5 && b[1] == 1) != true {
			tt.Errorf("5 times 1")
		}
		if (b[2] == 3 && b[3] == 2) != true {
			tt.Errorf("3 times 1")
		}
		if (b[4] == 1 && b[5] == 3) != true {
			tt.Errorf("1 times 3")
		}
		if (b[6] == 1 && b[7] == 4) != true {
			tt.Errorf("1 times 4")
		}
		if (b[8] == 4 && b[9] == 5) != true {
			tt.Errorf("4 times 5")
		}
	})
	t.Run("non same", func(tt *testing.T) {
		out := bytes.NewBuffer(nil)

		enc := NewEncoder(out)
		if err := enc.Encode([]byte{1, 2, 3, 4, 5}); err != nil {
			tt.Errorf("no error:%+v", err)
		}
		b := out.Bytes()
		tt.Logf("%v", b)

		if (b[0] == 1 && b[1] == 1) != true {
			tt.Errorf("1 times 1")
		}
		if (b[2] == 1 && b[3] == 2) != true {
			tt.Errorf("1 times 2")
		}
		if (b[4] == 1 && b[5] == 3) != true {
			tt.Errorf("1 times 3")
		}
		if (b[6] == 1 && b[7] == 4) != true {
			tt.Errorf("1 times 4")
		}
		if (b[8] == 1 && b[9] == 5) != true {
			tt.Errorf("1 times 5")
		}
	})
}
