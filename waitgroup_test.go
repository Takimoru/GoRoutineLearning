package GoRoutineLearn

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// SYNC.GROUP

// Digunakan untuk menunggu proses func/goroutine dapat di costum brp jumlah nya lewat group.Add()

//harus di tutup dengan group.Done jika tidak akan terjadi deadlock

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("halo")
	time.Sleep(1 * time.Second)

}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("kelar")

}

// SYNC ONCE

// Digunakan jika ingin menjalankan function hanya sekali (meski perulangan go routine banyak hanya mengeksekusi 1 function)

var counter = 0

func onlyOnce() {
	counter++

}

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 50; i++ {
		go func() {
			group.Add(1)
			once.Do(onlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("counter : ", counter)

}
