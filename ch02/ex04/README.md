# 実行結果

* loopで回したものと、bitshiftで64回ループを回したものでは、ループを回したもののほうが遅くなった。
* ループは8倍回っているはずだが、8倍遅くはなっていない。
* objdumpした結果を比較した結果、bitshiftの方は処理が単純なため行数が少なくなっていることがわかった

```
$ sh test.bash 
+ go test -bench=.
goos: linux
goarch: amd64
pkg: go_training/ch02/ex04
BenchmarkPopCount11-8               	2000000000	         0.27 ns/op
BenchmarkPopCount00-8               	2000000000	         0.27 ns/op
BenchmarkPopCount00rand-8           	100000000	        19.7 ns/op
BenchmarkPopCountRand-8             	100000000	        19.3 ns/op
BenchmarkPopCountLoop11-8           	100000000	        17.6 ns/op
BenchmarkPopCountLoop00-8           	100000000	        17.7 ns/op
BenchmarkPopCountLoop00rand-8       	50000000	        29.9 ns/op
BenchmarkPopCountLoopRand-8         	50000000	        31.4 ns/op
BenchmarkPopCountBitShift00-8       	30000000	        38.3 ns/op
BenchmarkPopCountBitShift11-8       	30000000	        39.5 ns/op
BenchmarkPopCountBitShift00rand-8   	30000000	        53.9 ns/op
BenchmarkPopCountBitShiftRand-8     	30000000	        55.0 ns/op
PASS
ok  	go_training/ch02/ex04	17.578s
```

# なぜ早くなったのか？

`go build main.go`して出てきた実行ファイルを `objdump -M intel -D` して比較

* ループを使った実装では `0000000000483280 <go_training/ch02/ex03/popcountloop.PopCount>:` が生成されており、関数をcallしてstackに積んで… みたいな処理が挟まっている上、cmp命令等も発行されているため遅くなっている
* ループを使わない実装ではまず関数そのものが定義されておらず、インライン展開されている。movやshrコマンドなど、かなり計算負荷が安い演算を使用しているためかなり早いと考えられる

## ループを使わない実装 

```
  4834f9:	48 8d 15 a0 2d 0c 00 	lea    rdx,[rip+0xc2da0]        # 5462a0 <go_training/ch02/ex03/popcount.pc>
  483500:	0f b6 0c 11          	movzx  ecx,BYTE PTR [rcx+rdx*1]
  483504:	48 89 c3             	mov    rbx,rax
  483507:	48 c1 e8 08          	shr    rax,0x8
  48350b:	0f b6 c0             	movzx  eax,al
  48350e:	0f b6 04 10          	movzx  eax,BYTE PTR [rax+rdx*1]
  483512:	48 89 de             	mov    rsi,rbx
  483515:	48 c1 eb 10          	shr    rbx,0x10
  483519:	0f b6 db             	movzx  ebx,bl
  48351c:	0f b6 1c 13          	movzx  ebx,BYTE PTR [rbx+rdx*1]
  483520:	48 89 f7             	mov    rdi,rsi
  483523:	48 c1 ee 18          	shr    rsi,0x18
  483527:	40 0f b6 f6          	movzx  esi,sil
  48352b:	0f b6 34 16          	movzx  esi,BYTE PTR [rsi+rdx*1]
  48352f:	49 89 f8             	mov    r8,rdi
  483532:	48 c1 ef 20          	shr    rdi,0x20
  483536:	40 0f b6 ff          	movzx  edi,dil
  48353a:	0f b6 3c 17          	movzx  edi,BYTE PTR [rdi+rdx*1]
  48353e:	4d 89 c1             	mov    r9,r8
  483541:	49 c1 e8 28          	shr    r8,0x28
  483545:	45 0f b6 c0          	movzx  r8d,r8b
  483549:	45 0f b6 04 10       	movzx  r8d,BYTE PTR [r8+rdx*1]
  48354e:	4d 89 ca             	mov    r10,r9
  483551:	49 c1 e9 30          	shr    r9,0x30
  483555:	45 0f b6 c9          	movzx  r9d,r9b
  483559:	45 0f b6 0c 11       	movzx  r9d,BYTE PTR [r9+rdx*1]
  48355e:	49 c1 ea 38          	shr    r10,0x38
  483562:	42 0f b6 14 12       	movzx  edx,BYTE PTR [rdx+r10*1]
  483567:	48 01 c8             	add    rax,rcx
  48356a:	48 01 d8             	add    rax,rbx
  48356d:	48 01 f0             	add    rax,rsi
  483570:	48 01 f8             	add    rax,rdi
  483573:	4c 01 c0             	add    rax,r8
  483576:	4c 01 c8             	add    rax,r9
  483579:	48 01 d0             	add    rax,rdx
  48357c:	48 89 44 24 50       	mov    QWORD PTR [rsp+0x50],rax
```

## ループを使った実装

```
0000000000483280 <go_training/ch02/ex03/popcountloop.PopCount>:
  483280:	48 8b 44 24 08       	mov    rax,QWORD PTR [rsp+0x8]
  483285:	31 c9                	xor    ecx,ecx
  483287:	31 d2                	xor    edx,edx
  483289:	eb 30                	jmp    4832bb <go_training/ch02/ex03/popcountloop.PopCount+0x3b>
  48328b:	48 8d 59 01          	lea    rbx,[rcx+0x1]
  48328f:	48 c1 e1 03          	shl    rcx,0x3
  483293:	48 89 c6             	mov    rsi,rax
  483296:	48 d3 e8             	shr    rax,cl
  483299:	48 83 f9 40          	cmp    rcx,0x40
  48329d:	48 19 ff             	sbb    rdi,rdi
  4832a0:	48 21 f8             	and    rax,rdi
  4832a3:	0f b6 f8             	movzx  edi,al
  4832a6:	4c 8d 05 f3 30 0c 00 	lea    r8,[rip+0xc30f3]        # 5463a0 <go_training/ch02/ex03/popcountloop.pc>
  4832ad:	41 0f b6 3c 38       	movzx  edi,BYTE PTR [r8+rdi*1]
  4832b2:	48 01 fa             	add    rdx,rdi
  4832b5:	48 89 f0             	mov    rax,rsi
  4832b8:	48 89 d9             	mov    rcx,rbx
  4832bb:	48 83 f9 08          	cmp    rcx,0x8
  4832bf:	72 ca                	jb     48328b <go_training/ch02/ex03/popcountloop.PopCount+0xb>
  4832c1:	48 89 54 24 10       	mov    QWORD PTR [rsp+0x10],rdx
  4832c6:	c3                   	ret    
  4832c7:	cc                   	int3   
  4832c8:	cc                   	int3   
  4832c9:	cc                   	int3   
  4832ca:	cc                   	int3   
  4832cb:	cc                   	int3   
  4832cc:	cc                   	int3   
  4832cd:	cc                   	int3   
  4832ce:	cc                   	int3   
  4832cf:	cc                   	int3   
```

## bitshiftを使った実装

```
0000000000483330 <go_training/ch02/ex04/popcountbitshift.PopCount>:
  483330:	48 8b 44 24 08       	mov    rax,QWORD PTR [rsp+0x8]
  483335:	31 c9                	xor    ecx,ecx
  483337:	31 d2                	xor    edx,edx
  483339:	eb 13                	jmp    48334e <go_training/ch02/ex04/popcountbitshift.PopCount+0x1e>
  48333b:	48 ff c1             	inc    rcx
  48333e:	48 89 c3             	mov    rbx,rax
  483341:	48 83 e0 01          	and    rax,0x1
  483345:	48 01 c2             	add    rdx,rax
  483348:	48 d1 eb             	shr    rbx,1
  48334b:	48 89 d8             	mov    rax,rbx
  48334e:	48 83 f9 40          	cmp    rcx,0x40
  483352:	72 e7                	jb     48333b <go_training/ch02/ex04/popcountbitshift.PopCount+0xb>
  483354:	48 89 54 24 10       	mov    QWORD PTR [rsp+0x10],rdx
  483359:	c3                   	ret    
```