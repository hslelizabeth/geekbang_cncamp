package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []int
	cond  *sync.Cond
}

func main() {
	q := Queue{
		queue: []int{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}
	go func() {
		for i := 0; i < 10; i++ {
			q.Enqueue(i)
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		fmt.Println(q.Dequeue())
		time.Sleep(time.Second)
	}
}

func (q *Queue) Enqueue(number int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, number)
	fmt.Printf("putting %d to queue, notify all\n", number)
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.queue) == 0 {
		fmt.Println("no data available, wait")
		q.cond.Wait()
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}
