package concurrent

import (
	"testing"
	"time"
)

func TestWaitGroup_Run(t *testing.T) {
	wg := NewWaitGroup()
	wg.Run(func() {
		panic(1)
	})
	wg.Run(func() {
		<-time.After(time.Second)
	})
	wg.Wait()
}
