package main

import (
	"fmt"
)

type rlsArgs struct{
	targ int
	arr []int
	idx int
}

type binArgs struct{
	rlsArgs
	low int
	high int
}

func main(){
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	val1 := binArgs{rlsArgs: rlsArgs{targ:9, arr: arr}, high: len(arr)-1}
	fmt.Println("-----Binary Search without recursion-----")
	found := binSearch(val1)
	print(found)
	fmt.Println()
	fmt.Println("-----Binary Search using Recursion-----")
	found = recBinSearch(val1)
	print(found)
	fmt.Println()
	fmt.Println("-----Binary Search along with Index to insert element if not found-----")
	arr = []int{3, 5, 5, 7, 7, 9}
	val2 := binArgs{rlsArgs: rlsArgs{targ:10, arr: arr}, high: len(arr)-1}
	found, index := binSearchIdx(val2)
	print(found)
	if index >= 0{
		fmt.Printf("Inserting element %v at index %v\n", val2.targ, index)
		val2.arr = append(val2.arr, 0)
		copy(val2.arr[index:], val2.arr[index-1:])
		val2.arr[index] = val2.targ
		fmt.Printf("After inserting, arr = %v\n", val2.arr)
	}
}



///Iterative Binary Search
func binSearch(v binArgs) bool {
	for v.low <= v.high{
		mid := (v.low+v.high)/2
		if v.targ == v.arr[mid] {
			return true
		} else if v.targ < v.arr[mid] {
			v.high = mid-1
		} else { // v.targ > v.arr[mid]
			v.low = mid+1
		}
	}
	return false
}

//Iterative Binary Search along with index to insert the element if not found
func binSearchIdx(v binArgs) (bool, int) {
	for v.low <= v.high{
		mid := (v.low+v.high)/2
		if v.targ == v.arr[mid] {
			return true, -1
		} else if v.targ < v.arr[mid] {
			v.high = mid-1
		} else { // v.targ > v.arr[mid]
			v.low = mid+1
		}
	}
	return false, v.low
}



// Binary Search using Recursion
func recBinSearch(v binArgs) bool{
	if v.low > v.high{
		return false
	}
	mid := (v.low+v.high)/2
	if v.targ == v.arr[mid] {
		return true
	} else if v.targ < v.arr[mid] {
		v.high = mid-1
	} else { // v.targ > v.arr[mid]
		v.low = mid+1
	}
	return recBinSearch(v)
}
