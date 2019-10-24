package intset

import "testing"

func isEqual(iset IntSet, mset map[int]bool) ([]int, []int) {
	onlyIset := []int{}
	onlyMset := []int{}
	for i := 0; i < 256*256; i++ {
		_, ok := mset[i]
		if iset.Has(i) != ok {
			if iset.Has(i) {
				onlyIset = append(onlyIset, i)
			}
			if ok {
				onlyMset = append(onlyIset, i)
			}
		}
	}
	return onlyIset, onlyMset
}

func TestAdd(t *testing.T) {

	var tests = []struct {
		addNums []int
	}{
		{[]int{}},
		{[]int{0, 1, 2}},
		{[]int{0, 0, 0, 0, 0}},
		{[]int{0, 64, 128}},
	}

	for _, test := range tests {
		iset := IntSet{}
		mset := map[int]bool{}

		for _, num := range test.addNums {
			iset.Add(num)
			mset[num] = true
		}

		onlyIset, onlyMset := isEqual(iset, mset)
		if len(onlyIset) != 0 || len(onlyMset) != 0 {
			t.Errorf("Add test failed: case %d, difference: exist only IntSet: %d, exist only mapset: %d ", test.addNums, onlyIset, onlyMset)
		}
	}
}

func TestUnionWith(t *testing.T) {

	var tests = []struct {
		addNumsLeft  []int
		addNumsRight []int
	}{
		{[]int{}, []int{}},
		{[]int{0, 1, 2}, []int{3, 4, 5}},
		{[]int{0, 1, 2}, []int{0, 1, 2}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{0, 64, 128}, []int{32, 128}},
	}

	for _, test := range tests {
		isetLeft := IntSet{}
		msetLeft := map[int]bool{}
		isetRight := IntSet{}
		msetRight := map[int]bool{}

		for _, num := range test.addNumsLeft {
			isetLeft.Add(num)
			msetLeft[num] = true
		}
		for _, num := range test.addNumsRight {
			isetRight.Add(num)
			msetRight[num] = true
		}

		isetLeft.UnionWith(&isetRight)
		for k, _ := range msetRight {
			msetLeft[k] = true
		}

		onlyIset, onlyMset := isEqual(isetRight, msetRight)
		if len(onlyIset) != 0 || len(onlyMset) != 0 {
			t.Errorf("Add test failed: case %d | %d , difference: exist only IntSet: %d, exist only mapset: %d ", test.addNumsLeft, test.addNumsRight, onlyIset, onlyMset)
		}
	}
}
