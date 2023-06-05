package main

import (
	"fmt"
	"math/rand"
  "sync"
)

type ProcessA struct {
	Weight int
}

type ProcessB struct {
	Color string
}

type ProcessC struct {
	Alive bool
}

func main() {
  
	// Create channels for communication
	A_to_B := make(chan int)
	B_to_C := make(chan string)
	C_to_A := make(chan bool)

  var wg sync.WaitGroup
  
  // Create and run three processes
  wg.Add(1)
  wg.Add(1)
  wg.Add(1)
  
  go func (from_C <-chan bool, to_B chan<- int) {

    defer wg.Done()
    
  	a := ProcessA{Weight: 10}

   //  // send my initial state
  	// to_B <- a.Weight
    
  	for {
  		alive := <-from_C
  		fmt.Printf("ProcessA: Received from ProcessC: Alive = %t\n", alive)

      new_weight := rand.Intn(10)
      
      // if a.Weight == new_weight {
      //   break
      // }
      
      a.Weight = new_weight
  		fmt.Printf("ProcessA: Weight = %d\n", a.Weight)
  		to_B <- a.Weight
  	}
  }(C_to_A, A_to_B)
  
  go func (from_A <-chan int, to_C chan<- string) {

    defer wg.Done()
    
  	b := ProcessB{Color: "blue"}

    // send my initial state
 		to_C <- b.Color

 		colors := []string{"red", "green", "blue", "N/A"}
    
  	for {
  		weight := <-from_A
  		fmt.Printf("ProcessB: Received from ProcessA: Weight = %d\n", weight)
  
  		new_color := colors[rand.Intn(len(colors))]
      
      // if b.Color == new_color {
      //   break
      // }
      
  		b.Color = new_color
  		fmt.Printf("ProcessB: Color = %s\n", b.Color)
  		to_C <- b.Color
  	}
  }(A_to_B, B_to_C)
  
  go func (from_B <-chan string, to_A chan<- bool) {

    defer wg.Done()
    
  	c := ProcessC{Alive: true}

    // send my initial state
 		to_A <- c.Alive
    
  	for {
  		color := <-from_B
  		fmt.Printf("ProcessC: Received from ProcessB: Color = %s\n", color)
  
  		new_alive := rand.Intn(2) == 1

      // if c.Alive == new_alive {
      //   break
      // }
      
  		c.Alive = new_alive
  		fmt.Printf("ProcessC: Received from ProcessB: Alive = %t\n", c.Alive)
  		to_A <- c.Alive
  	}
  }(B_to_C, C_to_A)
  
  wg.Wait()
	fmt.Println("All goroutines finished")
}