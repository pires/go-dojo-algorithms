package binary_search_test

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
		ok, pos := binary_search.BinarySearchInt(tt.items, tt.item)
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
		ok, pos := binary_search.BinarySearchString(tt.items, tt.item)
		if ok != tt.result.ok || pos != tt.result.pos {
			t.Fatalf("TestBinarySearchString: expected: %+v, found: %+v", tt.result, result{ok, pos})
		}
	}
}

func TestGreaterThan(t *testing.T) {
	var computations = []struct {
		v1, v2 interface{}
		result bool
	}{
		// is 1 greater than 2?
		{
			1,
			2,
			false,
		},
		// is 2 greater than 1?
		{
			2,
			1,
			true,
		},
		// is 1 greater than 1?
		{
			1,
			1,
			false,
		},
		// is a greater than b?
		{
			"a",
			"b",
			false,
		},
		// is b greater than a?
		{
			"b",
			"a",
			true,
		},
		// is a greater than a?
		{
			"a",
			"a",
			false,
		},
	}

	for _, tt := range computations {
		if result := binary_search.GreaterThan(tt.v1, tt.v2); result != tt.result {
			t.Fatalf("TestGreaterThan: expected: %t, found: %t", tt.result, result)
		}
	}
}
