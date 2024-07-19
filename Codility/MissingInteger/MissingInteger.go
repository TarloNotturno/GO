package MissingInteger

/*Write a function:

func MissingInteger(input_array []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
*/

import (
	"sort"
)

func MissingInteger(input_array []int) int {
	//input_array = A
	missingInt := 0
	if len(input_array) <= 0 {
		return 1
	} else {
		sort.Ints(input_array)
		for _, a := range input_array {
			if a != missingInt+1 && a != missingInt && missingInt >= 0 && a > 0 {
				return missingInt + 1
			} else if a > 0 {
				missingInt = a
			}
		}
		missingInt = missingInt + 1
	}
	if missingInt <= 0 {
		return 1
	}
	return missingInt
}
