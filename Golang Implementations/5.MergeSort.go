package main

import(
	"fmt"
)

type array struct{
	a []int
}


func main() {
	fmt.Println("--------Merge Sort--------")
	arr := array{[]int{3, 4, -5, 1, 6, 8, -2, -8}}
	fmt.Println("Before Sorting, arr =", arr)
	sortedArr := mergeSort(&arr)
	fmt.Println("After Sorting, arr =", *sortedArr)
}


//Merge Sort
func mergeSort(v *array) *array{
	if len(v.a) <= 1{
		return v
	}
	mid := (len(v.a)-1)/2
	leftHalf := array{v.a[:mid+1]}
	rightHalf := array{v.a[mid+1:]}
	left := mergeSort(&leftHalf)
	right := mergeSort(&rightHalf)
	merged := []int{}
	i, j := 0, 0
	for i < len(left.a) && j < len(right.a){
		if left.a[i] <= right.a[j]{
			merged = append(merged, left.a[i])
			i++
		} else {
			merged = append(merged, right.a[j])
			j++
		}
	}
	if i < len(left.a){
		merged = append(merged, left.a[i:]...)
	}
	if j < len(right.a){
		merged = append(merged, right.a[j:]...)
	}
	return &array{merged}
}