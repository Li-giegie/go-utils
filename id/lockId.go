package id

import "sync"

type LockId struct {
	*Range
	Id uint
	lock sync.Mutex
}

func NewLockId(_range ...Range) *LockId {
	r := getRangeSile(_range)
	return &LockId{
		Range: r,
		lock:  sync.Mutex{},
		Id: r.Start,
	}
}

func (l *LockId) Get() uint {
	l.lock.Lock()
	defer l.lock.Unlock()
	n := l.Id
	l.Id += l.Step
	if l.Id > l.End {
		l.Id = l.Start
	}
	return n
}
