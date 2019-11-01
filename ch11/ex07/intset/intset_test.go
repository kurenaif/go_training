package intset

import (
	"go_training/ch11/ex07/mapset"
	"testing"
)

// --------------------------------------------------

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s IntSet
		for j := 0; j < 100000; j++ {
			s.Add(j)
		}
	}

}

func BenchmarkFastAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s IntSet
		for j := 0; j < 100000; j++ {
			s.FastAdd(j)
		}
	}
}

func BenchmarkMapAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s mapset.IntSet
		for j := 0; j < 100000; j++ {
			s.Add(j)
		}
	}
}

// --------------------------------------------------

func BenchmarkUnionWith(b *testing.B) {
	var s1, s2 IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.UnionWith(&s2)
	}

}

func BenchmarkFastUnionWith(b *testing.B) {
	var s1, s2 IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.FastUnionWith(&s2)
	}
}

func BenchmarkMapUnionWith(b *testing.B) {
	var s1, s2 mapset.IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 mapset.IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.UnionWith(&s2)
	}

}

// --------------------------------------------------

var result int

func BenchmarkLen(b *testing.B) {
	var s IntSet
	for j := 0; j < 100000; j++ {
		s.Add(j)
	}
	a := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a += s.Len()
	}
	result = a
}

func BenchmarkMapLen(b *testing.B) {
	var s mapset.IntSet
	for j := 0; j < 100000; j++ {
		s.Add(j)
	}
	a := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a += s.Len()
	}
	result = a
}

// --------------------------------------------------

func BenchmarkRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := IntSet{}
		b.StopTimer()
		for j := 0; j < 100000; j++ {
			s.Add(j)
		}
		b.StartTimer()
		for j := 0; j < 100000; j++ {
			s.Remove(j)
		}
	}
}

func BenchmarkFastRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := IntSet{}
		b.StopTimer()
		for j := 0; j < 100000; j++ {
			s.Add(j)
		}
		b.StartTimer()
		for j := 0; j < 100000; j++ {
			s.FastRemove(j)
		}
	}
}

func BenchmarkMapRemove(b *testing.B) {
	s := mapset.IntSet{}
	b.StopTimer()
	for j := 0; j < 100000; j++ {
		s.Add(j)
	}
	b.StartTimer()
	for j := 0; j < 100000; j++ {
		s.Remove(j)
	}
}

// --------------------------------------------------

func BenchmarkIntersectWith(b *testing.B) {
	var s1, s2 IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.IntersectWith(&s2)
	}
}

func BenchmarkFastIntersectWith(b *testing.B) {
	var s1, s2 IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.FastIntersectWith(&s2)
	}
}

func BenchmarkMapIntersectWith(b *testing.B) {
	var s1, s2 mapset.IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 mapset.IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.IntersectWith(&s2)
	}

}

// --------------------------------------------------

func BenchmarkDifferenceWith(b *testing.B) {
	var s1, s2 IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.DifferenceWith(&s2)
	}
}

func BenchmarkFastDifferenceWith(b *testing.B) {
	var s1, s2 IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.FastDifferenceWith(&s2)
	}
}

func BenchmarkMapDifferenceWith(b *testing.B) {
	var s1, s2 mapset.IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 mapset.IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.DifferenceWith(&s2)
	}
}

// --------------------------------------------------

func BenchmarkSymmetricDifference(b *testing.B) {
	var s1, s2 IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.SymmetricDifference(&s2)
	}

}

func BenchmarkFastSymmetricDifference(b *testing.B) {
	var s1, s2 IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.FastSymmetricDifference(&s2)
	}
}

func BenchmarkMapSymmetricDifference(b *testing.B) {
	var s1, s2 mapset.IntSet
	for i := 0; i < 50000; i++ {
		s1.Add(i)
		s2.Add(i + 25000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer() //copy中はタイマーを止める
		var s3 mapset.IntSet
		s3 = *s1.Copy()
		b.StartTimer()
		s3.SymmetricDifference(&s2)
	}
}

/*

                          ./+o+-       kurenaif@kagerou
                  yyyyy- -yyyyyy+      OS: Ubuntu 19.04 disco
               ://+//////-yyyyyyo      Kernel: x86_64 Linux 5.0.0-32-generic
           .++ .:/++++++/-.+sss/`      Uptime: 5h 40m
         .:++o:  /++++++++/:--:/-      Packages: 1928
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
               /osyyyyyyo++ooo+++/     RAM: 5550MiB / 15987MiB
                   ````` +oo+++o\:
                          `oo++.
goos: linux
goarch: amd64
pkg: go_training/ch11/ex07/intset
BenchmarkAdd-8                       	    2494	    480196 ns/op
BenchmarkFastAdd-8                   	    5733	    189746 ns/op
BenchmarkMapAdd-8                    	     134	   8885993 ns/op
BenchmarkUnionWith-8                 	  294072	      4066 ns/op
BenchmarkFastUnionWith-8             	  360075	      3557 ns/op
BenchmarkMapUnionWith-8              	     199	   5970196 ns/op
BenchmarkLen-8                       	 1000000	      1116 ns/op
BenchmarkMapLen-8                    	1000000000	         0.782 ns/op
BenchmarkRemove-8                    	    4606	    261371 ns/op
BenchmarkFastRemove-8                	    8064	    155935 ns/op
BenchmarkMapRemove-8                 	1000000000	         0.00421 ns/op
BenchmarkIntersectWith-8             	 1000000	      1216 ns/op
BenchmarkFastIntersectWith-8         	  980601	      1226 ns/op
BenchmarkMapIntersectWith-8          	     309	   3850890 ns/op
BenchmarkDifferenceWith-8            	  928645	      1224 ns/op
BenchmarkFastDifferenceWith-8        	 1000000	      1207 ns/op
BenchmarkMapDifferenceWith-8         	     325	   3860140 ns/op
BenchmarkSymmetricDifference-8       	  305380	      4112 ns/op
BenchmarkFastSymmetricDifference-8   	  331898	      3646 ns/op
BenchmarkMapSymmetricDifference-8    	      68	  16670782 ns/op
PASS
ok  	go_training/ch11/ex07/intset	134.689s
goos: linux
goarch: 386
pkg: go_training/ch11/ex07/intset
BenchmarkAdd-8                       	    2445	    504794 ns/op
BenchmarkFastAdd-8                   	    5019	    222562 ns/op
BenchmarkMapAdd-8                    	     100	  11149541 ns/op
BenchmarkUnionWith-8                 	  167602	      7202 ns/op
BenchmarkFastUnionWith-8             	  272862	      4434 ns/op
BenchmarkMapUnionWith-8              	     172	   6840884 ns/op
BenchmarkLen-8                       	   45679	     26738 ns/op
BenchmarkMapLen-8                    	1000000000	         0.813 ns/op
BenchmarkRemove-8                    	    2734	    455184 ns/op
BenchmarkFastRemove-8                	    6955	    158885 ns/op
BenchmarkMapRemove-8                 	1000000000	         0.00390 ns/op
BenchmarkIntersectWith-8             	  488060	      2509 ns/op
BenchmarkFastIntersectWith-8         	  473642	      2551 ns/op
BenchmarkMapIntersectWith-8          	     324	   3753621 ns/op
BenchmarkDifferenceWith-8            	  391075	      3175 ns/op
BenchmarkFastDifferenceWith-8        	  414775	      3035 ns/op
BenchmarkMapDifferenceWith-8         	     315	   3880881 ns/op
BenchmarkSymmetricDifference-8       	  167940	      7289 ns/op
BenchmarkFastSymmetricDifference-8   	  271934	      4401 ns/op
BenchmarkMapSymmetricDifference-8    	      62	  17548933 ns/op
PASS
ok  	go_training/ch11/ex07/intset	93.770s

*/
