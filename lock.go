package gocache

import "sync/atomic"

type Lock struct {
	readerCount int32
	writeCount  int32
}

func (lock *Lock) BeginRead() {
	for {
		if atomic.LoadInt32(&lock.writeCount) == 0 {
			atomic.AddInt32(&lock.readerCount, 1)
			if atomic.LoadInt32(&lock.writeCount) == 0 {
				return
			} else {
				atomic.AddInt32(&lock.readerCount, -1)
			}
		}
	}
}

func (lock *Lock) EndRead() {
	atomic.AddInt32(&lock.readerCount, -1)
}

func (lock *Lock) BeginWrite() {
	for {
		if atomic.LoadInt32(&lock.writeCount) == 0 && atomic.LoadInt32(&lock.readerCount) == 0 {
			if atomic.CompareAndSwapInt32(&lock.writeCount, 0, 1) {
				return
			}
		}
	}
}

func (lock *Lock) EndWrite() {
	atomic.StoreInt32(&lock.writeCount, 0)
}
