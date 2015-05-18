package main

import (
	"fmt"
	"time"

	"github.com/janeczku/go-spinner"
)

func main() {
	fmt.Println("This may take some time:")
	s := spinner.StartNew("Task 1: Processing...")
	//s.Start()
	time.Sleep(3 * time.Second)
	s.Stop()
	fmt.Println("✓ Task 1: Completed")
	//time.Sleep(1 * time.Second)
	s = spinner.NewSpinner("Task 2: Processing...")
	s.Start()
	s.SetSpeed(100 * time.Millisecond)
	s.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})
	time.Sleep(3 * time.Second)
	s.Stop()
	fmt.Println("✓ Task 2: Completed")
	fmt.Println("All Done!")
}
