package sort

import (
	"fmt"
	"math/rand"
	"time"
)

// add changes for new commit
// initialization for randomizer
func init() {
	rand.Seed(time.Now().UnixNano())
}

// add additional changes for new commit version 2
// BUBBLE SORT
func bubbleSort(ar []int) {
	l := len(ar)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}

}

// add  changes for REMOTE
// SELECTION SORT
func selectionSort(ar []int) []int {
	l := len(ar)
	var min_index int
	var temp int
	for i := 0; i < l-1; i++ {
		min_index = i
		for j := i + 1; j < l; j++ {
			if ar[j] < ar[min_index] {
				min_index = j
			}
		}
		temp = ar[i]
		ar[i] = ar[min_index]
		ar[min_index] = temp
	}
	return ar
}

// INSERT SORT
func insertionSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	for i := 1; i < len(ar); i++ {
		for j := i; j > 0 && ar[j-1] > ar[j]; j-- {
			ar[j-1], ar[j] = ar[j], ar[j-1]
		}
	}
}

// MERGE SORT
func merge(ar []int) []int {
	if len(ar) <= 1 {
		return ar
	}
	l := len(ar)
	half := l / 2
	left := ar[:half]
	right := ar[half:]
	left, right = merge(left), merge(right)
	return mergeSort(left, right)
}

func mergeSort(a, b []int) []int {
	lenA, lenB := len(a), len(b)
	res := make([]int, 0, lenA+lenB)
	idxA, idxB := 0, 0
	for idxA+idxB < lenA+lenB {
		if idxB == lenB || (idxA < lenA && a[idxA] < b[idxB]) {
			res = append(res, a[idxA])
			idxA++
		} else {
			res = append(res, b[idxB])
			idxB++
		}
	}
	return res
}

// QUICK SORT
func quickSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	left, right := 0, len(ar)-1
	pivot_index := 0
	ar[pivot_index], ar[right] = ar[right], ar[pivot_index]
	for i := range ar {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}
	ar[left], ar[right] = ar[right], ar[left]

	quickSort(ar[:left])
	quickSort(ar[left+1:])
	return ar
}

func main() {

	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100
	}
	//bubbleSort(ar)
	fmt.Println(quickSort(ar))

}
