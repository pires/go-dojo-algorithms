package binary_search

import (
	"sort"
)

// Higher whether v1 is higher than v2.
// While equality works for interface{}, greater than or lower than, doesnt.
func GreaterThan(v1 interface{}, v2 interface{}) bool {
	switch v1.(type) {
	case int:
		return v1.(int) > v2.(int)
	case string:
		return v1.(string) > v2.(string)
	default:
		return v1.(int) > v2.(int)
	}
}

// binarySearch takes an array of items, sorts it and performs binary search on it,
// It returns true plus the position in the list, if the item exists in the list,
// false and zero otherwise.
func binarySearch(items []interface{}, item interface{}) (bool, int) {
	// Prevent underflow
	if len(items) == 0 {
		return false, 0
	}

	// No need to iterate on array with length == 1
	if len(items) == 1 {
		return items[0] == item, 0
	}

	// Initialize coordinate values
	low, mid := 0, 0
	high := len(items) - 1

	// Perform binary search
	for low <= high {
		mid = (low + high) / 2
		if items[mid] == item {
			// Found it
			return true, mid
		}

		// Cut elements to search in (the correct) half
		if GreaterThan(items[mid], item) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false, 0
}

// BinarySearchInt performs binary search on an array of integer values.
func BinarySearchInt(items []int, item int) (bool, int) {
	// Make sure items are sorted
	sort.Ints(items)

	var interfaceSlice []interface{} = make([]interface{}, len(items))
	for i, d := range items {
		interfaceSlice[i] = d
	}

	// Run binary search
	return binarySearch(interfaceSlice, item)
}

// BinarySearchString performs binary search on an array of string values.
func BinarySearchString(items []string, item string) (bool, int) {
	// Make sure items are sorted
	sort.Strings(items)

	var interfaceSlice []interface{} = make([]interface{}, len(items))
	for i, d := range items {
		interfaceSlice[i] = d
	}

	// Run binary search
	return binarySearch(interfaceSlice, item)
}
