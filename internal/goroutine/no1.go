package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func Routine(wg *sync.WaitGroup, nama string, durasi time.Duration) {
	defer wg.Done()
	fmt.Printf("Mulai %s...\n", nama)
	time.Sleep(durasi)
	fmt.Printf("Selesai %s\n", nama)
}
