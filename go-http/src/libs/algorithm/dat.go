package algorithm

import (
	"fmt"
)

/*
 Double Array Trie
*/

type Node struct {
	Depth    int
	Code     byte
	Children map[byte]*Node
}

func BuildDFA(arr [][]byte) *Node {
	var root = Node{
		Depth:    0,
		Code:     0,
		Children: make(map[byte]*Node),
	}

	for _, vector := range arr {
		var pre_node = &root

		for j, bs_code := range vector {
			var node = Node{
				Depth:    j + 1,
				Code:     bs_code,
				Children: make(map[byte]*Node),
			}

			fmt.Println(string(bs_code))

			pre_node.Children[bs_code] = &node
			pre_node = &node
		}
	}

	return &root
}

func BuildDAT(arr [][]byte) {
	fmt.Println(arr)
	var arr_len = len(arr)
	const num = 256

	var max_item_len = 0
	for _, item := range arr {
		if max_item_len <= len(item) {
			max_item_len = len(item)
		}
	}

	fmt.Println("max len", max_item_len)
	var base = make([]byte, max_item_len)
	var check = make([]byte, max_item_len)

	var depth_arr = make([][]byte, max_item_len)

	//TODO DELETE
	var depth_trie = Node{
		Depth:    0,
		Code:     0,
		Children: map[byte]*Node{},
	}

	fmt.Println(depth_trie)
	//

	for i := 0; i < max_item_len; i++ {
		var tmp_arr = make([]byte, arr_len)
		var tmp_i = 0
		for _, item := range arr {

			if len(item) > i {
				tmp_arr[tmp_i] = item[i]
			}
			tmp_i++
		}
		depth_arr[i] = tmp_arr

		fmt.Println(tmp_arr)
	}

	var dat_i = 0
	for i, vector := range depth_arr {
		for j, byte_code := range vector {
			if i == 0 {
				if j == 0 {
					base[dat_i] = byte_code
				}
				fmt.Println(byte_code)
				dat_i++
			} else {
			}
		}
	}

	fmt.Println(base, check)
	fmt.Println(max_item_len)
	fmt.Println(depth_arr)
}
