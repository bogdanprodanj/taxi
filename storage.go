package main

import (
	"math/rand"
	"sync"
	"time"
)

type storage struct {
	sync.Mutex
	active   []order
	canceled []order
}

type order struct {
	name  string
	count int
}

func newStorage() *storage {
	s := &storage{
		active: make([]order, 50),
	}
	for i := 0; i < len(s.active); i++ {
		s.active[i] = order{name: randomString(2)}
	}
	return s
}

func (s *storage) getAllRequests() (orders []order) {
	s.Lock()
	defer s.Unlock()
	for _, v := range append(s.canceled, s.active...) {
		if v.count != 0 {
			orders = append(orders, v)
		}
	}
	return
}

func (s *storage) makeRequest() string {
	s.Lock()
	defer s.Unlock()
	i := rand.Intn(50)
	s.active[i].count++
	return s.active[i].name
}

func (s *storage) cancel() {
	t := time.NewTicker(time.Millisecond * 200)
	var i int
loop:
	for {
		select {
		case <-t.C:
			s.Lock()
			i = rand.Intn(50)
			s.canceled = append(s.canceled, s.active[i])
			s.active[i].count = 0
			s.active[i].name = randomString(2)
			s.Unlock()
		case <-quit:
			break loop
		}
	}
	t.Stop()
}
