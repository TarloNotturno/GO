package DwarfsRafting

/*A company of dwarfs is travelling across the New Zealand. On reaching the Clutha River the dwarfs need to get across, but recent storms have washed away the bridge. Luckily, a small ferry, in the form of a square raft, is operating.

The raft is square and has N rows of seats, numbered from 1 to N. Each row contains N seats, labeled with consecutive letters (A, B, C, etc.). Each seat is identified by a string composed of its row number followed by its column number; for example, "9C" denotes the third seat in the 9th row.

The raft has already been loaded with barrels in some seat positions, and other seats are already occupied by dwarfs. Our company of dwarfs may only take the remaining unoccupied seats. The ferryman wants to accommodate as many dwarfs as possible, but the raft needs to be stable when making the crossing. That is, the following conditions must be satisfied:

the front and back halves of the raft (in terms of the rows of seats) must each contain the same number of dwarfs;
similarly, the left and right sides of the raft (in terms of the columns of seats) must each contain the same number of dwarfs.
You do not have to worry about balancing the barrels; you can assume that their weights are negligible.

For example, a raft of size N = 4 is shown in the following illustration:



Barrels are marked as brown squares, and seats that are already occupied by dwarfs are labeled d.

The positions of the barrels are given in string S. The occupied seat numbers are given in string T. The contents of the strings are separated by single spaces and may appear in any order. For example, in the diagram above, S = "1B 1C 4B 1D 2A" and T = "3B 2D".

In this example, the ferryman can accommodate at most six more dwarfs, as indicated by the green squares in the following diagram:



The raft is then balanced: both left and right halves have the same number of dwarfs (four), and both front and back halves have the same number of dwarfs (also four).

Write a function:

class Solution { public int solution(int N, String S, String T); }

that, given the size of the raft N and two strings S, T that describes the positions of barrels and occupied seats, respectively, returns the maximum number of dwarfs that can fit on the raft. If it is not possible to balance the raft with dwarfs, your function should return -1.

For instance, given N = 4, S = "1B 1C 4B 1D 2A" and T = "3B 2D", your function should return 6, as explained above.

Assume that:

N is an even integer within the range [2..26];
strings S, T consist of valid seat numbers, separated with spaces;
each seat number can appear no more than once in the strings;
no seat number can appear in both S and T simultaneously.
In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.*/

type area struct {
	beginX      int
	endX        int
	beginY      int
	endY        int
	weight      int
	emptySpaces int
}

func myMin(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func fromStringToInt(s string) (places [][]int) {
	i := 0
	number := 0
	for i < len(s) {
		app := int(s[i])
		if app != 32 {
			if app-48 >= 0 && app-48 <= 9 {
				number = number*10 + app - 48
			} else {
				number--
				places = append(places, []int{number, app - 65})
				number = 0
			}
		}
		i++
	}
	return
}

func findMaxCapacity(barrelPlaces [][]int, occupiedPlaces [][]int, a *area) {
	for _, value := range barrelPlaces {
		if value[0] >= a.beginY && value[0] <= a.endY && value[1] >= a.beginX && value[1] <= a.endX {
			a.emptySpaces--
		}
	}
	for _, value := range occupiedPlaces {
		if value[0] >= a.beginY && value[0] <= a.endY && value[1] >= a.beginX && value[1] <= a.endX {
			a.emptySpaces--
			a.weight++
		}
	}
}

func adjustAreas(area1 *area, area2 *area) int {
	var addedDwarf int
	if area1.weight > area2.weight {
		if area1.weight-area2.weight > area2.emptySpaces {
			return -1
		} else {
			addedDwarf = area1.weight - area2.weight
			area2.emptySpaces = area2.emptySpaces - area1.weight + area2.weight
			area1.weight = area2.weight
		}
	} else {
		if area2.weight-area1.weight > area1.emptySpaces {
			return -1
		} else {
			addedDwarf = area2.weight - area1.weight
			area1.emptySpaces = area1.emptySpaces - area2.weight + area1.weight
			area2.weight = area1.weight
		}
	}
	return addedDwarf + 2*myMin(area2.emptySpaces, area1.emptySpaces)
}

func DwarfsRafting(N int, S string, T string) int {
	barrelPlaces := fromStringToInt(S)
	occupiedPlaces := fromStringToInt(T)
	// la zattera moh va divisa in 4 aree
	sxUp := area{beginX: 0, endX: N/2 - 1, beginY: 0, endY: N/2 - 1, emptySpaces: N * N / 4}
	dxUp := area{beginX: N / 2, endX: N - 1, beginY: 0, endY: N/2 - 1, emptySpaces: N * N / 4}
	sxDown := area{beginX: 0, endX: N/2 - 1, beginY: N / 2, endY: N - 1, emptySpaces: N * N / 4}
	dxDown := area{beginX: N / 2, endX: N - 1, beginY: N / 2, endY: N - 1, emptySpaces: N * N / 4}
	findMaxCapacity(barrelPlaces, occupiedPlaces, &sxUp)
	findMaxCapacity(barrelPlaces, occupiedPlaces, &dxUp)
	findMaxCapacity(barrelPlaces, occupiedPlaces, &sxDown)
	findMaxCapacity(barrelPlaces, occupiedPlaces, &dxDown)
	balanceDiag1 := adjustAreas(&sxUp, &dxDown)
	balanceDiag2 := adjustAreas(&dxUp, &sxDown)
	if myMin(balanceDiag1, balanceDiag2) < 0 {
		return -1
	}
	return balanceDiag1 + balanceDiag2
}
