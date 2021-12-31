package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		t.Run("in Russian", func(t *testing.T) {
			got := Hello("Костя", "Russian")
			want := "Привет, Костя!"
			assertCorrectMessage(t, got, want)
		})

		t.Run("in English", func(t *testing.T) {
			got := Hello("F0rzend", "English")
			want := "Hello, F0rzend!"
			assertCorrectMessage(t, got, want)
		})

		t.Run("in other languge", func(t *testing.T) {
			got := Hello("F0rzend", "Unknown")
			want := "Hello, F0rzend!"
			assertCorrectMessage(t, got, want)
		})
	})

	t.Run("use World as name when an empty string is supplied", func(t *testing.T) {
		t.Run("in Russian", func(t *testing.T) {
			got := Hello("", "Russian")
			want := "Привет, Мир!"
			assertCorrectMessage(t, got, want)
		})

		t.Run("in English", func(t *testing.T) {
			got := Hello("", "English")
			want := "Hello, World!"
			assertCorrectMessage(t, got, want)
		})

		t.Run("in other languge", func(t *testing.T) {
			got := Hello("", "Unknown")
			want := "Hello, World!"
			assertCorrectMessage(t, got, want)
		})
	})
}
