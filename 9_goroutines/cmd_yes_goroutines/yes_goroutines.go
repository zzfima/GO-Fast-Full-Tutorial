package yes_goroutines

import (
	"fmt"
	"sync"
	"time"
)

var dbData []string = []string{"id1", "id2", "id3", "id4", "id5", "id6"}
var res = []string{}
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func DoYesGoroutines() {
	fmt.Println("***YesNoGoroutines***")
	t0 := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("time passed ", time.Since(t0))
	fmt.Println(res)
}

func dbCall(i int) {
	var delay float32 = 1000
	duration := time.Duration(delay) * time.Millisecond
	fmt.Println("DB read time is ", duration)
	time.Sleep(duration)
	fmt.Println("DB read is ", dbData[i])

	m.Lock()
	res = append(res, dbData[i])
	m.Unlock()

	//will locked when Lock
	m.RLock()
	fmt.Println(res)
	m.RUnlock()

	wg.Done()
}
