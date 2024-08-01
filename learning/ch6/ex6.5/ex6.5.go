// 根据机器型号， 来进行操作, 直接把字长设置成wordSize = 32 << (^uint(0) >> 32 & 1)
package main

import (
	"bytes"
	"fmt"
)

const (
	wordSize = 32 << (^uint(0) >> 32 & 1)
)

type IntSet struct {
	words []uint64 
}

func  (s *IntSet) Has(x int) bool {
	word, bit := x / wordSize, uint(x % wordSize)
	return word < len(s.words) && s.words[word] & (1 << bit) != 0 
}

// 按位或
func (s *IntSet) Add(x int) {
	word, bit := x / wordSize, uint(x % wordSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// 合并
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer 
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word & (1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordSize * i + j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main(){
	var x, y IntSet 
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42144}")
}