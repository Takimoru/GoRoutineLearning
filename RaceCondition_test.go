package GoRoutineLearn

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("counter = ", x)
}

//jikalau seperti ini saat go routine di jalankan skenarionya 1 variable akan di pakai berbarengan oleh go routine yang lain
// mengakibatkan hasilnya yang harusnya 100.000 tidak sampai segitu ini yang di namakan **RaceCondition**
