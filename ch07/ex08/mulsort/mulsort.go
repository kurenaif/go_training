package mulsort

import (
	"log"
	"sort"
)

type MultipleSortIntarface struct {
	sortInterface []sort.Interface
	seek          int
}

func (s *MultipleSortIntarface) Add(sorter sort.Interface) {
	s.sortInterface = append(s.sortInterface, sorter)
}

func (s *MultipleSortIntarface) Next() *sort.Interface {
	if s.seek >= len(s.sortInterface) {
		return nil
	}
	res := s.sortInterface[s.seek]
	s.seek++
	return &res
}

func (s MultipleSortIntarface) Len() int {
	if s.sortInterface == nil || len(s.sortInterface) == 0 {
		log.Printf("sort key is not setted. (may be sort is not enabled)")
		return 0
	}
	return s.sortInterface[0].Len()
}

func (s MultipleSortIntarface) Less(i, j int) bool {
	for _, sorter := range s.sortInterface {
		if sorter.Less(i, j) != sorter.Less(j, i) { //入れ替えると結果が変わる=>iとjは等しくない
			return sorter.Less(i, j)
		}
	}
	return false
}

func (s MultipleSortIntarface) Swap(i, j int) {
	if s.sortInterface == nil || len(s.sortInterface) == 0 {
		log.Printf("sort key is not setted. (may be sort is not enabled)")
		return
	}
	s.sortInterface[0].Swap(i, j)
}
