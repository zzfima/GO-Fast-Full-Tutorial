package yes_goroutines

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var dbData []string = []string{"id1", "id2", "id3", "id4", "id5", "id6"}
var wg = sync.WaitGroup{}

func DoYesGoroutines() {
	fmt.Println("***YesNoGoroutines***")
	t0 := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("time passed ", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	duration := time.Duration(delay) * time.Millisecond
	fmt.Println("DB read time is ", duration)
	time.Sleep(duration)
	fmt.Println("DB read is ", dbData[i])
	wg.Done()
}
