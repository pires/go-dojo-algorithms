package algos_test

import (
	"testing"

	"github.com/pires/go-dojo-algorithms/grokking_algorithms/ch02"
)

type resultInt struct {
	itemsAsc  []int
	itemsDesc []int
}

func TestSortInts(t *testing.T) {
	var computations = []struct {
		items  []int
		result resultInt
	}{
		{
			[]int{1, 9, 2, 8, 3, 0, 7, 4, 6, 5},
			resultInt{
				itemsAsc:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				itemsDesc: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			},
		},
		{
			[]int{1, 1},
			resultInt{
				itemsAsc:  []int{1, 1},
				itemsDesc: []int{1, 1},
			},
		},
		{
			[]int{1},
			resultInt{
				itemsAsc:  []int{1},
				itemsDesc: []int{1},
			},
		},
		{
			[]int{},
			resultInt{
				itemsAsc:  []int{},
				itemsDesc: []int{},
			},
		},
	}

	for _, tt := range computations {
		itemsAsc, itemsDesc := algos.SortInts(tt.items)
		for pos, v := range itemsAsc {
			if v != tt.result.itemsAsc[pos] {
				t.Fatalf("TestSortInts: expected: %+v, found: %+v", tt.result.itemsAsc, itemsAsc)
			}
		}
		for pos, v := range itemsDesc {
			if v != tt.result.itemsDesc[pos] {
				t.Fatalf("TestSortInts: expected: %+v, found: %+v", tt.result.itemsDesc, itemsDesc)
			}
		}
	}
}

type resultString struct {
	itemsAsc  []string
	itemsDesc []string
}

func TestSortStrings(t *testing.T) {
	var computations = []struct {
		items  []string
		result resultString
	}{
		{
			[]string{"d", "a", "c", "e", "b"},
			resultString{
				itemsAsc:  []string{"a", "b", "c", "d", "e"},
				itemsDesc: []string{"e", "d", "c", "b", "a"},
			},
		},
		{
			[]string{"a", "a"},
			resultString{
				itemsAsc:  []string{"a", "a"},
				itemsDesc: []string{"a", "a"},
			},
		},
		{
			[]string{"a"},
			resultString{
				itemsAsc:  []string{"a"},
				itemsDesc: []string{"a"},
			},
		},
		{
			[]string{},
			resultString{
				itemsAsc:  []string{},
				itemsDesc: []string{},
			},
		},
	}

	for _, tt := range computations {
		itemsAsc, itemsDesc := algos.SortStrings(tt.items)
		for pos, v := range itemsAsc {
			if v != tt.result.itemsAsc[pos] {
				t.Fatalf("TestSortStrings: expected: %+v, found: %+v", tt.result.itemsAsc, itemsAsc)
			}
		}
		for pos, v := range itemsDesc {
			if v != tt.result.itemsDesc[pos] {
				t.Fatalf("TestSortStrings: expected: %+v, found: %+v", tt.result.itemsDesc, itemsDesc)
			}
		}
	}
}
