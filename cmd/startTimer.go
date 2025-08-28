package cmd

import (
	"fmt"
	"time"
)

func StartTimer(seconds int, done chan bool) {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for i := seconds; i > 0; i-- {
		<-ticker.C // wait for 1 second
		fmt.Printf("⏳ %d seconds left\n", i)
	}
	fmt.Println("⏰ Time's up!")
	done <- true
}
