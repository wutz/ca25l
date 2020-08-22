package ca25l

// PerfQuickSort implement a performanced quick sort
func PerfQuickSort(input []int) {
	for len(input) > 1 {
		first, last := 0, len(input)-1
		base := middle(input[first], input[last/2], input[last])
		for first < last {
			for input[first] < base {
				first++
			}
			for input[last] > base {
				last--
			}
			if first < last {
				input[first], input[last] = input[last], input[first]
			}
		}
		PerfQuickSort(input[last+1:])
		input = input[:first]
	}
}

func middle(a, b, c int) int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	return b
}
