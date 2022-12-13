package lamport_clocks

import (
	"reflect"
	"testing"
)

func TestTimestamps(t *testing.T) {
	for _, n := range timeTests {
		got := GetTimestamps(n.input)
		if len(got) != len(n.expected) {
			t.Fatalf("Length of output %d != length of expected %d", len(got), len(n.expected))
		}
		if !reflect.DeepEqual(got, n.expected) {
			t.Fatalf("Output %v != want %v", got, n.expected)
		}
	}
	t.Log(len(timeTests), "test cases")
}

func BenchmarkTimestamps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range timeTests {
			GetTimestamps(n.input)
		}
	}
}