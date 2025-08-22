package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func Routine(wg *sync.WaitGroup, nama string, durasi time.Duration) {
	defer wg.Done()
	start := time.Now()
	fmt.Printf("[%s] Mulai %s...\n", start.Format("15:04:05"), nama)

	time.Sleep(durasi)

	end := time.Now()
	fmt.Printf("[%s] Selesai %s\n", end.Format("15:04:05"), nama)
}
