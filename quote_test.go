package quote

import "testing"

func TestSay(t *testing.T) {
	want := "Hello"
	if got := Say(); got != want {
		t.Errorf("Failed Expect %s, but got %s", want, got)
	}
}
