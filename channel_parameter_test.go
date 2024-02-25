package GoRoutineLearn

import (
	"fmt"
	"testing"
	"time"
)

func TestBuatchanel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "apa iyh deck"
		fmt.Println("kelar buat channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "dini jangan marah yak"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string, 5)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestBuffCh(t *testing.T) {
	cenel := make(chan string)
	defer close(cenel)

	//ini kalo pake go routine + anonymous func
	go func() {
		cenel <- "Suidini"
		cenel <- "Pekodini"
		cenel <- "Wanwandini"
		cenel <- "Namidini"
		cenel <- "Suichan waaaa"
	}()

	go func() {
		fmt.Println(<-cenel)
		fmt.Println(<-cenel)
		fmt.Println(<-cenel)
		fmt.Println(<-cenel)
		fmt.Println(<-cenel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Waifu gwe")
}
