package pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Payload1 struct {
	locker *sync.Mutex
	n      int
}

func (p *Payload1) Payload() {
	p.locker.Lock()
	p.n = p.n + 1
	fmt.Printf("%d\n", p.n)
	p.locker.Unlock()
}

func Test1(t *testing.T) {
	p := NewPool(2, 2)
	go p.Run()
	job := &Payload1{&sync.Mutex{}, 0}
	p.Submit(job)
	p.Submit(job)
	p.Submit(job)

	time.Sleep(time.Second * 3)
	exp := 3
	if exp != job.n {
		t.Errorf("exp %d, actual %d\n", exp, job.n)
	}
}
