package main

import (
	"context"
	"math"
	"math/rand"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}
var rwmutex = sync.RWMutex{}
var readyCond = sync.NewCond(rwmutex.RLocker())

var once = sync.Once{}

var squares = map[int]int{}

func processRequest(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			Printfln("Stopping processing - request cancelled")
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}

func generateSquares(max int) {
	// rwmutex.Lock()
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	// rwmutex.Unlock()
	// Printfln("Broadcasting condition")
	// readyCond.Broadcast()
	// waitGroup.Done()
}

func readSquares(id, max, iterations int) {
	once.Do(func() {
		generateSquares(max)
	})
	// readyCond.L.Lock()
	// for len(squares) == 0 {
	//     readyCond.Wait()
	// }
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	// readyCond.L.Unlock()
	waitGroup.Done()
}

func calculateSquares(max, iterations int) {
	for i := 0; i < iterations; i++ {
		val := rand.Intn(max)
		rwmutex.RLock()
		square, ok := squares[val]
		rwmutex.RUnlock()
		if ok {
			Printfln("Cached value: %v = %v", val, square)
		} else {
			rwmutex.Lock()
			if _, ok := squares[val]; !ok {
				squares[val] = int(math.Pow(float64(val), 2))
				Printfln("Added value: %v = %v", val, squares[val])
			}
			rwmutex.Unlock()
		}
	}
	waitGroup.Done()
}

func doSum(count int, val *int) {
	time.Sleep(time.Second)
	for i := 0; i < count; i++ {
		*val++
	}
	waitGroup.Done()
}

func main() {
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	ctx, cancel := context.WithCancel(context.Background())
	go processRequest(ctx, &waitGroup, 10)

	time.Sleep(time.Second)
	Printfln("Canceling request")
	cancel()

	waitGroup.Wait()
}
