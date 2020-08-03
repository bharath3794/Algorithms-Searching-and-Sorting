package main

import(
	"fmt"
)

type array struct{
	a []int
}


func main() {
	fmt.Println("--------Selection Sort--------")
	arr := array{[]int{3, 4, -5, 1, 6, 8, -2, -8}}
	fmt.Println("Before Sorting, arr =", arr)
	selectionSort(&arr)
	fmt.Println("After Sorting, arr =", arr)
}


// Selection Sort
func selectionSort(v *array){
	for i:=0; i<len(v.a)-1; i++{
		minIdx := i
		for j:=i+1; j<len(v.a); j++{
			if v.a[j] <= v.a[minIdx]{
				minIdx = j
			}
		}
		v.a[i], v.a[minIdx] = v.a[minIdx], v.a[i]
	}
}