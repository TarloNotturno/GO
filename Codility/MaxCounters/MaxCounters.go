package MaxCounters

func MaxCounters(N int, A []int) []int {
	result := make([]int, N)

	if cap(A) == 0 {
		return result
	}
	var max_Val int
	max_ValOld := 0
	max_Val = 0
	for i := 0; i < cap(A); i++ {
		if A[i] == N+1 {
			max_ValOld = max_Val
		} else {
			if result[A[i]-1] < max_ValOld {
				result[A[i]-1] = max_ValOld
			}
			result[A[i]-1] = result[A[i]-1] + 1
			if result[A[i]-1] > max_Val {
				max_Val = result[A[i]-1]
			}
		}
	}
	for i := 0; i < N; i++ {
		if result[i] < max_ValOld {
			result[i] = max_ValOld
		}
	}

	return result
}
