package main

import(
	"fmt"
)

type array struct{
	a []int
}

type array1 struct{
	s []string
}

func main() {
	fmt.Println("--------Radix Sort using countSort2Less() function--------")
	arr1 := array1{[]string{"F6", "E5", "R6", "X6", "X2", "T5", "F2", "T3"}}
	fmt.Println("Before Sorting, arr =", arr1)
	radixSort(&arr1)
	fmt.Println("After Sorting, arr =", arr1.s)
	arr1 = array1{[]string{"XI7FS6", "PL4ZQ2", "JI8FR9", "XL8FQ6", "PY2ZR5", "KV7WS9", "JL2ZV3", "KI4WR2"}}
	fmt.Println("Before Sorting, arr =", arr1)
	radixSort(&arr1)
	fmt.Println("After Sorting, arr =", arr1.s)
}



// Radix Sort
// This function uses Count Sort Method 2 Dependency i.e. countSort2Less() function
func radixSort(v *array1){
	temp := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ref := map[string]int{}
	for i,val := range temp{
		ref[string(val)] = i
	}
	for i:=len(v.s[0])-1; i>=0; i--{
		temp := []int{}
		for _, val := range v.s{
			temp = append(temp, ref[string(val[i])])
		}
		less := countSort2Less(&array{temp}, 36)
		sortedIdx := make([]int, len(v.s))
		for i, key := range temp{
			sortedIdx[less.a[key]] = i
			less.a[key]++
		}
		sortedArr := make([]string, len(v.s))
		for i, idx := range sortedIdx{
			sortedArr[i] = v.s[idx]
		}
		v.s = sortedArr
	}
}


// radixSort Dependency
func countSort2Less(v *array, m int) *array{
	equal := make([]int, m)
	for _, val := range v.a{
		equal[val]++
	}
	less := []int{0}
	for i:=0; i<len(equal)-1; i++{
		less = append(less, less[i]+equal[i])
	}
	return &array{less}
}
