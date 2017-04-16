package main

import (
	"fmt"
	"github.com/house-lee/SoarGO/dsa"
	"sync"
	"time"
)

var gq dsa.IQueue
var testArr []bool

const (
	maxTests  = 20000
	writerCnt = 10
	readerCnt = 10
)

func main() {
	testArr = make([]bool, maxTests)
	for i := 0; i != maxTests; i++ {
		testArr[i] = false
	}
	gq = dsa.NewQueue(readerCnt, 0)

	var wg sync.WaitGroup
	for i := 0; i != writerCnt; i++ {
		wg.Add(1)
		go func(idx int) {
			numPerThread := maxTests / writerCnt
			start := idx * numPerThread
			time.Sleep(2 * time.Second)
			fmt.Println(start, numPerThread)
			for j := 0; j != numPerThread; j++ {
				if err := gq.Enqueue(start + j); err != nil {
					fmt.Println(err.Error())
				}
			}
			wg.Done()
		}(i)
	}

	for i := 0; i != readerCnt; i++ {
		wg.Add(1)
		go func() {
			var idx int
			numPerThread := maxTests / readerCnt
			for j := 0; j != numPerThread; j++ {
				idx = gq.Dequeue().(int)
				testArr[idx] = true
			}
			wg.Done()
		}()
	}
	wg.Wait()
	//for i := 0;i != maxTests ;i++  {
	//	idx := gq.Dequeue().(int)
	//	testArr[idx] = true
	//}
	for i := 0; i != maxTests; i++ {
		if !testArr[i] {
			fmt.Printf("%d not set\n", i)
		}
	}
	fmt.Println("Done")
}
