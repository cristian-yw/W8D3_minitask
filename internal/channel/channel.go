package channel

import "fmt"

type Msg struct {
	Pengirim string
	Isi      string
}

func Pesan(ch chan<- Msg, pengirim string, isi string) {
	pesan := Msg{Pengirim: pengirim, Isi: isi}
	ch <- pesan
}

func Papan(ch <-chan Msg) {
	for pesan := range ch {
		fmt.Println(pesan.Pengirim, ":", pesan.Isi)
	}
	fmt.Println("Semoga sukses")
}
