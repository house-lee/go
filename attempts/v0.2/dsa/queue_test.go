package dsa

import (
    "testing"
    "sync"
    "fmt"
)
const (
    testEnqueue = iota
    testEnqueueNDequeueSeparate
    testEnqueueNDequeueMix
)


func TestLfQueue(t *testing.T)  {
    err := launchTest(50,20,0,1000000, true, testEnqueueNDequeueMix)
    if err != nil {
        t.Error(err.Error())
    }
}

func BenchmarkLfQueueW50R20C10M(b *testing.B) {
    for i := 0; i != b.N; i++ {
        err := launchTest(50,20,0,1000000, true, testEnqueueNDequeueMix)
        if err != nil {
            b.Error(err.Error())
        }
    }
}
func BenchmarkLQueueW50R20C10M(b *testing.B) {
    for i := 0; i != b.N; i++ {
        err := launchTest(50,20,0,1000000, false, testEnqueueNDequeueMix)
        if err != nil {
            b.Error(err.Error())
        }
    }
}




func launchTest(writerCnt, readerCnt, maxItems, maxTests uint32, lockFree bool,  testMode int) error {
    testArr := make([]bool, maxTests)
    for i := 0; i != int(maxTests); i++ {
        testArr[i] = false
    }
    var gq IQueue
    if lockFree {
        gq = NewLockFreeQueue(readerCnt, maxItems)
    } else {
        gq = NewQueueWithLock(readerCnt, maxItems)
    }
    var wg sync.WaitGroup
    startSig := make(chan bool)
    for i := 0; i != int(writerCnt); i++ {
        wg.Add(1)
        go func(q IQueue, idx int) {
            <- startSig
            numberPerThread := int(maxTests/writerCnt)
            startOffset := idx * numberPerThread
            for j := 0; j != numberPerThread; j++ {
                if err := gq.Enqueue(startOffset + j); err != nil {
                    fmt.Println(err.Error())
                }
            }
            wg.Done()
        }(gq, i)
    }
    if testMode == testEnqueue {
        close(startSig) // trigger all threads start
        wg.Wait()
        for i := 0; i != int(maxTests); i++ {
            idx := gq.Dequeue().(int)
            testArr[idx] = true
        }
    } else if testMode == testEnqueueNDequeueSeparate {
        close(startSig) // trigger all threads start
        wg.Wait()
        for i := 0; i != int(readerCnt); i++ {
            wg.Add(1)
            go func() {
                var idx int
                numPerThread := int(maxTests / readerCnt)
                for j := 0; j != numPerThread ; j++ {
                    idx = gq.Dequeue().(int)
                    testArr[idx] = true
                }
                wg.Done()
            }()
        }
        wg.Wait()

    } else if testMode == testEnqueueNDequeueMix {
        for i := 0; i != int(readerCnt); i++ {
            wg.Add(1)
            go func() {
                var idx int
                numPerThread := int(maxTests / readerCnt)
                for j := 0; j != numPerThread ; j++ {
                    idx = gq.Dequeue().(int)
                    testArr[idx] = true
                }
                wg.Done()
            }()
        }
        close(startSig)
        wg.Wait()
    }
    for i := 0; i != int(maxTests); i++ {
        if !testArr[i] {
            return fmt.Errorf("%d not set\n", i)
        }
    }
    return nil
}