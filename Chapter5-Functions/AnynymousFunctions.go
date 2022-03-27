package main

func multiplyBy(factor int, a []int) []int {
	for i := 0; i < len(a); i++ {
		func(factor int, j int) {
			a[j] = a[j] * factor
		}(factor, i)
	}
	return a
}
