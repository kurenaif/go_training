package main

import "testing"

func BenchmarkMandebrot(b *testing.B) {
	for count := 0; count < b.N; count++ {
		calcImage()
	}
}

func BenchmarkMandebrotParallel1(b *testing.B) {
	for count := 0; count < b.N; count++ {
		calcImageParallel(1)
	}
}

func BenchmarkMandebrotParallel2(b *testing.B) {
	for count := 0; count < b.N; count++ {
		calcImageParallel(2)
	}
}

func BenchmarkMandebrotParallel4(b *testing.B) {
	for count := 0; count < b.N; count++ {
		calcImageParallel(4)
	}
}

func BenchmarkMandebrotParallel8(b *testing.B) {
	for count := 0; count < b.N; count++ {
		calcImageParallel(8)
	}
}

/*
                          ./+o+-       kurenaif@kurenaif-home
                  yyyyy- -yyyyyy+      OS: Ubuntu 18.10 cosmic
               ://+//////-yyyyyyo      Kernel: x86_64 Linux 4.18.0-25-generic
           .++ .:/++++++/-.+sss/`      Uptime: 3h 17m
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
               /osyyyyyyo++ooo+++/     RAM: 5043MiB / 15989MiB
                   ````` +oo+++o\:
                          `oo++.

*/
