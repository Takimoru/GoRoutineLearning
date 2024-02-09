package GoRoutineLearn

import (
	"fmt"
	"testing"
	"time"
)

func suiseisayang() {
	fmt.Println("Hai honey")
}

func TestGoRoutine(t *testing.T) {
	go suiseisayang()
	fmt.Println("nico pacar dini")

	time.Sleep(1 * time.Second)
}
