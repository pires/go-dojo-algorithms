package algos

import (
	"github.com/pires/go-dojo-algorithms/grokking_algorithms/util"
)

const (
	ASC  = 0
	DESC = 1
)

// getSignificantValuePositionInArray returns the position in the array of the lower or higher value stored, based on an order.
// If order is ascending, the position of the lower value will be returned.
// If order is descending, the position of the higher value will be returned.
func getSignificantValuePositionInArray(items []interface{}, order int) int {
	// Start with first position in the array
	v := items[0]
	pos := 0
	// Iterate array and evaluate
	for i := 0; i < len(items); i++ {
		// Respect the order
		if (order == ASC && util.LowerThanOrEqual(items[i], v)) || (order == DESC && util.GreaterThanOrEqual(items[i], v)) {
			v = items[i]
			pos = i
		}
	}

	return pos
}

// sort performs selection sorting and returns a sorted array.
func sort(items []interface{}, order int) []interface{} {
	// Prevent unnecessary execution
	if len(items) <= 1 {
		return items
	}

	// Reserve needed memory
	sorted := make([]interface{}, len(items))

	// Prepare temporary copy to work on top of
	tempItems := make([]interface{}, len(items))
	copy(tempItems, items)

	// Perform selection sort
	for i := 0; i < len(items); i++ {
		// Retrieve significant value position in the array
		pos := getSignificantValuePositionInArray(tempItems, order)
		// Append value to sorted array
		sorted[i] = tempItems[pos]
		// Delete from items to avoid being processed again
		tempItems = append(tempItems[:pos], tempItems[pos+1:]...)
	}

	return sorted
}

// SortInts performs selection sort on an array of integers.
// The result is two arrays sorted in ascending and descending order, respectively.
func SortInts(items []int) (asc []int, desc []int) {
	var interfaceSlice []interface{} = make([]interface{}, len(items))
	for i, d := range items {
		interfaceSlice[i] = d
	}

	asc = make([]int, len(items))
	ascSlice := sort(interfaceSlice, ASC)
	for i, d := range ascSlice {
		asc[i] = d.(int)
	}

	desc = make([]int, len(items))
	descSlice := sort(interfaceSlice, DESC)
	for i, d := range descSlice {
		desc[i] = d.(int)
	}

	return
}

// SortStrings performs selection sort on an array of strings.
// The result is two arrays sorted in ascending and descending order, respectively.
func SortStrings(items []string) (asc []string, desc []string) {
	var interfaceSlice []interface{} = make([]interface{}, len(items))
	for i, d := range items {
		interfaceSlice[i] = d
	}

	asc = make([]string, len(items))
	ascSlice := sort(interfaceSlice, ASC)
	for i, d := range ascSlice {
		asc[i] = d.(string)
	}

	desc = make([]string, len(items))
	descSlice := sort(interfaceSlice, DESC)
	for i, d := range descSlice {
		desc[i] = d.(string)
	}

	return
}
