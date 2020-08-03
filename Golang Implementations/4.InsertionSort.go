package main

import(
	"fmt"
)

type array struct{
	a []int
}


func main() {
	fmt.Println("--------Insertion Sort Method 1--------")
	arr := array{[]int{3, 4, -5, 1, 6, 8, -2, -8}}
	fmt.Println("Before Sorting, arr =", arr)
	insertionSort1(&arr)
	fmt.Println("After Sorting, arr =", arr)
	fmt.Println("--------Insertion Sort Method 2--------")
	arr = array{[]int{3, 4, -5, 1, 6, 8, -2, -8}}
	fmt.Println("Before Sorting, arr =", arr)
	insertionSort1(&arr)
	fmt.Println("After Sorting, arr =", arr)
}




// Insertion Sort Method1 (more efficient)
func insertionSort1(v *array){
	for i:=1; i<len(v.a); i++{
		key := v.a[i]
		j := i-1
		for j>=0 && key<v.a[j] {
			v.a[j+1] = v.a[j]
			j--
		}
		v.a[j+1] = key
	}
}

// Insertion Sort Method2
func insertionSort2(v *array){
	for i:=0; i<len(v.a)-1; i++{
		idx := i+1
		for j:=i; j>=0; j--{
			if v.a[j] > v.a[idx]{
				v.a[j], v.a[idx] = v.a[idx], v.a[j]
				idx = j
			} else {
				break
			}
		}
	}
}
