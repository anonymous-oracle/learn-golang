// Package acdc asks if you like rock
package acdc

// Sum adds the ints to return the sum of all the ints
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}

	return s
}
