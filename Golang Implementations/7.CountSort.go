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
	fmt.Println("--------Count Sort Method 1--------")
	arr := array{[]int{4,1,5,0,1,6,5,1,5,3}}
	fmt.Println("Before Sorting, arr =", arr)
	sortedArr := countSort1(&arr, 7)
	fmt.Println("After Sorting, arr =", *sortedArr)
	arr = array{[]int{8, 7, 6, 5, 4, 3, 2, 1000}}
	fmt.Println("Before Sorting, arr =", arr)
	sortedArr = countSort1(&arr, 1001)
	fmt.Println("After Sorting, arr =", *sortedArr)
	fmt.Println("--------Count Sort Method 2--------")
	arr = array{[]int{4,1,5,0,1,6,5,1,5,3}}
	fmt.Println("Before Sorting, arr =", arr)
	sortedArr = countSort2(&arr, 7)
	fmt.Println("After Sorting, arr =", *sortedArr)
	arr = array{[]int{8, 7, 6, 5, 4, 3, 2, 1000}}
	fmt.Println("Before Sorting, arr =", arr)
	sortedArr = countSort2(&arr, 1001)
	fmt.Println("After Sorting, arr =", *sortedArr)
}



// Count Sort Method 1
func countSort1(v *array, m int) *array{
	equal := countSort1Equal(v, m)
	sortedArr := []int{}
	for i, val := range equal.a{
		for j:=0; j<val; j++{
			sortedArr = append(sortedArr, i)
		}
	}
	return &array{sortedArr}
}

// Count Sort Method 1 Dependency
func countSort1Equal(v *array, m int) *array{
	equal := make([]int, m)
	for _, key := range v.a{
		equal[key]++
	}
	return &array{equal}
}

// Count Sort Method 2
func countSort2(v *array, m int) *array{
	less := countSort2Less(v, m)
	sortedArr := make([]int, len(v.a))
	for _, key := range v.a{
		sortedArr[less.a[key]] = key
		less.a[key]++
	}
	return &array{sortedArr}
}


// Count Sort Method 2 Dependency
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
