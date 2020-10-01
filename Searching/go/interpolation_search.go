package main
import "fmt"
 
func interpolationSearch(target int, array []int) int {
 
    min, max := array[0], array[len(array)-1]
 
    low, high := 0, len(array)-1
 
    for {
        if target < min {
            return low
        }
 
        if target > max {
            return high + 1
        }
 
        var guess int
        if high == low {
            guess = high
        } else {
            size := high - low
            offset := int(float64(size-1) * (float64(target-min) / float64(max-min)))
            guess = low + offset
        }
 
        if array[guess] == target {
            for guess > 0 && array[guess-1] == target {
                guess--
            }
            return guess
        }
 
        if array[guess] > target {
            high = guess - 1
            max = array[high]
        } else {
            low = guess + 1
            min = array[low]
        }
    }
}

func main(){
    items := []int{1,2, 4, 8, 16, 32, 64}
    fmt.Println(interpolationSearch(64, items))
}