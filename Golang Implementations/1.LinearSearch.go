package main

import (
	"fmt"
)

type rlsArgs struct{
	targ int
	arr []int
	idx int
}


func main(){
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}


	found := bls(9, arr)
	fmt.Println("-----Better Linear Search-----")
	print(found)
	fmt.Println("-----Sentinel Linear Search-----")
	found = sls(9, arr)
	print(found)
	fmt.Println("-----Modified Sentinel Linear Search-----")
	found = msls(9, arr)
	print(found)
	fmt.Println("-----Recursive Linear Search-----")
	vals := rlsArgs{targ:9, arr: arr}
	found = rls(vals)
	print(found)
}

func print(found bool){
	if found == true{
		fmt.Println("Found target element")
	} else {
		fmt.Println("Not Found target element")
	}
}

// Better-Linear-Search
func bls(target int, arr []int) bool {
	for _, v := range arr{
		if v == target{
			return true
		}
	}
	return false
}

// Sentinel Linear Search
func sls(target int, arr []int) bool {
	temp := arr[len(arr)-1]
	arr[len(arr)-1] = target
	i := 0
	for arr[i] != target{
		i++
	}
	arr[len(arr)-1] = temp
	if i < len(arr)-1 || arr[len(arr)-1] == target{
		return true
	}
	return false
}


// Modified Sentinel Linear Saerch
func msls(target int, arr []int) bool {
	arr = append(arr, target)
	i := 0
	for arr[i] != target{
		i++
	}
	arr = arr[:len(arr)-1]
	if i == len(arr){
		return false
	}
	return true
	
}

// Recursive Linear Search
func rls(v rlsArgs) bool {
	if v.idx>=len(v.arr){
		return false
	} 
	if v.targ == v.arr[v.idx]{
		return true
	} 
	v.idx += 1
	return rls(v)
}
