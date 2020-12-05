package stack

import "testing"

func TestDecodeString(t *testing.T) {
	result := decodeString("3[a]2[bc]")
	if result != "aaabcbc" {
		t.Fail()
	}
}
