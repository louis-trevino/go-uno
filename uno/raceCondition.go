/*
 * raceCondition.go
 * Author: LTrevino
 *
 * Compiled and tested using Go version go1.13.8 windows/amd64
 *
 * Please note that I like using semicolons -- valid in Go.
 */
package uno

import (
	"fmt"
)

var Num int;

func SetTo20() {
	Num = 20;
}

func SetTo50() {
	Num = 50;
}

func RaceConditionDemo() {
	fmt.Println("\n* Demo of raceCondition.go");
	num := 5; // the main thread is setting the value of num=5.
	fmt.Printf("Initial value of num in main thread is: %d. \n", num);
	// the main thread spawns a new worker thread (goroutine) that sets num=20, but does NOT wait for worker thread to finish.
	go SetTo20();
	// the main thread spawns a new worker thread (goroutine) that sets num=50, but does NOT wait for worker thread to finish.
	go SetTo50();
	// Main thread prints value of num, which will NOT necessarily be the expected value of 50.
	fmt.Printf("Expected value of num would be 50 (if threads were synchronized), but actual Value of num is: %d \n", num);	
}