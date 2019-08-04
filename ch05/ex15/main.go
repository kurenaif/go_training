package main

import "fmt"

func min(vals ...int) (res *int) {
	for _, val := range vals {
		if res == nil || *res > val {
			if res == nil {
				res = new(int)
			}
			*res = val
		}
	}
	return res
}

func max(vals ...int) (res *int) {
	for _, val := range vals {
		if res == nil || *res < val {
			if res == nil {
				res = new(int)
			}
			*res = val
		}
	}
	return res
}

// 少なくとも一つの引数が必要
// 引数が0の場合、errorを返す
func min2(vals ...int) (int, error) {
	if len(vals) <= 0 {
		return 0, fmt.Errorf("one argument is required at least. (given %v)", vals)
	}

	// ここまで来る処理は必ず[0]は存在する
	res := vals[0]

	for _, val := range vals {
		if res > val {
			res = val
		}
	}
	return res, nil
}

// 少なくとも一つの引数が必要
// 引数が0の場合、errorを返す
func max2(vals ...int) (int, error) {
	if len(vals) <= 0 {
		return 0, fmt.Errorf("one argument is required at least. (given %v)", vals)
	}

	// ここまで来る処理は必ず[0]は存在する
	res := vals[0]

	for _, val := range vals {
		if res < val {
			res = val
		}
	}
	return res, nil
}

func main() {
	fmt.Println(*min(1, 3, 5, 7)) //  "1"
	fmt.Println(*max(1, 3, 5, 7)) //  "7"
	fmt.Println(min2(1, 3, 5, 7)) //  "7"
	fmt.Println(max2(1, 3, 5, 7)) //  "7"
}
