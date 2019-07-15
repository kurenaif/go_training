#!/usr/local/bin/python3

"""
コマンドライン引数で渡された文字列をsha256して二進数のdiffを取る
"""

import hashlib
import sys

b1 = format(int(hashlib.sha256(sys.argv[1].encode()).hexdigest(), 16), "0255b")
b2 = format(int(hashlib.sha256(sys.argv[2].encode()).hexdigest(), 16), "0255b")
print(b1)
print(b2)

# b1とb2の先頭に0bはつくけど、差分を取る分には影響しないので無視

cnt = 0
for i in range(len(b1)):
    if b1[i] != b2[i]:
        cnt += 1

print(cnt)
