#include <stdio.h>

int f(int a, int b )  {
	int memo[100] = {};
	int index = 0;
	for (int i = a; i < b; i++) {
		memo[index] = i;
		index++;
	}

	int res = 0;
	for (; index > 0; index--) {
		res += memo[index-1];
	}
	return res;
}

int main() {
	printf("%d", f(1,2));
}
