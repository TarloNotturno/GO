package main

/*SocksLaundering   https://app.codility.com/programmers/trainings/3/socks_laundering/
Bob is about to go on a trip. But first he needs to take care of his supply of socks. Each sock has its own color.
Bob wants to take as many pairs of clean socks as possible (both socks in the pair should be of the same color).

Socks are divided into two drawers: clean and dirty socks. Bob has time for only one laundry and his washing machine can clean at
most K socks. He wants to pick socks for laundering in such a way that after washing he will have a maximal number of clean,
same-colored pairs of socks. It is possible that some socks cannot be paired with any other sock, because Bob may have lost
some socks over the years.

Bob has exactly N clean and M dirty socks, which are described in arrays C and D, respectively. The colors of the socks are
represented as integers (equal numbers representing identical colors).

For example, given four clean socks and five dirty socks:
clean: 1 2 1 1
dirty: 1 4 3 2 4

If Bob's washing machine can clean at most K = 2 socks, then he can take a maximum of three pairs of clean socks. He can wash one
red sock and one green sock, numbered 1 and 2 respectively. Then he will have two pairs of red socks and one pair of green socks.

Write a function:

class Solution { public int solution(int K, int[] C, int[] D); }

that, given an integer K (the number of socks that the washing machine can clean), two arrays C and D (containing the color
representations of N clean and M dirty socks respectively), returns the maximum number of pairs of socks that Bob can take on the
trip.

For example, given K = 2, C = [1, 2, 1, 1] and D = [1, 4, 3, 2, 4], the function should return 3, as explained above.

Assume that:

K is an integer within the range [0..50];
each element of arrays C and D is an integer within the range [1..50];
C and D are not empty and each of them contains at most 50 elements.
In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.


*/

// K capacity washing machine
// C list of clean socks
// D list of dirty socks

func myMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func SocksLaundering(K int, C []int, D []int) int {
	var pairOfSocks int
	capacityWashingmachine := K

	cleanSocks := make(map[int]int)
	dirtySocks := make(map[int]int)

	// find couple of clean socks ----------------------------------------------------------
	for _, value := range C {
		cleanSocks[value] = cleanSocks[value] + 1
	}

	// find list of dirty socks ------------------------------------------------------------
	for _, value := range D {
		dirtySocks[value] = dirtySocks[value] + 1
	}

	//add all odd number of clean socks
	for i, value := range cleanSocks {
		if value%2 == 0 {
			pairOfSocks = pairOfSocks + value/2
			cleanSocks[i] = 0
		} else {
			pairOfSocks = pairOfSocks + myMax(value-1, 0)/2
			cleanSocks[i] = 1
		}
	}

	// add the remaining clean socks pairing with a dirty one if possible
	for i, value := range cleanSocks {
		if value > 0 {
			valDirt, ok := dirtySocks[i]
			if ok {
				if valDirt > 0 && capacityWashingmachine > 0 {
					capacityWashingmachine--
					pairOfSocks = pairOfSocks + 1
					dirtySocks[i] = dirtySocks[i] - 1
					cleanSocks[i] = cleanSocks[i] - 1
				}
			}
		}
	}

	// count number of dirty odds
	var dirtyCouples int
	for i, value := range dirtySocks {
		var coupleToAdd int
		if value%2 == 0 {
			coupleToAdd = myMin(value/2, int(capacityWashingmachine/2))
			dirtySocks[i] = dirtySocks[i] - coupleToAdd
		} else {
			if value-1 > 0 {
				coupleToAdd = myMin((value-1)/2, int(capacityWashingmachine/2))
				dirtySocks[i] = dirtySocks[i] - coupleToAdd
			}
		}

		if coupleToAdd > 0 {
			dirtyCouples = dirtyCouples + coupleToAdd
			dirtySocks[i] = 0
			capacityWashingmachine = capacityWashingmachine - coupleToAdd*2
			coupleToAdd = 0
		}
	}

	pairOfSocks = pairOfSocks + dirtyCouples

	return pairOfSocks
}
