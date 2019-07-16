package main

import "fmt"

func gcd(lhs, rhs int) int {
	if rhs == 0 {
		return lhs
	}
	return gcd(rhs, lhs%rhs)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(s, 2)
	fmt.Println(s)
}

func rotate(s []int, offset int) {
	n := len(s)
	startMax := gcd(n, offset)
	for start := 0; start < startMax; start++ {
		cnt := n / startMax
		last := s[((cnt-1)*offset+start)%n]
		for i := cnt - 1; i >= 1; i-- {
			l := ((i-1)*offset + start) % n
			r := (i*offset + start) % n
			s[r] = s[l]
		}
		s[start] = last
	}
}
