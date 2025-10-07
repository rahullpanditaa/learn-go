package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat character given number of times", func(t *testing.T) {
		got := Repeat("w", 7)
		want := "wwwwwww"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("given number is 0", func(t *testing.T) {
		got := Repeat("z", 0)
		want := ""

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

}
