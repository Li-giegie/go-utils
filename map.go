package go_utils

import "sync"

// MapString 并发安全键string map
type MapString struct {
	sync.RWMutex
	M map[string]interface{}
}

type MapStringSlice struct {
	Key   string
	Value interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapString(n ...int) *MapString {
	if len(n) == 0 {
		n = []int{0}
	}
	m := new(MapString)
	m.M = make(map[string]interface{}, n[0])
	return m
}

func (m *MapString) Get(k string) (interface{}, bool) {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	i, ok := m.M[k]
	return i, ok
}

func (m *MapString) Set(k string, i interface{}) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	m.M[k] = i
	return
}

func (m *MapString) Delete(k string) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	delete(m.M, k)
}

func (m *MapString) GetMap() map[string]interface{} {
	return m.M
}

func (m *MapString) ToSlice() []*MapStringSlice {
	var ms = make([]*MapStringSlice, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, v := range m.M {
		ms[i] = &MapStringSlice{
			Key:   k,
			Value: v,
		}
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapString) KeyToSlice() []string {
	var ms = make([]string, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, _ := range m.M {
		ms[i] = k
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapString) ValueToSlice() []interface{} {
	var ms = make([]interface{}, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, _ := range m.M {
		ms[i] = m.M[k]
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapString) Range(call func(k string, v interface{})) {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	for s, i := range m.M {
		call(s, i)
	}
}

// MapString 并发安全键string map
type MapInt64 struct {
	sync.RWMutex
	M map[int64]interface{}
}

type MapInt64Slice struct {
	Key   int64
	Value interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapInt64(n ...int) *MapInt64 {
	if len(n) == 0 {
		n = []int{0}
	}
	m := new(MapInt64)
	m.M = make(map[int64]interface{}, n[0])
	return m
}

func (m *MapInt64) Get(k int64) (interface{}, bool) {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	i, ok := m.M[k]
	return i, ok
}

func (m *MapInt64) Set(k int64, i interface{}) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	m.M[k] = i
	return
}

func (m *MapInt64) Delete(k int64) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	delete(m.M, k)
}

func (m *MapInt64) GetMap() map[int64]interface{} {
	return m.M
}

func (m *MapInt64) ToSlice() []*MapInt64Slice {
	var ms = make([]*MapInt64Slice, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, v := range m.M {
		ms[i] = &MapInt64Slice{
			Key:   k,
			Value: v,
		}
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapInt64) KeyToSlice() []int64 {
	var ms = make([]int64, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, _ := range m.M {
		ms[i] = k
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapInt64) ValueToSlice() []interface{} {
	var ms = make([]interface{}, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, _ := range m.M {
		ms[i] = m.M[k]
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapInt64) Range(call func(k int64, v interface{})) {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	for s, i := range m.M {
		call(s, i)
	}
}

// MapString 并发安全键string map
type MapUint64 struct {
	sync.RWMutex
	M map[uint64]interface{}
}

type MapUint64Slice struct {
	Key   uint64
	Value interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapUint64(n ...int) *MapUint64 {
	if len(n) == 0 {
		n = []int{0}
	}
	m := new(MapUint64)
	m.M = make(map[uint64]interface{}, n[0])
	return m
}

func (m *MapUint64) Get(k uint64) (interface{}, bool) {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	i, ok := m.M[k]
	return i, ok
}

func (m *MapUint64) Set(k uint64, i interface{}) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	m.M[k] = i
	return
}

func (m *MapUint64) Delete(k uint64) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	delete(m.M, k)
}

func (m *MapUint64) GetMap() map[uint64]interface{} {
	return m.M
}

func (m *MapUint64) ToSlice() []*MapUint64Slice {
	var ms = make([]*MapUint64Slice, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, v := range m.M {
		ms[i] = &MapUint64Slice{
			Key:   k,
			Value: v,
		}
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapUint64) KeyToSlice() []uint64 {
	var ms = make([]uint64, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, _ := range m.M {
		ms[i] = k
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapUint64) ValueToSlice() []interface{} {
	var ms = make([]interface{}, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, _ := range m.M {
		ms[i] = m.M[k]
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapUint64) Range(call func(k uint64, v interface{})) {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	for s, i := range m.M {
		call(s, i)
	}
}

// MapString 并发安全键string map
type MapInt32 struct {
	sync.RWMutex
	M map[int32]interface{}
}

type MapInt32Slice struct {
	Key   int32
	Value interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapInt32(n ...int) *MapInt32 {
	if len(n) == 0 {
		n = []int{0}
	}
	m := new(MapInt32)
	m.M = make(map[int32]interface{}, n[0])
	return m
}

func (m *MapInt32) Get(k int32) (interface{}, bool) {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	i, ok := m.M[k]
	return i, ok
}

func (m *MapInt32) Set(k int32, i interface{}) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	m.M[k] = i
	return
}

func (m *MapInt32) Delete(k int32) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	delete(m.M, k)
}

func (m *MapInt32) GetMap() map[int32]interface{} {
	return m.M
}

func (m *MapInt32) ToSlice() []*MapInt32Slice {
	var ms = make([]*MapInt32Slice, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, v := range m.M {
		ms[i] = &MapInt32Slice{
			Key:   k,
			Value: v,
		}
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapInt32) KeyToSlice() []int32 {
	var ms = make([]int32, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, _ := range m.M {
		ms[i] = k
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapInt32) ValueToSlice() []interface{} {
	var ms = make([]interface{}, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, _ := range m.M {
		ms[i] = m.M[k]
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapInt32) Range(call func(k int32, v interface{})) {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	for s, i := range m.M {
		call(s, i)
	}
}

// MapString 并发安全键string map
type MapUint32 struct {
	sync.RWMutex
	M map[uint32]interface{}
}

type MapUint32Slice struct {
	Key   uint32
	Value interface{}
}

// NewMapString 创建一个并发安全键string map ,n 为map初始容量
func NewMapUint32(n ...int) *MapUint32 {
	if len(n) == 0 {
		n = []int{0}
	}
	m := new(MapUint32)
	m.M = make(map[uint32]interface{}, n[0])
	return m
}

func (m *MapUint32) Get(k uint32) (interface{}, bool) {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	i, ok := m.M[k]
	return i, ok
}

func (m *MapUint32) Set(k uint32, i interface{}) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	m.M[k] = i
	return
}

func (m *MapUint32) Delete(k uint32) {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	delete(m.M, k)
}

func (m *MapUint32) GetMap() map[uint32]interface{} {
	return m.M
}

func (m *MapUint32) ToSlice() []*MapUint32Slice {
	var ms = make([]*MapUint32Slice, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, v := range m.M {
		ms[i] = &MapUint32Slice{
			Key:   k,
			Value: v,
		}
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapUint32) KeyToSlice() []uint32 {
	var ms = make([]uint32, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, _ := range m.M {
		ms[i] = k
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapUint32) ValueToSlice() []interface{} {
	var ms = make([]interface{}, len(m.M))
	var i int
	m.RWMutex.RLock()
	for k, _ := range m.M {
		ms[i] = m.M[k]
		i++
	}
	m.RWMutex.RUnlock()
	return ms
}

func (m *MapUint32) Range(call func(k uint32, v interface{})) {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	for s, i := range m.M {
		call(s, i)
	}
}
