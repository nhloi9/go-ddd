package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Unsafe map without locks
	fmt.Println("=== UNSAFE MAP (will crash) ===")
	unsafeMapDemo()

	fmt.Println("\n=== SAFE MAP (with locks) ===")
	safeMapDemo()
}

func unsafeMapDemo() {
	m := make(map[string]int)
	
	// Start multiple goroutines writing to the same map
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 1000; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				m[key] = id*1000 + j // THIS WILL CAUSE PANIC!
			}
		}(i)
	}
	
	time.Sleep(100 * time.Millisecond) // Give goroutines time to race
	fmt.Printf("Map size: %d\n", len(m))
}

func safeMapDemo() {
	m := make(map[string]int)
	var mu sync.Mutex
	
	var wg sync.WaitGroup
	
	// Start multiple goroutines writing to the same map SAFELY
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				
				mu.Lock()
				m[key] = id*1000 + j // SAFE with mutex
				mu.Unlock()
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Map size: %d\n", len(m))
	fmt.Println("All writes completed successfully!")
}