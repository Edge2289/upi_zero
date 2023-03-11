package tool

import "testing"

func TestKrand(t *testing.T) {
	s := Krand(12, KC_RAND_KIND_NUM)
	t.Log(s)
}
