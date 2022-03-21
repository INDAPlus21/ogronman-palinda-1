package main

import (
	"fmt"
	"time"
)

func main() {
	go Remind("Time to eat", time.Second*10)
	go Remind("Time to work", time.Second*30)
	Remind("Time to sleep", time.Second*60)
}

func Remind(text string, delay time.Duration) {
	for {
		fmt.Printf("The time is %s: %s\n", time.Now().Format("15.04.05"), text)
		time.Sleep(delay)
	}
}
