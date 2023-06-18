package lru

import (
	"github.com/erfanmomeniii/queue"
)

const DefaultSize int64 = 2000

// Lru is an instantiation of the lru.
type Lru struct {
	q    *queue.Queue
	m    map[any]any
	size int64
}

// Get returns the value associated with the inputted key.
func (l *Lru) Get(key any) any {
	return l.m[key]
}

// moveToFront moves key to the front of the queue.
func (l *Lru) moveToFront(key any) {
	q := queue.New()

	f := l.q.Front()
	for f != nil {
		if f != key {
			q.PushBack(f)
		}

		l.q.PopFront()
		f = l.q.Front()
	}

	q.PushFront(key)
	l.q = q

	return
}

// purge removes the last element from cache.
func (l *Lru) purge() {
	if l.q.Size() > int(l.size) {
		b := l.q.Back()
		l.q.PopBack()
		delete(l.m, b)
	}

	return
}

// Set adds or updates with inputted key and value.
func (l *Lru) Set(key any, value any) {
	if _, ok := l.m[key]; ok {
		l.moveToFront(key)
	} else {
		l.q.PushFront(key)
	}

	l.m[key] = value
	l.purge()

	return
}

// Has checks whether the key exists or not.
func (l *Lru) Has(key any) bool {
	_, ok := l.m[key]

	return ok
}

// New creates a new instance of a lru.
func New(size ...int64) *Lru {
	s := DefaultSize
	if len(size) > 0 {
		s = size[0]
	}

	return &Lru{q: queue.New(), m: make(map[any]any), size: s}
}
