package go_utils

import "sync"

// MapString 并发安全键string map
type MapString struct {
	lock  sync.RWMutex
	cache map[string]interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapString(n ...int) *MapString {
	if len(n) == 0 {
		n = []int{0}
	}
	m := new(MapString)
	m.cache = make(map[string]interface{}, n[0])
	return m
}

func (m *MapString) Get(k string) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	i, ok := m.cache[k]
	return i, ok
}

func (m *MapString) Set(k string, i interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.cache[k] = i
	return
}

func (m *MapString) Delete(k string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.cache, k)
}

func (m *MapString) GetMap() map[string]interface{} {
	return m.cache
}

type MapSlice struct {
	Key   string
	Value interface{}
}

func (m *MapString) ToSlice() []*MapSlice {
	var ms = make([]*MapSlice, len(m.cache))
	var i int
	for k, v := range m.cache {
		ms[i] = &MapSlice{
			Key:   k,
			Value: v,
		}
		i++
	}
	return ms
}

func (m *MapString) KeyToSlice() []string {
	var ms = make([]string, len(m.cache))
	var i int
	for k, _ := range m.cache {
		ms[i] = k
		i++
	}
	return ms
}

func (m *MapString) ValueToSlice() []interface{} {
	var ms = make([]interface{}, len(m.cache))
	var i int
	for k, _ := range m.cache {
		ms[i] = m.cache[k]
		i++
	}
	return ms
}
