package kata

func Tribonacci(signature [3]float64, n int) (r []float64) {
	r = signature[:]

	for i := len(r); i < n; i++ {
		r = append(r, r[i-1]+r[i-2]+r[i-3])
	}

	return r[:n]
}
