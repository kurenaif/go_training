package intset

import (
	"bytes"
	"fmt"
	"math/bits"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() (cnt int) {
	for _, word := range s.words {
		cnt += bits.OnesCount64(word)
	}
	return cnt
}

// ないものをremoveしても特に何もしない
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &= (^(1 << bit))
	}
}

func (s *IntSet) Clear() {
	s.words = s.words[:0] // nil代入でも良い
}

func (s *IntSet) Copy() *IntSet {
	var clone IntSet
	if s == nil {
		return nil
	}
	clone.words = make([]uint64, len(s.words))
	copy(clone.words, s.words)
	return &clone
}
