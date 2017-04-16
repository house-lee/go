package main

import (
	"fmt"
	"github.com/house-lee/SoarGO/dsa"
	"sync"
    "math/rand"
    "time"
)

var gq dsa.IQueue

func main() {

    rand.Seed(time.Now().Unix())
	gq = dsa.NewQueue(10, 1024)
	var wg sync.WaitGroup
	for i := 0; i != 100; i++ {
		wg.Add(1)
		go func() {
            for j := 0; j != 10; j++ {
                randint := rand.Intn(100000)
                fmt.Println(randint)
                gq.Enqueue(randint)
            }
            wg.Done()
		}()
	}

	for i := 0; i != 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j != 100; j++ {
				item := gq.Dequeue().(int)
				fmt.Println(item)
			}
            wg.Done()
		}()
	}
	wg.Wait()
}
