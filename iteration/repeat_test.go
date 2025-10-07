package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("z")
	expected := "zzzzz"

	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}
