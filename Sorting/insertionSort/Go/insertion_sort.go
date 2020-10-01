package main

import "fmt"

func insertionSort(array []int) []int {
    for i := 0; i < len(array); i++ {
        for j := 0; j < i; j++ {
            if array[j] > array[i] {
                array[j], array[i] = array[i], array[j]
            }
        }
    }
    return array 
}

func main() {
    items := []int{8, 6 , 2, -9, 128, 48}
    res := insertionSort(items)
    for _, val := range res {
        fmt.Println(val)
    }
}