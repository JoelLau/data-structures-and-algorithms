package main

import "fmt"

const NUM_TO_POPULATE = 5

func main() {
	// slice initialisation (default is `nil`)
	// time: O(n) -> append 1 by 1
	var arr []int = make([]int, 5)
	fmt.Printf("arr aft init  : %+v\n", arr) // -> [0 0 0 0 0]

	// append items
	// time: O(n) -> take worst case
	//    - [best  case] write in memory location that's already allocated
	//    - [worst case] create new array with twice(?) as much space, copy items one at a time
	arr = append(arr, 6)
	fmt.Printf("arr aft append: %+v\n", arr) // -> [0 0 0 0 0 6]

	// iterate
	// time: O(n) -> go through each item
	for k, v := range arr {
		fmt.Printf("%+v:%+v\n", k, v) // -> 0:0, 1:0, 2:0, 3:0, 4:0, 5:6
	}

	// random access
	// time: O(1) -> easy to calculate memory address based on data type
	//    - e.g. index 2:  0x0000 + (2 x 8 bytes) = new address
	index := 2
	fmt.Printf("random access [%+v]: %+v\n", index, arr[index]) // [0]: 2
}
