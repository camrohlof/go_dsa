package main

import "fmt"

func main() {
	fmt.Printf("Hello World.")
	arr := []int{1, 2, 3, 4, 5, 6}
	pos := BinarySearch(arr, 4)
	if pos > 0 {
		fmt.Println(pos)
	}
	arr = []int{2, 1, 4, 3}
	BubbleSort(&arr)
	for _, val := range arr {
		fmt.Println(val)
	}

	ll := new(LinkedList)

	ll.append(5)
	ll.append(4)
	ll.append(3)
	ll.append(2)
	ll.append(1)
	ll.prepend(44)
	ll.print()
}
