package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	re := add(1, 2)
	if re != 3 {
		t.Errorf("1 + 2 + 3 = %d", re)
	}
}

func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	re := add(1, 2)
	if re != 3 {
		t.Errorf("1 + 2 + 3 = %d", re)
	}
}
func TestAdd3(t *testing.T) {
	var dataset = []struct {
		x   int
		y   int
		out int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{333, 666, 999},
	}
	for _, d := range dataset {
		res := add(d.x, d.y)
		if res != d.out {
			t.Errorf("%d+%d = %d, want %d", d.x, d.y, res, d.out)
		}
	}

}
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re := add(1, 2)
		if re != 3 {
			b.Errorf("1 + 2 + 3 = %d", re)
		}

	}
}

func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < 10000; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 10000; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}
