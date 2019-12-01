#include <iostream>
#include <vector>
#include <cstdlib>

using namespace std;

const int MOD = 1000000007;

vector<int> memo;

int fibo(int a){
	if (memo.at(a) != 0) {
		return memo.at(a);
	}
	if (a == 0) {
		return 0;
	}
	if(a == 1){
		return 1;
	}
	memo.at(a) = (fibo(a-1) + fibo(a-2)) % MOD;
	return memo.at(a);
}

int main(int args, char* argv[]){
	int num = atoi(argv[1]);
	memo = vector<int>(num+1, 0);
	cout << fibo(num) << endl;
}
