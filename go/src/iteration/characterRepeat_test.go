package iteration

import "testing"

func TestCharacterRepeat(t *testing.T) {
	output := Repeat("a", 5)
	expected := "aaaaa"

	if output != expected {
		t.Errorf("got %q expected %q", output, expected)
	}
}

/* Sample Benchmarking Test */
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
