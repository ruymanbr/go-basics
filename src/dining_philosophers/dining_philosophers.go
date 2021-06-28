/*
 * File:    dining_philosophers.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	28-6-21
 * Course:  	Concurrency in Go (University of California IRVINE)
 * Assignment:  Module 4 - Activity dining_philosophers.go
 *
 * Summary of File:
 *
 *  The program implements the dining philosopher’s problem with the following constraints/modifications.
 *	
 *	1. There are 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of 
 *	philosophers.
 *
 *	2. Each philosopher should eat only 3 times (not in an infinite loop)
 *
 *	3. The philosophers pick up the chopsticks in any order, not lowest-numbered first
 *	
 *	4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
 *	
 *	5. The host allows no more than 2 philosophers to eat concurrently.
 *	
 *	6. Each philosopher is numbered, 1 through 5.
 *	
 *	7. When a philosopher starts eating (after it has obtained necessary locks) it prints: 
 *	   	
 *		“starting to eat <number>” 
 *
 *		(on a line by itself, where <number> is the number of the philosopher)
 *	
 *	8. When a philosopher finishes eating (before it has released its locks) it prints:
 *
 *		“finishing eating <number>” 
 *
 *		(on a line by itself, where <number> is the number of the philosopher)
 *
 */

package main

import (
	"fmt"
	"sync"
	"math"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	PNum float64
	EatTimes int
}

type Host struct {
	p1, p2 *Philo
}

// eat prints 2 lines of strings and reduces times allowed to eat from a given Philo type
func (p Philo) eat(notEating chan Philo, wg *sync.WaitGroup) {
	
	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Println("starting to eat ", p.PNum)	
	
	if (p.EatTimes > 0) {
		p.EatTimes -= 1		
	}

	if (p.EatTimes != 0) {
		notEating <- p
	}
	fmt.Println("finishing eating ", p.PNum)

	p.leftCS.Unlock()
	p.rightCS.Unlock()

}

// allowToEat calls method eat() from 2 given Philo types
func (h Host) allowToEat(notEating chan Philo, p1 Philo, p2 Philo, wg *sync.WaitGroup) {
	p1.eat(notEating, wg)
	p2.eat(notEating, wg)
	wg.Done()
}



func main() {
	notEating 	:= make(chan Philo, 5)
	var p1Granted, p2Granted Philo
	var wg sync.WaitGroup
	var canEat, ok bool	
	
	_, philos, host := InitData()

	startDinner(notEating, philos)
	fmt.Println(len(notEating))

	fmt.Println("Host will begin dinner now")

	
	for {
		
		if (len(notEating) > 1) {			
			select {
				case p1Granted = <- notEating:
					p2Granted, ok = <- notEating
					if (ok) {

						canEat = canPhilosEatTogether(notEating, p1Granted, p2Granted, &wg)
						
						if (canEat) {
							
							wg.Add(1)
							go host.allowToEat(notEating, p1Granted, p2Granted, &wg)				
							wg.Wait()

						} else {
							
							notEating <- p2Granted
							notEating <- p1Granted	
						}
					} else {
						notEating <- p1Granted
					}				
			}
		} 
		
		if (len(notEating) == 1) {
			p1Granted = <- notEating			
			p1Granted.eat(notEating, &wg)			
		}		

		if (len(notEating) == 0) {
			fmt.Println("Diner has finished")
			break
		}
	}
}

func InitData() ([]*ChopS, []Philo, *Host) {
	
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := philoList(CSticks, 5, true)
	
	host 	:= new(Host)
	
	return CSticks, philos, host
}

// philoList creates a []Philo slice of a given qty int value
func philoList(CSticks []*ChopS, qty int, setNum bool) []Philo {
	fmt.Println()
	philoList := make([]Philo, qty)
	for i := 0; i < qty; i++ {
		philoList[i].EatTimes = 3
		if (setNum) {
			philoList[i].PNum = float64(i + 1)
			
		}
		philoList[i].leftCS 	= CSticks[i]
		philoList[i].rightCS 	= CSticks[(i+1)%qty]
	}
	return philoList
}

// startDinner initializes channel notEating with a given []Philo slice
func startDinner(notEating chan Philo, philos []Philo) {
	fmt.Println("Dinner has started")
	for i := 0; i < 5; i++ {
		notEating <- philos[i]
		fmt.Printf("Philosopher %v not eating", i + 1)
		fmt.Println()
	}
}

// canPhilosEatTogether returns true if philosophers pair given isn't seated together
func canPhilosEatTogether(notEating chan Philo, phil1 Philo, phil2 Philo, wg *sync.WaitGroup) bool {
	
	calc := math.Abs(phil1.PNum - phil2.PNum)

	if calc > 1 && calc != 4 {
		if phil1.EatTimes != 0 && phil2.EatTimes != 0 {
			
			return true
		} else {
			if phil1.EatTimes == 0 && phil2.EatTimes != 0{
				return false
			}

			if phil1.EatTimes != 0 && phil2.EatTimes == 0{
				return false
			} 

			if phil1.EatTimes == 0 && phil2.EatTimes == 0 {
				return false
			}
			
		}
		
	} 
	return false
}