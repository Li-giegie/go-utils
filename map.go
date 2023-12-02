package go_utils

import "sync"

type MapStringI interface {
	Get(k string) (interface{}, bool)
	Set(k string, i interface{})
	Delete(k string)
	GetMap() map[string]interface{}
	ToSlice() []*MapStringSlice
	KeyToSlice() []string
	ValueToSlice() []interface{}
}

// MapString 并发安全键string map
type MapString struct {
	lock  sync.RWMutex
	cache map[string]interface{}
}

type MapStringSlice struct {
	Key   string
	Value interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapString(n ...int) MapStringI {
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

func (m *MapString) ToSlice() []*MapStringSlice {
	var ms = make([]*MapStringSlice, len(m.cache))
	var i int
	for k, v := range m.cache {
		ms[i] = &MapStringSlice{
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

type MapInt64I interface {
	Get(k int64) (interface{}, bool)
	Set(k int64, i interface{})
	Delete(k int64)
	GetMap() map[int64]interface{}
	ToSlice() []*MapInt64Slice
	KeyToSlice() []int64
	ValueToSlice() []interface{}
}

// MapString 并发安全键string map
type MapInt64 struct {
	lock  sync.RWMutex
	cache map[int64]interface{}
}

type MapInt64Slice struct {
	Key   int64
	Value interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapInt64(n ...int) MapInt64I {
	if len(n) == 0 {
		n = []int{0}
	}
	m := new(MapInt64)
	m.cache = make(map[int64]interface{}, n[0])
	return m
}

func (m *MapInt64) Get(k int64) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	i, ok := m.cache[k]
	return i, ok
}

func (m *MapInt64) Set(k int64, i interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.cache[k] = i
	return
}

func (m *MapInt64) Delete(k int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.cache, k)
}

func (m *MapInt64) GetMap() map[int64]interface{} {
	return m.cache
}

func (m *MapInt64) ToSlice() []*MapInt64Slice {
	var ms = make([]*MapInt64Slice, len(m.cache))
	var i int
	for k, v := range m.cache {
		ms[i] = &MapInt64Slice{
			Key:   k,
			Value: v,
		}
		i++
	}
	return ms
}

func (m *MapInt64) KeyToSlice() []int64 {
	var ms = make([]int64, len(m.cache))
	var i int
	for k, _ := range m.cache {
		ms[i] = k
		i++
	}
	return ms
}

func (m *MapInt64) ValueToSlice() []interface{} {
	var ms = make([]interface{}, len(m.cache))
	var i int
	for k, _ := range m.cache {
		ms[i] = m.cache[k]
		i++
	}
	return ms
}

type MapUint64I interface {
	Get(k uint64) (interface{}, bool)
	Set(k uint64, i interface{})
	Delete(k uint64)
	GetMap() map[uint64]interface{}
	ToSlice() []*MapUint64Slice
	KeyToSlice() []uint64
	ValueToSlice() []interface{}
}

// MapString 并发安全键string map
type MapUint64 struct {
	lock  sync.RWMutex
	cache map[uint64]interface{}
}

type MapUint64Slice struct {
	Key   uint64
	Value interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapUint64(n ...int) MapUint64I {
	if len(n) == 0 {
		n = []int{0}
	}
	m := new(MapUint64)
	m.cache = make(map[uint64]interface{}, n[0])
	return m
}

func (m *MapUint64) Get(k uint64) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	i, ok := m.cache[k]
	return i, ok
}

func (m *MapUint64) Set(k uint64, i interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.cache[k] = i
	return
}

func (m *MapUint64) Delete(k uint64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.cache, k)
}

func (m *MapUint64) GetMap() map[uint64]interface{} {
	return m.cache
}

func (m *MapUint64) ToSlice() []*MapUint64Slice {
	var ms = make([]*MapUint64Slice, len(m.cache))
	var i int
	for k, v := range m.cache {
		ms[i] = &MapUint64Slice{
			Key:   k,
			Value: v,
		}
		i++
	}
	return ms
}

func (m *MapUint64) KeyToSlice() []uint64 {
	var ms = make([]uint64, len(m.cache))
	var i int
	for k, _ := range m.cache {
		ms[i] = k
		i++
	}
	return ms
}

func (m *MapUint64) ValueToSlice() []interface{} {
	var ms = make([]interface{}, len(m.cache))
	var i int
	for k, _ := range m.cache {
		ms[i] = m.cache[k]
		i++
	}
	return ms
}

type MapInt32I interface {
	Get(k int32) (interface{}, bool)
	Set(k int32, i interface{})
	Delete(k int32)
	GetMap() map[int32]interface{}
	ToSlice() []*MapInt32Slice
	KeyToSlice() []int32
	ValueToSlice() []interface{}
}

// MapString 并发安全键string map
type MapInt32 struct {
	lock  sync.RWMutex
	cache map[int32]interface{}
}

type MapInt32Slice struct {
	Key   int32
	Value interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapInt32(n ...int) MapInt32I {
	if len(n) == 0 {
		n = []int{0}
	}
	m := new(MapInt32)
	m.cache = make(map[int32]interface{}, n[0])
	return m
}

func (m *MapInt32) Get(k int32) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	i, ok := m.cache[k]
	return i, ok
}

func (m *MapInt32) Set(k int32, i interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.cache[k] = i
	return
}

func (m *MapInt32) Delete(k int32) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.cache, k)
}

func (m *MapInt32) GetMap() map[int32]interface{} {
	return m.cache
}

func (m *MapInt32) ToSlice() []*MapInt32Slice {
	var ms = make([]*MapInt32Slice, len(m.cache))
	var i int
	for k, v := range m.cache {
		ms[i] = &MapInt32Slice{
			Key:   k,
			Value: v,
		}
		i++
	}
	return ms
}

func (m *MapInt32) KeyToSlice() []int32 {
	var ms = make([]int32, len(m.cache))
	var i int
	for k, _ := range m.cache {
		ms[i] = k
		i++
	}
	return ms
}

func (m *MapInt32) ValueToSlice() []interface{} {
	var ms = make([]interface{}, len(m.cache))
	var i int
	for k, _ := range m.cache {
		ms[i] = m.cache[k]
		i++
	}
	return ms
}

type MapUint32I interface {
	Get(k uint32) (interface{}, bool)
	Set(k uint32, i interface{})
	Delete(k uint32)
	GetMap() map[uint32]interface{}
	ToSlice() []*MapUint32Slice
	KeyToSlice() []uint32
	ValueToSlice() []interface{}
}

// MapString 并发安全键string map
type MapUint32 struct {
	lock  sync.RWMutex
	cache map[uint32]interface{}
}

type MapUint32Slice struct {
	Key   uint32
	Value interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapUint32(n ...int) MapUint32I {
	if len(n) == 0 {
		n = []int{0}
	}
	m := new(MapUint32)
	m.cache = make(map[uint32]interface{}, n[0])
	return m
}

func (m *MapUint32) Get(k uint32) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	i, ok := m.cache[k]
	return i, ok
}

func (m *MapUint32) Set(k uint32, i interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.cache[k] = i
	return
}

func (m *MapUint32) Delete(k uint32) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.cache, k)
}

func (m *MapUint32) GetMap() map[uint32]interface{} {
	return m.cache
}

func (m *MapUint32) ToSlice() []*MapUint32Slice {
	var ms = make([]*MapUint32Slice, len(m.cache))
	var i int
	for k, v := range m.cache {
		ms[i] = &MapUint32Slice{
			Key:   k,
			Value: v,
		}
		i++
	}
	return ms
}

func (m *MapUint32) KeyToSlice() []uint32 {
	var ms = make([]uint32, len(m.cache))
	var i int
	for k, _ := range m.cache {
		ms[i] = k
		i++
	}
	return ms
}

func (m *MapUint32) ValueToSlice() []interface{} {
	var ms = make([]interface{}, len(m.cache))
	var i int
	for k, _ := range m.cache {
		ms[i] = m.cache[k]
		i++
	}
	return ms
}
