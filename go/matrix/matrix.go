// package matrix

package main

import (
	"fmt"
	"strconv"
)

type Matrix struct {
	len  int
	vals []int
}

func (m Matrix) Rows() [][]int {
	return [][]int{{1, 2, 3}}
}

func (m Matrix) Cols() [][]int {
	return [][]int{{1, 2, 3}}
}

func New(s string) (*Matrix, error) {

	var m Matrix
	l := 0
	firstColComplete := false
	for i := range s {
		matlen := len(s)
		fmt.Println(matlen)
		bo := firstColComplete && (s[i] == 10 || i >= len(s)-1) && l+1 != m.len
		fmt.Println(bo)
		fmt.Println("first complete %b", firstColComplete)
		fmt.Println("at the end %b", s[i] == 10 || i >= len(s)-1)
		fmt.Println("over length %b", l+1 != m.len)
		c := s[i] //10 \n 32 " "
		str := string(s[i])
		fmt.Println(c, str)
		if firstColComplete && (s[i] == 10 || i >= len(s)-1) && l+1 != m.len {
			return nil, fmt.Errorf("Cols are not equal")
		}
		if s[i] == 32 {
			continue
		}
		l++
		if s[i] == 10 {
			if !firstColComplete {
				firstColComplete = true
				m.len = l // TODO: not true
			}
			l = 0
		} else {
			v, _ := strconv.Atoi(string(s[i]))
			m.vals = append(m.vals, v)
		}
	}
	// fmt.Println(string(num))

	var m1 Matrix = Matrix{}
	return &m1, nil
}

func (m Matrix) Set(row, column, value int) bool {
	return true
}

func main() {
	fmt.Println(New("1 2 \n 1 2 3"))
}
