package main

import (
	"fmt"
)

type occurencyFound struct {
	pos   int
	occur int
}

/*  FirstUnique --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

A non-empty array A consisting of N integers is given. The unique number is the number that occurs exactly once in array A.

For example, the following array A:

  A[0] = 4
  A[1] = 10
  A[2] = 5
  A[3] = 4
  A[4] = 2
  A[5] = 10
contains two unique numbers (5 and 2).

You should find the first unique number in A. In other words, find the unique number with the lowest position in A.

For above example, 5 is in second position (because A[2] = 5) and 2 is in fourth position (because A[4] = 2). So, the first unique number is 5.

Write a function:

class Solution { public int solution(int[] A); }

that, given a non-empty array A of N integers, returns the first unique number in A. The function should return −1 if there are no unique numbers in A.

For example, given:

  A[0] = 1
  A[1] = 4
  A[2] = 3
  A[3] = 3
  A[4] = 1
  A[5] = 2
the function should return 4. There are two unique numbers (4 and 2 occur exactly once). The first one is 4 in position 1 and the second one is 2 in position 5. The function should return 4 bacause it is unique number with the lowest position.

Given array A such that:

  A[0] = 6
  A[1] = 4
  A[2] = 4
  A[3] = 6
the function should return −1. There is no unique number in A (4 and 6 occur more than once).

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [0..1,000,000,000].

Result 100%*/

func FirstUnique(A []int) int {
	firstUnique := -1
	occurrencies := make(map[int]*occurencyFound)
	for i, value := range A {
		addre, ok := occurrencies[value]
		if ok {
			addre.occur++

		} else {
			occurrencies[value] = &occurencyFound{pos: i, occur: 1}
		}

	}
	minPosition := len(A)
	for value, occurr := range occurrencies {
		if occurr.occur < 2 && occurr.pos < minPosition {
			firstUnique = value
			minPosition = occurr.pos
		}
	}
	return firstUnique
}

/* StrSymmetryPoint ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Write a function:

class Solution { public int solution(String S); }

that, given a string S, returns the index (counting from 0) of a character such that the part of the string to the left of that character is a reversal of the part of the string to its right. The function should return −1 if no such index exists.

Note: reversing an empty string (i.e. a string whose length is zero) gives an empty string.

For example, given a string:

"racecar"

the function should return 3, because the substring to the left of the character "e" at index 3 is "rac", and the one to the right is "car".

Given a string:

"x"

the function should return 0, because both substrings are empty.

Write an efficient algorithm for the following assumptions:

the length of string S is within the range [0..2,000,000].

Result 100%*/

func StrSymmetryPoint(S string) int {
	i := 0
	N := len(S) - 1
	if N%2 == 1 || N < 0 {
		return -1
	}
	for 2*i < N {
		if S[i] == S[N-1] {
			i++
		} else {
			return -1
		}
	}
	return i
}

func newTree(x int, l *Tree, r *Tree) *Tree {
	var t Tree
	t.X = x
	t.L = l
	t.R = r
	return &t
}

func myMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type branch struct {
	lenghtBranch int
	val          int
	Leaf         *Tree
}

func TreeHeight(T *Tree) int {
	sol := 0
	queuBranch := make([]branch, 0)

	if T.L != nil {
		queuBranch = append(queuBranch, branch{lenghtBranch: 0, Leaf: T.L})
	}
	if T.R != nil {
		queuBranch = append(queuBranch, branch{lenghtBranch: 0, Leaf: T.R})
	}
	for len(queuBranch) > 0 {
		if queuBranch[0].Leaf.L == nil && queuBranch[0].Leaf.R == nil {
			queuBranch[0].lenghtBranch = queuBranch[0].lenghtBranch + 1
		}
		if queuBranch[0].Leaf.L != nil {
			queuBranch = append(queuBranch, branch{lenghtBranch: queuBranch[0].lenghtBranch + 1, Leaf: queuBranch[0].Leaf.L})
		}
		if queuBranch[0].Leaf.R != nil {
			queuBranch = append(queuBranch, branch{lenghtBranch: queuBranch[0].lenghtBranch + 1, Leaf: queuBranch[0].Leaf.R})
		}
		if queuBranch[0].lenghtBranch > sol {
			sol = queuBranch[0].lenghtBranch
		}
		queuBranch = queuBranch[1:]
	}
	return sol

}

/* ArrayInversionCount -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

An array A consisting of N integers is given. An inversion is a pair of indexes (P, Q) such that P < Q and A[Q] < A[P].

Write a function:

class Solution { public int solution(int[] A); }

that computes the number of inversions in A, or returns −1 if it exceeds 1,000,000,000.

For example, in the following array:

 A[0] = -1 A[1] = 6 A[2] = 3
 A[3] =  4 A[4] = 7 A[5] = 4
there are four inversions:

   (1,2)  (1,3)  (1,5)  (4,5)
so the function should return 4.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].

Result 72% solution not optimal */

func ArrayInversionCountNotOptimal(A []int) int {
	N := len(A)
	sol := 0
	for p := 0; p < N-1; p++ {
		for q := p + 1; q < N; q++ {
			if A[q] < A[p] {
				sol++
			}
		}
	}
	return sol
}

func ArrayInversionCount(A []int) int {
	N := len(A)
	sol := 0
	for i := 2; i <= N; i++ {
		check := A[N-i] > A[N-i+1]
		j := 0
		for check {
			sol++
			j++
			if j < i-1 {
				check = (A[N-i] > A[N-i+1+j])
			} else {
				check = false
			}
		}
		heapSort(A[N-i:])
	}
	return sol
}

/* DisappearingPairs ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
A string S containing only the letters "A", "B" and "C" is given. The string can be transformed by removing one occurrence of "AA", "BB" or "CC".

Transformation of the string is the process of removing letters from it, based on the rules described above. As long as at least one rule can be applied,
the process should be repeated. If more than one rule can be used, any one of them could be chosen.

Write a function:

class Solution { public String solution(String S); }

that, given a string S consisting of N characters, returns any string that can result from a sequence of transformations as described above.

For example, given string S = "ACCAABBC" the function may return "AC", because one of the possible sequences of transformations is as follows:



Also, given string S = "ABCBBCBA" the function may return "", because one possible sequence of transformations is:



Finally, for string S = "BABABA" the function must return "BABABA", because no rules can be applied to string S.

Write an efficient algorithm for the following assumptions:

the length of string S is within the range [0..50,000];
string S is made only of the following characters: 'A', 'B' and/or 'C'.

Solution 100% */

func DisappearingPairs(S string) string {
	queuToCancel := make([]int, 0)
	if len(S) > 1 {
		queuToCancel = append(queuToCancel, int(S[0]))
	} else {
		return S
	}
	index := 1
	N := len(S)
	for len(queuToCancel) > 0 && index < N {
		if queuToCancel[len(queuToCancel)-1] == int(S[index]) {
			queuToCancel = queuToCancel[:(len(queuToCancel) - 1)]
			if len(queuToCancel) == 0 && index < N-1 {
				index++
				queuToCancel = append(queuToCancel, int(S[index]))
			}
		} else {
			queuToCancel = append(queuToCancel, int(S[index]))
		}
		index++
	}
	var sol string
	for i := 0; i < len(queuToCancel); i++ {
		if queuToCancel[i] == 65 {
			sol = sol + "A"
		}
		if queuToCancel[i] == 66 {
			sol = sol + "B"
		}
		if queuToCancel[i] == 67 {
			sol = sol + "C"
		}
	}
	return sol
}

/* Calling function */
func main() {
	FirstUnique([]int{6, 4, 4, 6})
	FirstUnique([]int{1, 4, 3, 3, 1, 2})
	FirstUnique([]int{})
	fmt.Println(StrSymmetryPoint("racecar"))
	fmt.Println(StrSymmetryPoint("raccar"))
	fmt.Println(StrSymmetryPoint("fortinite"))
	fmt.Println(StrSymmetryPoint("legenegel"))
	fmt.Println(StrSymmetryPoint("'a'"))
	fmt.Println(StrSymmetryPoint(""))
	t := newTree(5 /*1 L*/, newTree(3, newTree(20, nil, nil), newTree(21, nil, nil)) /*1 R*/, newTree(10, newTree(1, nil, nil), nil))
	fmt.Println(TreeHeight(t))
	A := []int{1, 4, 5, 3, 1, 2}
	fmt.Println(A)
	fmt.Println(ArrayInversionCountNotOptimal(A))
	fmt.Println(ArrayInversionCount(A))
	A = []int{-1, 6, 3, 4, 7, 4}
	fmt.Println(ArrayInversionCount(A))
	fmt.Println(DisappearingPairs("ABCBBCBA"))
	fmt.Println(DisappearingPairs("AABBBBCBAABCBABCCCA"))
	fmt.Println(DisappearingPairs("A"))
	fmt.Println(DisappearingPairs(""))

}
