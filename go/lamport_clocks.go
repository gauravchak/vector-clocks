package lamport_clocks

func GetTimestamps(input [][]Event) [][]int {
	// uncomment to use simple readable solution
	// return GetTimestampsSimple(input)

	// uncomment to test O(num events) solution
	// return GetTimestampsOptimized(input)

	// uncomment to test solution that uses essentially topological sort.
	return GetTimestampsTopSort(input)
}
