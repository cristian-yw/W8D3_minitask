package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/cristian-yw/W8D3_minitask/internal/channel"
	"github.com/cristian-yw/W8D3_minitask/internal/goroutine"
	"github.com/cristian-yw/W8D3_minitask/internal/mutex"
)

func main() {
	var wg sync.WaitGroup

	routine := []struct {
		nama   string
		durasi time.Duration
	}{
		{"Mandi", 3 * time.Second},
		{"Makan", 2 * time.Second},
		{"Ngopi", 4 * time.Second},
		{"Menyapu", 1 * time.Second},
	}

	wg.Add(len(routine))
	for _, r := range routine {
		go goroutine.Routine(&wg, r.nama, r.durasi)
	}
	wg.Wait()
	fmt.Println("Berangkat Kerja")

	ch := make(chan channel.Msg)
	go channel.Papan(ch)

	go channel.Pesan(ch, "Cegans", "Anjir gw ganteng bet dah")
	time.Sleep(500 * time.Millisecond)

	go channel.Pesan(ch, "Miya", "Iya cok si ce ganteng bat cok")
	time.Sleep(500 * time.Millisecond)

	go channel.Pesan(ch, "Johnson", "Gila si ce ganteng banget cok")
	time.Sleep(500 * time.Millisecond)

	close(ch)

	time.Sleep(5 * time.Second)

	var wg2 sync.WaitGroup
	var mut sync.Mutex

	orang := []string{"Ce", "Oped", "Radif", "Sidiq", "Ucup"}

	wg2.Add(len(orang))
	for _, o := range orang {
		go mutex.Microwavew(o, &wg2, &mut)
	}
	wg2.Wait()
	fmt.Println("Sudah selesai")
}
