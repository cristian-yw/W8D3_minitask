package mutex

import (
	"fmt"
	"sync"
	"time"
)

func Microwavew(nama string, wg *sync.WaitGroup, mut *sync.Mutex) {
	defer wg.Done()
	mut.Lock()
	fmt.Printf("%s ingin pake microwave\n", nama)
	time.Sleep(2 * time.Second)
	fmt.Printf("%s selesai pake microwave\n", nama)
	mut.Unlock()
}
