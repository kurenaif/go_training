package main

import (
	"math/big"
)

func main() {
	bigTwo := new(big.Float).SetFloat64(2.0)
	// bigFour := new(big.Float).SetFloat64(4.0)
	vx := new(big.Float).SetFloat64(5.0)
	vy := new(big.Float).SetFloat64(2.0)

	nvx := new(big.Float).Sub(new(big.Float).Mul(vx, vx), new(big.Float).Mul(vy, vy))
	nvy := new(big.Float).Mul(new(big.Float).Mul(bigTwo, vx), vy)
	vx = nvx
	vy = nvy
	a, _ := vx.Float64()
	b, _ := vy.Float64()
	println(a, b)
}
