# 出力結果

## 所感

* 何度実行しても1回目より2回目のほうがfetchしている時間は短い
* 一度実行終了すると、再び大きくなるのでキャッシュはアプリが落ちるまで有効

```
+ cat out1.txt
0.06s	  13859	http://www.in2white.com/
0.09s	 106382	https://www.nicovideo.jp/
0.11s	  38427	https://www.pixiv.net/
0.13s	   6483	https://earth.google.com/web/
0.14s	  28574	http://kurenaif.html.xdomain.jp/hanahudagd2/
0.22s	 301526	https://twitter.com/
1.19s	2278556	http://ksnctf.sweetduet.info/log
1.19s elapced
+ cat out2.txt
0.04s	  38427	https://www.pixiv.net/
0.04s	  13859	http://www.in2white.com/
0.05s	   6483	https://earth.google.com/web/
0.06s	 106612	https://www.nicovideo.jp/
0.06s	  28574	http://kurenaif.html.xdomain.jp/hanahudagd2/
0.14s	 301526	https://twitter.com/
0.93s	2278556	http://ksnctf.sweetduet.info/log
0.93s elapced

```

```
+ cat out1.txt
0.07s	  28574	http://kurenaif.html.xdomain.jp/hanahudagd2/
0.08s	  13859	http://www.in2white.com/
0.15s	 106555	https://www.nicovideo.jp/
0.15s	  38427	https://www.pixiv.net/
0.16s	   6483	https://earth.google.com/web/
0.27s	 301526	https://twitter.com/
1.26s	2278556	http://ksnctf.sweetduet.info/log
1.26s elapced
+ cat out2.txt
0.03s	  13859	http://www.in2white.com/
0.04s	  38427	https://www.pixiv.net/
0.06s	   6483	https://earth.google.com/web/
0.08s	  28574	http://kurenaif.html.xdomain.jp/hanahudagd2/
0.09s	 106789	https://www.nicovideo.jp/
0.15s	 301526	https://twitter.com/
1.09s	2278556	http://ksnctf.sweetduet.info/log
1.09s elapced
```
