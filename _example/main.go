package main

import (
	"fmt"
	"time"

	"github.com/janeczku/go-spinner"
)

func NewSpinner(message string) *spinner.Spinner {
	s := spinner.NewSpinner(message)
	s.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})
	return s
}

func main() {
	fmt.Println("This may take some time:")

	// Task 1
	s := NewSpinner("Task 1: Processing...")
	s.Start()
	time.Sleep(3 * time.Second)
	s.Stop()
	fmt.Println("✓ Task 1: Completed")

	// Task 2
	s = NewSpinner("Task 2: Processing...")
	s.Start()
	s.SetSpeed(100 * time.Millisecond)
	time.Sleep(3 * time.Second)
	s.Stop()
	fmt.Println("✓ Task 2: Completed")

	// Task 3
	s = NewSpinner("Task 3: Testing Success (No message)")
	s.Start()
	s.SetSpeed(100 * time.Millisecond)
	time.Sleep(3 * time.Second)
	s.Success()

	// Task 4
	s = NewSpinner("Task 4: Testing Success (New Message)")
	s.Start()
	s.SetSpeed(100 * time.Millisecond)
	time.Sleep(3 * time.Second)
	s.Success("Task 4: I am a success message!")

	// Task 5
	s = NewSpinner("Task 5: Testing Error (No message)")
	s.Start()
	s.SetSpeed(100 * time.Millisecond)
	time.Sleep(3 * time.Second)
	s.Error()

	// Task 6
	s = NewSpinner("Task 6: Testing Error (New Message)")
	s.Start()
	s.SetSpeed(100 * time.Millisecond)
	time.Sleep(3 * time.Second)
	s.Error("Task 6: I am an error message!")

	fmt.Println("All Done!")
}
