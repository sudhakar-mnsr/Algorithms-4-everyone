// This program takes the number of integers to be sorted as first input
// Recieves the numbers from user and sorts using merge sort

package main

// importing fmt and bytes package
import (
	"fmt"
	"math/rand"
	"time"
)

// MergeSorter algorithm
func MergeSorter(array []int) []int {

	if len(array) < 2 {
		return array
	}
	var middle int
	middle = (len(array)) / 2
	return JoinArrays(MergeSorter(array[:middle]), MergeSorter(array[middle:]))
}

// Join Arrays method
func JoinArrays(leftArr []int, rightArr []int) []int {

	var num int
	var i int
	var j int
	num, i, j = len(leftArr)+len(rightArr), 0, 0
	var array []int
	array = make([]int, num, num)

	var k int
	for k = 0; k < num; k++ {
		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
			array[k] = rightArr[j]
			j++
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
			array[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] {
			array[k] = leftArr[i]
			i++
		} else {
			array[k] = rightArr[j]
			j++
		}
	}
	return array
}

// main method
func main() {

        var num int
        
        fmt.Print("Enter Number of Elements: ")
        fmt.Scan(&num)
        
        var array = make([]int, num)

        var i int
        for i = 0; i < num; i++ {
                fmt.Print("array[", i, "]: ")
                fmt.Scan(&array[i])
        }       
        

	fmt.Println("\n Before Sorting \n\n", array)
	fmt.Println("\n-After Sorting\n\n", MergeSorter(array), "\n")
}
