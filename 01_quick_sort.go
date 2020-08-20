package ca25l

// QuickSort implement a quick sort.
func QuickSort(input []int) {
	if len(input) <= 1 {
		return
	}

	first, last := 0, len(input)-1
	base := input[first]
	for first < last {
		for first < last && input[last] >= base {
			last--
		}
		if first < last {
			input[first] = input[last]
			first++
		}
		for first < last && input[first] <= base {
			first++
		}
		if first < last {
			input[last] = input[first]
			last--
		}
	}
	input[first] = base

	if first > 1 {
		QuickSort(input[:first])
	}
	if last < len(input)-2 {
		QuickSort(input[last+1:])
	}

	return
}
