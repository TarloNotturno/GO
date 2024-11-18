package main

type peakAndDepth struct {
	listOfPeak  []int
	listOfDepth []int
}

// myMax and myMin finds the min and max between two integer
func myMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func myMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// function that finds points higher than previous
func FindLocalMax(A []int, pointOfInterest *peakAndDepth) {
	var (
		currentMin   int
		previousPeak int
		deltaH       int
		oldDeltaH    int
	)

	pointOfInterest.listOfDepth = append(pointOfInterest.listOfDepth, 0)
	for i, val := range A {
		// find the list of local peaks
		oldDeltaH = deltaH
		deltaH = val - previousPeak
		previousPeak = val

		if deltaH < 0 && oldDeltaH > 0 {
			pointOfInterest.listOfPeak = append(pointOfInterest.listOfPeak, i-1)
			if len(pointOfInterest.listOfPeak) > 1 {
				pointOfInterest.listOfDepth = append(pointOfInterest.listOfDepth, currentMin)
			}
			currentMin = i
		}

		if val < A[currentMin] {
			currentMin = i
		}
	}

	//last point in this way is not added to the peaks
	if A[len(A)-1] > A[len(A)-2] {
		pointOfInterest.listOfPeak = append(pointOfInterest.listOfPeak, len(A)-1)
		pointOfInterest.listOfDepth = append(pointOfInterest.listOfDepth, currentMin)
	}

}

// keep only peaks
func ClearPeaks(A []int, pointOfInterest *peakAndDepth) {
	lastPeak := pointOfInterest.listOfPeak[0]
	cleanedPeaks := make([]int, 0)
	cleanedDepths := make([]int, 0)
	minDepth := pointOfInterest.listOfDepth[0]
	peakQueu := make([]int, 0)
	insertedPeak := make(map[int]int)

	// if the current value is higher than the last peak found means it is a peak
	// otherwise i add it to a queu that will be checked in the opposite direction,
	// from end to beginning
	for i, currPeak := range pointOfInterest.listOfPeak {

		if A[minDepth] >= A[pointOfInterest.listOfDepth[i]] {
			minDepth = pointOfInterest.listOfDepth[i]
		}

		if A[lastPeak] <= A[currPeak] {
			cleanedPeaks = append(cleanedPeaks, currPeak)
			cleanedDepths = append(cleanedDepths, minDepth)
			minDepth = currPeak
			lastPeak = currPeak
			peakQueu = []int{}
			insertedPeak[currPeak] = 2
		} else {
			peakQueu = append(peakQueu, currPeak)
		}
	}

	// start from end of vector, if current value is higher than last peak add it
	// It is checked in the queu populated before
	j := len(peakQueu) - 1
	app := []int{}
	addDepth := []int{}
	if j >= 0 {
		lastPeak = peakQueu[j]
		for j >= 0 {
			if A[lastPeak] <= A[peakQueu[j]] {
				gino := make([]int, 0)
				gino = append(gino, peakQueu[j])
				app = append(gino, app...)
				lastPeak = peakQueu[j]
				insertedPeak[peakQueu[j]] = 1
			} else {
				insertedPeak[peakQueu[j]] = 0
			}
			j--
		}

		// for each peak add an abyss
		currmin := pointOfInterest.listOfPeak[0]
		for j = 0; j < len(A); j++ {
			if A[currmin] >= A[j] {
				currmin = j
			}

			if insertedPeak[j] == 1 {
				addDepth = append(addDepth, currmin)
				currmin = j
			} else if insertedPeak[j] == 2 {
				currmin = j

			}
		}
	}

	cleanedPeaks = append(cleanedPeaks, app...)
	cleanedDepths = append(cleanedDepths, addDepth...)
	pointOfInterest.listOfPeak = cleanedPeaks
	pointOfInterest.listOfDepth = cleanedDepths
}
