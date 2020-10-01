package main

import "fmt"

func bubbleSort(array []int) []int {
    for i := 0; i < len(array)-1; i++ {
        for j := 0; j < len(array)-i-1; j++ {
            if array[j] > array[j+1] {
                array[j], array[j+1] = array[j+1], array[j]
            }
        }
    }
    return array 
}

func main() {
    items := []int{8, 6 , 2, -9, 128, 48}
    res := bubbleSort(items)
    for _, val := range res {
        fmt.Println(val)
    }
}