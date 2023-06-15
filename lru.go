package lru

import (
	"github.com/erfanmomeniii/queue"
)

const defaultSize int64 = 2000

type Lru struct {
	q    *queue.Queue
	m    map[any]any
	size int64
}

func New(size ...int64) *Lru {
	s := defaultSize
	if len(size) > 0 {
		s = size[0]
	}

	return &Lru{q: queue.New(), size: s}
}

func (l *Lru) Get(key any) any {
	return l.m[key]
}

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

func (l *Lru) purge() {
	if l.q.Size() > int(l.size) {
		b := l.q.Back()
		l.q.PopBack()
		delete(l.m, b)
	}

	return
}

func (l *Lru) Set(key any, value any) {
	if _, ok := l.m[key]; ok {
		l.moveToFront(key)
	}
	l.q.PushFront(key)
	l.m[key] = value
	l.purge()

	return
}
