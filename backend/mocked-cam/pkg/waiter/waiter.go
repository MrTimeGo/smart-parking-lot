package waiter

import (
	"math/rand"
	"time"
)

func WaitRandom(minDelay, maxDelay time.Duration) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		time.Sleep(minDelay + time.Duration(rand.Int63n(int64(maxDelay-minDelay))))
	}()
	return done
}
