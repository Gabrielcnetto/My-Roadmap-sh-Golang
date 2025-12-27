package bench

import "testing"

func JoinStrings(strs []string) string {
	var result string
	for _, s := range strs {
		result += s
	}
	return result
}

func BenchmarkJoinStrings(b *testing.B) {
	strs := []string{"Hello", ", ", "world", "!"}

	// The benchmark runner will call this function b.N times
	for i := 0; i < b.N; i++ {
		JoinStrings(strs)
	}
}
