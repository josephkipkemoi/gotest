package math

// Finds the average of a series of number
func Average(xs []float64) float64 {

 	total := float64(0)
	if float64(len(xs)) == 0 {
		return 0
	}
	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}

// Finds the maximum number of a series of number
func Max(xs []float64) float64 {

	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}

	return max
}

// Finds the minimum number of a series of number
func Min(xs []float64) float64 {

	min := xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	
	return min
}

// Find the fibonacci sequence of number
func Fib(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}