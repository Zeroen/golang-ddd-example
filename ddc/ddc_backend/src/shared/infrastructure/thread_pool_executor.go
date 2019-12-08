package infrastructure

import (
	"sync"
)

type Runnable func()

type ThreadPoolExecutor struct {
	mux sync.Mutex
	numberThreads int
	closed        bool
	exit          chan int
	task          chan Runnable
}

func NewThreadPoolExecutor(coreThreads int) *ThreadPoolExecutor {
	th := &ThreadPoolExecutor{
		numberThreads: coreThreads,
		exit:          make(chan int),
		task:          make(chan Runnable),
	}

	for i := 0; i < coreThreads; i++ {
		go func() {
			for {
				select {
				case r := <-th.task:
					r()
				case <-th.exit:
					return
				}
			}
		}()
	}

	return th
}

func (th *ThreadPoolExecutor) Submit(runnable Runnable) {
	th.mux.Lock()
	if !th.closed {
		th.task <- runnable
	}
	th.mux.Unlock()
}

func (th *ThreadPoolExecutor) Close() {
	th.mux.Lock()
	th.closed = true
	for i := 0; i < th.numberThreads; i++ {
		th.exit <- -1
	}
	th.mux.Unlock()
}
