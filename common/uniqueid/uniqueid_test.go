package uniqueid

import "testing"

func TestGenId(t *testing.T) {
	s := GenId()
	t.Log(s)
}
