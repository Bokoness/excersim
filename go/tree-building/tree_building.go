package main

import (
	"fmt"
)

type Record struct {
	ID     int
	Parent int
}
type Node struct {
	ID       int
	Children []*Node
}

func main() {
	arr := []Record{
		{ID: 3, Parent: 0},
		{ID: 2, Parent: 0},
		{ID: 1, Parent: 0},
		{ID: 0},
	}
	a, b := Build(arr)
	fmt.Println(a, b)
}

func Build(r []Record) (Node, error) {
	var p Node
	for _, v := range r {
		if v.ID == 0 {
			p.ID = v.ID
		}
	}
	return BuildHelper(r, p, 0)
}
func BuildHelper(r []Record, root Node, i int) (Node, error) {
	for i < len(r) {
		v := r[i]
		if v.Parent == root.ID {
			n := Node{
				ID: v.ID,
			}
			root.Children = append(root.Children, &n)
			BuildHelper(r, n, i+1)
		}
		i++
	}
	return root, nil
}

// []Record{
// 	{ID: 3, Parent: 0},
// 	{ID: 2, Parent: 0},
// 	{ID: 1, Parent: 0},
// 	{ID: 0},
// }
// expected: &Node{
// 	ID: 0,
// 	Children: []*Node{
// 		{
// 			ID: 1,
// 			Children: []*Node{
// 				{ID: 4},
// 				{ID: 5},
// 			},
// 		},
// 		{
// 			ID: 2,
// 			Children: []*Node{
// 				{ID: 3},
// 				{ID: 6},
// 			},
// 		},
// 	},
// },
