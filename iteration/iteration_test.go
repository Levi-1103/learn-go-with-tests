package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("7 repeats set", func(t *testing.T) {
		repeated := Repeat("a", 7)
		expected := "aaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func ExampleRepeat() {
	repeat := Repeat("a", 7)
	fmt.Println(repeat)
	// Output: aaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}
