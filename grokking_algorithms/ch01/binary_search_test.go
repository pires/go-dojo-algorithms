package algos_test

import (
	"testing"

	"github.com/pires/go-dojo-algorithms/grokking_algorithms/ch01"
)

type result struct {
	ok  bool
	pos int
}

func TestBinarySearchInt(t *testing.T) {
	var computations = []struct {
		items  []int
		item   int
		result result
	}{
		{
			[]int{1, 2, 3, 4, 5},
			4,
			result{true, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			6,
			result{false, 0},
		},
		{
			[]int{1},
			0,
			result{false, 0},
		},
		{
			[]int{},
			0,
			result{false, 0},
		},
	}

	for _, tt := range computations {
		ok, pos := algos.BinarySearchInt(tt.items, tt.item)
		if ok != tt.result.ok || pos != tt.result.pos {
			t.Fatalf("TestBinarySearchInt: expected: %+v, found: %+v", tt.result, result{ok, pos})
		}
	}
}

func TestBinarySearchString(t *testing.T) {
	var computations = []struct {
		items  []string
		item   string
		result result
	}{
		{
			[]string{"a", "b", "c", "d", "e"},
			"d",
			result{true, 3},
		},
		{
			[]string{"a", "b", "c", "d", "e"},
			"f",
			result{false, 0},
		},
		{
			[]string{"a"},
			"f",
			result{false, 0},
		},
		{
			[]string{},
			"a",
			result{false, 0},
		},
	}

	for _, tt := range computations {
		ok, pos := algos.BinarySearchString(tt.items, tt.item)
		if ok != tt.result.ok || pos != tt.result.pos {
			t.Fatalf("TestBinarySearchString: expected: %+v, found: %+v", tt.result, result{ok, pos})
		}
	}
}
