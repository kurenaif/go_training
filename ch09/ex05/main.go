package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	cnt := 0
	done := make(chan struct{})

	go func(done chan struct{}) {
		tick := time.Tick(1 * time.Second)
		start := time.Now()
		for {
			<-tick
			fmt.Printf("%v,%d\n", time.Since(start), cnt)
		}
		close(done)
	}(done)

	go func(ch1 chan struct{}, ch2 chan struct{}) {
		for {
			<-ch1
			cnt++
			ch2 <- struct{}{}
		}
	}(ch1, ch2)

	go func(ch1 chan struct{}, ch2 chan struct{}) {
		for {
			ch1 <- struct{}{}
			<-ch2
		}
	}(ch1, ch2)
	<-done
}

/*
1.00006211s,2123020
2.000083713s,4121554
3.000080204s,6155867
4.000068363s,8212511
5.000070256s,10274536
6.000067769s,12317915
7.000072089s,14319499
8.000067865s,16375361
9.000086696s,18399173
10.00006999s,20387574
^Csignal: interrupt
*/
