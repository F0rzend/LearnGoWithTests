package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertResult := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertResult(t, repeated, expected)
	})

	t.Run("repeat 100 times", func(t *testing.T) {
		repeated := Repeat("b", 1)
		expected := "b"
		assertResult(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeat := Repeat("qw", 4)
	fmt.Println(repeat)
	// Output: qwqwqwqw
}
