package inversion_number

// InversionNumber returns an inversion number of array A.
// Time complexity: O(|A|log|A|)
func InversionNumber(A []int64) int64 {
	var _rec func(A []int64, l, r int) int64

	// _rec sorts an array A
	_rec = func(A []int64, l, r int) int64 {
		if r-l == 1 {
			return 0
		}

		// divide and conquer
		mid := (l + r) / 2
		res := _rec(A, l, mid) + _rec(A, mid, r)

		// collect left and right elements separately
		L, R := []int64{}, []int64{}
		for i := l; i < r; i++ {
			if i < mid {
				L = append(L, A[i])
			} else {
				R = append(R, A[i])
			}
		}

		// process like shakutori method
		cur, j := l, 0
		for i := 0; i < len(L); i++ {
			for j < len(R) && L[i] > R[j] {
				A[cur] = R[j]
				cur, j = cur+1, j+1
			}

			res += int64(j) // update result

			A[cur] = L[i]
			cur++
		}
		for ; j < len(R); j++ {
			A[cur] = R[j]
			cur++
		}

		return res
	}

	B := make([]int64, len(A))
	copy(B, A)
	return _rec(B, 0, len(B))
}
