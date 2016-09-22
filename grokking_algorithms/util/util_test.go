package util_test

import (
	"testing"

	"github.com/pires/go-dojo-algorithms/grokking_algorithms/util"
)

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
		if result := util.GreaterThan(tt.v1, tt.v2); result != tt.result {
			t.Fatalf("TestGreaterThan: expected: %t, found: %t", tt.result, result)
		}
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
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
			true,
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
			true,
		},
	}

	for _, tt := range computations {
		if result := util.GreaterThanOrEqual(tt.v1, tt.v2); result != tt.result {
			t.Fatalf("TestGreaterThanOrEqual: expected: %t, found: %t", tt.result, result)
		}
	}
}

func TestLowerThan(t *testing.T) {
	var computations = []struct {
		v1, v2 interface{}
		result bool
	}{
		// is 1 lower than 2?
		{
			1,
			2,
			true,
		},
		// is 2 lower than 1?
		{
			2,
			1,
			false,
		},
		// is 1 lower than 1?
		{
			1,
			1,
			false,
		},
		// is a lower than b?
		{
			"a",
			"b",
			true,
		},
		// is b lower than a?
		{
			"b",
			"a",
			false,
		},
		// is a lower than a?
		{
			"a",
			"a",
			false,
		},
	}

	for _, tt := range computations {
		if result := util.LowerThan(tt.v1, tt.v2); result != tt.result {
			t.Fatalf("TestGreaterThan: expected: %t, found: %t", tt.result, result)
		}
	}
}

func TestLowerThanOrEqual(t *testing.T) {
	var computations = []struct {
		v1, v2 interface{}
		result bool
	}{
		// is 1 lower than 2?
		{
			1,
			2,
			true,
		},
		// is 2 lower than 1?
		{
			2,
			1,
			false,
		},
		// is 1 lower than 1?
		{
			1,
			1,
			true,
		},
		// is a lower than b?
		{
			"a",
			"b",
			true,
		},
		// is b lower than a?
		{
			"b",
			"a",
			false,
		},
		// is a lower than a?
		{
			"a",
			"a",
			true,
		},
	}

	for _, tt := range computations {
		if result := util.LowerThanOrEqual(tt.v1, tt.v2); result != tt.result {
			t.Fatalf("TestLowerThanOrEqual: expected: %t, found: %t", tt.result, result)
		}
	}
}
