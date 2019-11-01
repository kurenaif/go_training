package main

func main() {
}

/*
BenchmarkPopCount11-8                   1000000000           2.81 ns/op
BenchmarkPopCount00-8                   1000000000           2.82 ns/op
BenchmarkPopCount00rand-8               100000000           20.9 ns/op
BenchmarkPopCountRand-8                 100000000           20.1 ns/op
BenchmarkPopCountLoop11-8               100000000           17.6 ns/op
BenchmarkPopCountLoop00-8               100000000           18.6 ns/op
BenchmarkPopCountLoop00rand-8           50000000            28.9 ns/op
BenchmarkPopCountLoopRand-8             50000000            32.4 ns/op
BenchmarkPopCountBitShift00-8           50000000            37.9 ns/op
BenchmarkPopCountBitShift11-8           30000000            37.7 ns/op
BenchmarkPopCountBitShift00rand-8       30000000            51.0 ns/op
BenchmarkPopCountBitShiftRand-8         30000000            53.1 ns/op
BenchmarkPopCountLSB00-8                2000000000           1.79 ns/op
BenchmarkPopCountLSB11-8                30000000            41.5 ns/op
BenchmarkPopCountLSB00rand-8            100000000           20.2 ns/op
BenchmarkPopCountLSBRand-8              30000000            44.6 ns/op
*/
