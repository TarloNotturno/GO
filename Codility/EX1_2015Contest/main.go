package main

import "fmt"

/* LongestPassword -----------------------------------------------------------------------------------------------------------------------------------------------------------------

You would like to set a password for a bank account. However, there are three restrictions on the format of the password:

it has to contain only alphanumerical characters (a−z, A−Z, 0−9);
there should be an even number of letters;
there should be an odd number of digits.
You are given a string S consisting of N characters. String S can be divided into words by splitting it at, and removing, the spaces.
The goal is to choose the longest word that is a valid password. You can assume that if there are K spaces in string S then there are
exactly K + 1 words.

For example, given "test 5 a0A pass007 ?xy1", there are five words and three of them are valid passwords: "5", "a0A" and "pass007".
Thus the longest password is "pass007" and its length is 7. Note that neither "test" nor "?xy1" is a valid password, because "?" is not
an alphanumerical character and "test" contains an even number of digits (zero).

Write a function:

func Solution(S string) int

that, given a non-empty string S consisting of N characters, returns the length of the longest word from the string that is a valid
password. If there is no such word, your function should return −1.

For example, given S = "test 5 a0A pass007 ?xy1", your function should return 7, as explained above.

Assume that:

N is an integer within the range [1..200];
string S consists only of printable ASCII characters and spaces.
In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.
*/

func LongestPassword(S string) int {
	//solution var
	numberOfPassw := -1

	countLetters := 0
	countDigits := 0
	// if I have a not acceptable element the password is not ok and so i cannot
	// consider it
	psswNotOk := false
	// space is equal to 32
	// numbers are from 48 to 57
	// letters from 97 to 122
	// capital letters from 65 to 90
	for i := 0; i < len(S); i++ {
		if S[i] == 32 {
			if countLetters%2 == 0 && countDigits%2 == 1 && !psswNotOk {
				if countLetters+countDigits > numberOfPassw {
					numberOfPassw = countLetters + countDigits
				}
			}
			countLetters = 0
			countDigits = 0
			psswNotOk = false
		} else if (S[i] >= 97 && S[i] <= 122) || (S[i] >= 65 && S[i] <= 90) {
			countLetters++
		} else if S[i] >= 48 && S[i] <= 57 {
			countDigits++
		} else {
			psswNotOk = true
		}
	}
	if countLetters%2 == 0 && countDigits%2 == 1 && !psswNotOk {
		if countLetters+countDigits > numberOfPassw {
			numberOfPassw = countLetters + countDigits
		}
	}
	return numberOfPassw
}

/* FloodDepth  --------------------------------------------------------------------------------------------------------------------------
https://app.codility.com/programmers/trainings/1/flood_depth/
You are helping a geologist friend investigate an area with mountain lakes. A recent heavy rainfall has flooded these lakes and their
water levels have reached the highest possible point. Your friend is interested to know the maximum depth in the deepest part of these lakes.

We simplify the problem in 2-D dimensions. The whole landscape can be divided into small blocks and described by an array A of length N.
Each element of A is the altitude of the rock floor of a block (i.e. the height of this block when there is no water at all).
After the rainfall, all the low-lying areas (i.e. blocks that have higher blocks on both sides) are holding as much water as possible.
You would like to know the maximum depth of water after this entire area is flooded. You can assume that the altitude outside this area
is zero and the outside area can accommodate infinite amount of water.

For example, consider array A such that:

    A[0] = 1
    A[1] = 3
    A[2] = 2
    A[3] = 1
    A[4] = 2
    A[5] = 1
    A[6] = 5
    A[7] = 3
    A[8] = 3
    A[9] = 4
    A[10] = 2
The following picture illustrates the landscape after it has flooded:



The gray area is the rock floor described by the array A above and the blue area with dashed lines represents the water
filling the low-lying areas with maximum possible volume. Thus, blocks 3 and 5 have a water depth of 2 while blocks 2, 4, 7 and 8 have a
water depth of 1. Therefore, the maximum water depth of this area is 2.

Write a function:

class Solution { public int solution(int[] A); }

that, given a non-empty array A consisting of N integers, returns the maximum depth of water.

Given array A shown above, the function should return 2, as explained above.

For the following array:

    A[0] = 5
    A[1] = 8
the function should return 0, because this landscape cannot hold any water.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [1..100,000,000].
*/

func FloodDepth(A []int) int {
	if len(A) <= 2 {
		return 0
	}
	var (
		pointOfInterest peakAndDepth
		highestDiff     int
	)
	pointOfInterest.listOfPeak = make([]int, 0)

	// find the list of local peaks
	FindLocalMax(A, &pointOfInterest)

	if len(pointOfInterest.listOfPeak) > 0 {
		//remove the peaks that do not brings to a "lake"
		ClearPeaks(A, &pointOfInterest)
		i := 0
		j := 1
		for j < len(pointOfInterest.listOfDepth) {
			// check between to peaks which is the higher and subtract the hole
			depth := myMin(A[pointOfInterest.listOfPeak[i]], A[pointOfInterest.listOfPeak[j]]) - A[pointOfInterest.listOfDepth[j]]
			highestDiff = myMax(depth, highestDiff)
			i++
			j++

		}
	}
	return highestDiff
}

func main() {
	fmt.Println(LongestPassword("test 5 a0A pass007 ?xy1"))
	fmt.Println(LongestPassword("apex911"))
	fmt.Println(LongestPassword("1"))
	fmt.Println(FloodDepth([]int{1, 3, 2, 1, 2, 1, 5, 3, 3, 4, 2}), "expected", 2)
	fmt.Println(FloodDepth([]int{20, 4, 10, 1, 3, 5, 7}), "expected", 6)
	fmt.Println(FloodDepth([]int{1, 9, 8, 7, 4, 8, 7, 10, 31, 2}), "expected", 5)
	fmt.Println(FloodDepth([]int{3, 1, 10, 4, 20}), "expected", 6)
	fmt.Println(FloodDepth([]int{5000, 2000, 5000, 1, 3000, 4000, 5000, 1, 2, 2, 1, 2, 5000}), "expected", 4999)
}
