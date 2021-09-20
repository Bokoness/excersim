// package secret
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Handshake(19))
}

//Handshake converts number to appropriate sequence of events for a secret handshake
func Handshake(n uint) []string {
	var set []string
	fmt.Printf("%b \n", n)
	b := fmt.Sprintf("%b", n)
	var z string
	for i := len(b) - 1; i >= 0; i-- {
		z += "0"
		if b[i] == "0" {
			continue
		}
		event(string(v)+zeros, &set)
	}
}

func event(s string, set *[]string) {
	switch s {
	case "1":
		set = append(set, "wink")
	case "10":
		set = append(set, "double blink")
	case "100":
		set = append(set, "close your eyes")
	case "1000":
		set = append(set, "jump")
	case "10000":
		for i, j := 0, len(set)-1; i < j; i, j = i+1, j-1 {
			set[i], set[j] = set[j], set[i]
		}
	}
}

func intToB(n int) int {
	s := ""
	rs := ""
	for n > 0 {
		s += strconv.Itoa(n % 2)
		n /= 2
	}
	for i := len(s) - 1; i >= 0; i-- {
		rs += string(s[i])
	}
	result, _ := strconv.ParseInt(rs, 0, 64)
	return int(result)
}

func reverse(s []string) []string {
	rs := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		rs = append(rs, s[i])
	}
	return rs
}
