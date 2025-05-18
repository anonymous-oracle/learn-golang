// package mymath provides ACME inc math solutions
package mymath

// sums adds an unlimited number of values of type int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
