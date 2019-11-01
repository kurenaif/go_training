package mapset

import (
	"bytes"
	"fmt"
	"sort"
)

type IntSet struct {
	set map[uint]bool
}

func (s *IntSet) initialize() {
	if s.set == nil {
		s.set = make(map[uint]bool)
	}
}

func (s *IntSet) Has(x int) bool {
	s.initialize()
	_, ok := s.set[uint(x)]
	return ok
}

func (s *IntSet) Add(x int) {
	s.initialize()
	s.set[uint(x)] = true
}

func (s *IntSet) UnionWith(t *IntSet) { // cnt fieldをもたせても良いが、どうせここでO(要素数)かかるし、毎回数える方針にする
	s.initialize()
	for k, _ := range t.set {
		s.set[k] = true
	}
}

// Stringは常に小さい順に出力されることが保証されている
func (s *IntSet) String() string {
	s.initialize()
	var buf bytes.Buffer

	var res []uint
	for k, _ := range s.set {
		res = append(res, k)
	}
	sort.Slice(res, func(i int, j int) bool { return res[i] < res[j] })

	buf.WriteByte('{')
	for i := 0; i < len(res); i++ {
		fmt.Fprintf(&buf, "%d", res[i])
		if i != len(res)-1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() (cnt int) {
	s.initialize()
	return len(s.set)
}

// ないものをremoveしても特に何もしない
func (s *IntSet) Remove(x int) {
	s.initialize()
	delete(s.set, uint(x))
}

func (s *IntSet) Clear() {
	s.initialize()
	s.set = make(map[uint]bool)
}

func (s *IntSet) Copy() *IntSet {
	s.initialize()
	var clone IntSet
	if s == nil {
		return nil
	}
	for k, _ := range s.set {
		clone.Add(int(k))
	}
	return &clone
}

func (s *IntSet) AddAll(values ...int) {
	s.initialize()
	for _, value := range values {
		s.Add(value)
	}
}

// s & t
func (s *IntSet) IntersectWith(t *IntSet) {
	s.initialize()
	for k, _ := range s.set {
		if _, ok := t.set[k]; !ok {
			delete(s.set, k)
		}
	}
}

// s - t
func (s *IntSet) DifferenceWith(t *IntSet) {
	s.initialize()
	for k, _ := range s.set {
		if _, ok := t.set[k]; ok {
			delete(s.set, k)
		}
	}
}

// s XOR t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	s.initialize()
	memo := map[uint]bool{}
	for k, _ := range t.set {
		if _, ok := s.set[k]; !ok { // tのみ存在
			s.set[k] = true
			memo[k] = true
		}
	}
	for k, _ := range s.set {
		if _, ok := t.set[k]; ok { // s も tも存在
			if !memo[k] {
				delete(s.set, k)
			}
		}
	}
}

func (s *IntSet) Elems() []uint {
	s.initialize()
	elems := []uint{}

	for k, _ := range s.set {
		elems = append(elems, k)
	}

	sort.Slice(elems, func(i int, j int) bool { return elems[i] < elems[j] })

	return elems
}
