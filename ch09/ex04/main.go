package main

import "fmt"

type T struct {
	c  chan struct{}
	id int
}

// prv->nxtなチャネルを作る
func (prv T) chain(nxt T) {
	for {
		<-prv.c
		// fmt.Println(prv.id)
		nxt.c <- struct{}{}
	}
}

// numで指定した分の連結したチャネルを作成する
func makeChain(cnt int) (chan struct{}, chan struct{}) {
	first := T{make(chan struct{}), 0}
	prv := first
	var last T
	for i := 0; i < cnt; i++ {
		nxt := T{make(chan struct{}), prv.id + 1}
		go prv.chain(nxt)
		prv = nxt
		last = nxt
	}
	return first.c, last.c
}

func main() {
	// 2966段でメモリ不足で死亡
	for i := 1; ; i++ {
		fmt.Println(i)
		first, last := makeChain(i)
		first <- struct{}{}
		<-last
	}
}

/*
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
