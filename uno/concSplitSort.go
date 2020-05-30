/*
 * concSplitSort.go
 * Author: LTrevino
 *
 * Concurrent Split Sort
 *
 * Compiled and tested using Go version go1.13.8 windows/amd64
 *
 * Please note that I like using semicolons -- valid in Go.
 */
package uno

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

const N_THREADS int = 4;



func ParseIntegers(intStrList string) (intSlice []int) {
	intStrings := strings.Split(intStrList, " ");
	intSlice = make([]int, 0, 20);
	for _, str := range intStrings {
		nn, err := strconv.Atoi(str);
		if (err!=nil) {
			fmt.Printf("Number '%s' could not be parsed -- ignored. \n")
		} else {
			intSlice = append(intSlice, nn);
		}
	}
	return intSlice;
}

func sortChunk(ch chan string, chunk []int, threadId int, fromIdx int, toIdx int) {
	fmt.Printf("Thread %d <-- Raw chunk    [%d:%d] : %v \n", threadId, fromIdx, toIdx, chunk);
	sort.Ints(chunk);
	fmt.Printf("Thread %d --> Sorted chunk [%d:%d] : %v \n", threadId, fromIdx, toIdx, chunk);
	ch <- "done";
}

func SplitSort(ch chan string, numList []int) {
	var chunkSize int = int(math.Ceil( float64(len(numList)) / float64(N_THREADS) ));
	// fmt.Printf("Chunk size: %d \n", chunkSize);
	var thCount = 0;
	for iPage := 0; iPage < N_THREADS; iPage++ {
		iFrom := iPage * chunkSize;
		if iFrom >= len(numList) {
			break;
		}
		thCount++;
		iTo := iFrom + chunkSize;
		if (iTo > len(numList)) {
			iTo = len(numList)
		}
		iChunk := numList[iFrom:iTo];
		go sortChunk(ch, iChunk, iPage, iFrom, iTo);
	}
	for idx:=0; idx < thCount; idx++ {
		<- ch;
	}
	// fmt.Printf("Merged list:      %d \n", numList);
	sort.Ints(numList);
	fmt.Printf("Sorted full list: %d \n", numList);
}

func ConcSplitSortDemo() {
	ch := make(chan string);
	fmt.Println("\n* Demo of SplitSort (concurrent split sort)");
	var numList = []int {20,19,18,17};
	//fmt.Println("Enter integer numbers: separated by space (e.g. \"7 6 5 4 3 2 1\")");
	//usrInput, err := uno.AcceptLn();
	//if (err!=nil) {fmt.Printf("Error reading numbers. %s\n", err)}
	//fmt.Println("usr input: ", usrInput)
	//numList := ParseIntegers(usrInput);
	fmt.Printf("Original list:     %v \n" , numList);
	SplitSort(ch, numList);	
}

