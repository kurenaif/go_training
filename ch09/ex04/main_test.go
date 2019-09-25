package main

import "testing"

func BenchmarkChannel1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		first, last := makeChain(1)
		first <- struct{}{}
		<-last
	}
}

func BenchmarkChannel10(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		first, last := makeChain(10)
		first <- struct{}{}
		<-last
	}
}

func BenchmarkChannel100(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		first, last := makeChain(100)
		first <- struct{}{}
		<-last
	}
}

func BenchmarkChannel1000(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		first, last := makeChain(1000)
		first <- struct{}{}
		<-last
	}
}

/*
goos: linux
goarch: amd64
pkg: go_training/ch09/ex04
BenchmarkChannel1-8      	 1000000	      2527 ns/op	     716 B/op	       4 allocs/op
BenchmarkChannel10-8     	  100000	     22230 ns/op	    6289 B/op	      31 allocs/op
BenchmarkChannel100-8    	   10000	    160425 ns/op	   62694 B/op	     301 allocs/op
BenchmarkChannel1000-8   	    1000	   1503047 ns/op	  610813 B/op	    3001 allocs/op
PASS
ok  	go_training/ch09/ex04	18.496s

                          ./+o+-       kurenaif@kurenaif-home
                  yyyyy- -yyyyyy+      OS: Ubuntu 18.10 cosmic
               ://+//////-yyyyyyo      Kernel: x86_64 Linux 4.18.0-25-generic
           .++ .:/++++++/-.+sss/`      Uptime: 1m
         .:++o:  /++++++++/:--:/-      Packages: 2830
        o:+o+:++.`..```.-/oo+++++/     Shell: fish 3.0.2
       .:+o:+o/.          `+sssoo+/    Resolution: 5760x2160
  .++/+:+oo+o:`             /sssooo.   DE: GNOME 
 /+++//+:`oo+o               /::--:.   WM: GNOME Shell
 \+/+o+++`o++o               ++////.   WM Theme: Adwaita
  .++.o+++oo+:`             /dddhhh.   GTK Theme: Yaru [GTK2/3]
       .+.o+oo:.          `oddhhhh+    Icon Theme: Yaru
        \+.++o+o``-````.:ohdhhhhh+     Font: Ubuntu 11
         `:o+++ `ohhhhhhhhyo++os:      CPU: Intel Core i7-4790 @ 8x 4GHz [27.8°C]
           .o:`.syhhhhhhh/.oo++o`      GPU: GeForce GTX 970
               /osyyyyyyo++ooo+++/     RAM: 1597MiB / 15989MiB
                   ````` +oo+++o\:    
                          `oo++.      
*/
