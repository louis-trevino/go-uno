/*
 * diners.go
 * Author: Louis Trevino
 * Copyright(C) Torino Consulting, 2020.
 *
 * Dining Philosophers with a twist: a Host issues a Dining Permit before a diner can eat.
 *   There are 5 dining philosophers, each of which should eat 3 times.
 *
 * Compiled and tested using Go version go1.13.8 windows/amd64
 *
 * Please note that I like using semicolons -- valid in Go.
 */
package uno

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup;

// structure for chopstick
type Chop struct{
	mx sync.Mutex;
}

type Diner struct {
	Id int;
	LeftChop  *Chop;
	RightChop *Chop;
}

type Host struct {
	permitsTray chan string;
}

func getOrdinal(num int) string {
	switch num {
	case 1:
		return "1st";
	case 2:
		return "2nd";
	case 3:
		return "3rd";
	default:
		return fmt.Sprintf("%dth", num);
	}
}

func (host *Host) IssueInitialDiningPermits() {
	host.permitsTray = make(chan string, 2);
	const PERMIT_A string = "A";
	const PERMIT_B string = "B";
	fmt.Printf("Host issuing dining permit %s. \n", PERMIT_A);
	host.permitsTray <- PERMIT_A;
	fmt.Printf("Host issuing dining permit %s. \n", PERMIT_B);
	host.permitsTray <- PERMIT_B;
	wg.Done();
}

func (host *Host) waitForDiningPermit() string {
	var permit string = <- host.permitsTray;
	return permit;
}

func (host *Host) returnDiningPermit(permit string) {
	host.permitsTray <- permit;
}

func (diner Diner) Eat(host *Host) {
	for idx:=0; idx < 3; idx++ {
		// wait for dining permit from host
		permit := host.waitForDiningPermit();
		diner.LeftChop.mx.Lock();    // get left chopstick
		diner.RightChop.mx.Lock();   // get right chopstick
		fmt.Printf("Philosopher Diner %d starting to eat (%s time) using dining permit %s.\n", diner.Id, getOrdinal(idx + 1), permit);
		time.Sleep(time.Millisecond * 200);
		diner.RightChop.mx.Unlock()  // return right chopstick
		diner.LeftChop.mx.Unlock();  // return right chopstick
		fmt.Printf("Philosopher Diner %d finishing eating (%s time) using dining permit %s. \n", diner.Id, getOrdinal(idx + 1), permit);
		// return dining permit to host
		host.returnDiningPermit(permit);
		time.Sleep(time.Millisecond * 100);
	}
	wg.Done()
}


func DinersDemo() {
	fmt.Println("\n* Demo of diners.go (concurrent Diner Philosophers++)");
	const NUM_DINERS = 5;
	chops := make([]*Chop, NUM_DINERS)
	diners := make([]*Diner, NUM_DINERS)
	// host issues initial two "dining permits", which will be consumed by diners
	host := &Host{};
	wg.Add(1);
	go host.IssueInitialDiningPermits();
	wg.Wait();
	// populate chopsticks list
	for i := 0; i < NUM_DINERS; i++ {
		chops[i] = new(Chop)
	}
	// populate diners list
	for i := 0; i < NUM_DINERS; i++ {
		diners[i] = &Diner{Id: i+1, LeftChop: chops[i], RightChop: chops[(i+1) % NUM_DINERS]}
	}
	// spawn new diner threads and call Eat() method
	for i := 0; i < NUM_DINERS; i++ {
		wg.Add(1)
		go diners[i].Eat(host);
	}

	wg.Wait() // wait for worker threads to finish
}
