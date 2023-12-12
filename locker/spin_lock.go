package locker

import (
	"runtime"
	"sync/atomic"
)

type Locker interface {
	Lock()
	Unlock()
}

type SpinLock uint32

const maxBackoff = 32

func (s *SpinLock) Lock() {
	backoff := 1

	// 尝试获取自旋锁
	for !atomic.CompareAndSwapUint32((*uint32)(s), 0, 1) { // 获取锁失败
		for i := 0; i < backoff; i++ {
			runtime.Gosched() // 让出cpu
		}

		// 简单指数退避算法
		if backoff < maxBackoff {
			backoff <<= 1
		}
	}
}

func (s *SpinLock) Unlock() {
	atomic.CompareAndSwapUint32((*uint32)(s), 1, 0)
}
