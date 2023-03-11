package uniqueid

import "testing"

func TestSn(t *testing.T) {
	s := GenSn(SN_RETAIL_ORDER)
	t.Log(s)
}