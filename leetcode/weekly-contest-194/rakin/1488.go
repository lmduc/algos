package main

import (
	"container/heap"
	"fmt"
)

func genNextRains(rains []int) []int {
	tracker := map[int]int{}
	nextRains := make([]int, len(rains))
	for i := len(rains) - 1; i >= 0; i-- {
		nxt, ok := tracker[rains[i]]
		if !ok {
			nxt = len(rains) + 10
		}

		nextRains[i] = nxt
		tracker[rains[i]] = i
	}
	return nextRains
}

type PQueue []int

func (p PQueue) Len() int {
	return len(p)
}

func (p PQueue) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p PQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQueue) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *PQueue) Pop() interface{} {
	old := *p
	l := len(old)
	v := old[l-1]
	*p = old[0 : l-1]
	return v
}

func avoidFlood(rains []int) []int {
	nextRains := genNextRains(rains)
	res := []int{}
	q := PQueue{}
	heap.Init(&q)
	fullLakes := map[int]int{}

	for i, rain := range rains {
		if rain == 0 {
			needToEmpty := 1
			if len(q) > 0 {
				ind := heap.Pop(&q).(int)
				fmt.Println("ind", ind)
				if ind > len(rains) {
					needToEmpty = 1
				} else {
					needToEmpty = rains[ind]
				}
			}

			res = append(res, needToEmpty)
			fullLakes[needToEmpty] = 0
			continue
		}
		if fullLakes[rain] > 0 {
			return []int{}
		}
		fullLakes[rain] = 1
		heap.Push(&q, nextRains[i])
		fmt.Println("test", nextRains[i])
		res = append(res, -1)
	}

	return res
}
