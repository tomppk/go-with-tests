package main

// range iterates over and array
// every time it is called returns two values, the index and the value
// we are choosing to ignore index by using _ blank identifier
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes in varying number of slices and returns a new slice containing
// the totals for each slice passed in.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails calculates the totals of each slice, excluding the first item (the "head")
// numbers[1:] means "take from 1 to the end"
func SumAllTails(numbersToSum ...[]int) []int {
		var sums []int
		for _, numbers := range numbersToSum {
			if len(numbers) == 0 {
				sums = append(sums, 0)
			} else {
				tail := numbers[1:]
				sums = append(sums, Sum(tail))
			}


		}

		return sums
	}
