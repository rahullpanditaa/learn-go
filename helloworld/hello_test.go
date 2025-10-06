package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello", func(t *testing.T) {
		got := Hello()
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
