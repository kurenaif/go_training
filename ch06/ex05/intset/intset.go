package intset

import (
	"bytes"
	"fmt"
	"math/bits"
)

const (
	BIT_SIZE = 32 << (^uint(0) >> 63)
)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/BIT_SIZE, uint(x%BIT_SIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/BIT_SIZE, uint(x%BIT_SIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) { // cnt fieldをもたせても良いが、どうせここでO(要素数)かかるし、毎回数える方針にする
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Stringは常に小さい順に出力されることが保証されている
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < BIT_SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", BIT_SIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() (cnt int) {
	for _, word := range s.words {
		cnt += bits.OnesCount(word)
	}
	return cnt
}

// ないものをremoveしても特に何もしない
func (s *IntSet) Remove(x int) {
	word, bit := x/BIT_SIZE, uint(x%BIT_SIZE)
	if word < len(s.words) {
		s.words[word] &= (^(1 << bit))
	}
}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

func (s *IntSet) Copy() *IntSet {
	var clone IntSet
	if s == nil {
		return nil
	}
	clone.words = make([]uint, len(s.words))
	copy(clone.words, s.words)
	return &clone
}

func (s *IntSet) AddAll(values ...int) {
	for _, value := range values {
		s.Add(value)
	}
}

// s & t
func (s *IntSet) IntersectWith(t *IntSet) {
	// cnt fieldをもたせても良いが、どうせここでO(要素数)かかるし、毎回数える方針にする
	mi := len(s.words)
	if mi > len(t.words) {
		mi = len(t.words)
	}

	// サイズが違った時の処理がめんどくさいので小さい方に合わせる
	s.words = s.words[:mi]
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// s - t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= (^tword)
		}
	}
}

// s XOR t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	// cnt fieldをもたせても良いが、どうせここでO(要素数)かかるし、毎回数える方針にする
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else { // 片方しか存在しないときはそのまま追加すればいい
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Elems() []uint {
	elems := []uint{}

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < BIT_SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, uint(BIT_SIZE*i+j))
			}
		}
	}
	return elems
}
