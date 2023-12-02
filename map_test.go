package go_utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func TestNewMapString(t *testing.T) {
	m := NewMapString()
	m.Set("1", 123)
	m.Set("2", 456)
	m.Set("3", 789)
	fmt.Println(m.GetMap())
	fmt.Println(m.Get("1"))
	for _, slice := range m.ToSlice() {
		fmt.Println(slice)
	}
	for _, slice := range m.KeyToSlice() {
		fmt.Println(slice)
	}
	for _, i := range m.ValueToSlice() {
		fmt.Println(i)
	}
}

func TestNewMapInt64(t *testing.T) {
	m := NewMapInt64()
	m.Set(1, 123)
	m.Set(2, 456)
	m.Set(3, 789)
	fmt.Println(m.GetMap())
	fmt.Println(m.Get(1))
	for _, slice := range m.ToSlice() {
		fmt.Println(slice)
	}
	for _, slice := range m.KeyToSlice() {
		fmt.Println(slice)
	}
	for _, i := range m.ValueToSlice() {
		fmt.Println(i)
	}
}

func TestNewMapUint64(t *testing.T) {
	m := NewMapUint64()
	m.Set(1, 123)
	m.Set(2, 456)
	m.Set(3, 789)
	fmt.Println(m.GetMap())
	fmt.Println(m.Get(1))
	for _, slice := range m.ToSlice() {
		fmt.Println(slice)
	}
	for _, slice := range m.KeyToSlice() {
		fmt.Println(slice)
	}
	for _, i := range m.ValueToSlice() {
		fmt.Println(i)
	}
}

func TestNewMapInt32(t *testing.T) {
	m := NewMapInt32()
	m.Set(1, 123)
	m.Set(2, 456)
	m.Set(3, 789)
	fmt.Println(m.GetMap())
	fmt.Println(m.Get(1))
	for _, slice := range m.ToSlice() {
		fmt.Println(slice)
	}
	for _, slice := range m.KeyToSlice() {
		fmt.Println(slice)
	}
	for _, i := range m.ValueToSlice() {
		fmt.Println(i)
	}
}

func TestNewMapUint32(t *testing.T) {
	m := NewMapUint32()
	m.Set(1, 123)
	m.Set(2, 456)
	m.Set(3, 789)
	fmt.Println(m.GetMap())
	fmt.Println(m.Get(1))
	for _, slice := range m.ToSlice() {
		fmt.Println(slice)
	}
	for _, slice := range m.KeyToSlice() {
		fmt.Println(slice)
	}
	for _, i := range m.ValueToSlice() {
		fmt.Println(i)
	}
}

// bench测试仅供参考
// cpu: AMD Ryzen 5 5600H with Radeon Graphics
// BenchmarkNewMapString-12                 1274150               934.3 ns/op
func BenchmarkNewMapString(b *testing.B) {
	m := NewMapString()
	var w sync.WaitGroup
	for i := 0; i < b.N; i++ {
		w.Add(1)
		go func(_i int) {
			defer w.Done()
			m.Set(strconv.Itoa(_i), _i)
		}(i)
	}
	w.Wait()
}

// cpu: AMD Ryzen 5 5600H with Radeon Graphics
// BenchmarkNewMapStringAsyncRW-12          1000000              1782 ns/op
func BenchmarkNewMapStringAsyncRW(b *testing.B) {
	m := NewMapString()
	var w sync.WaitGroup
	for i := 0; i < b.N; i++ {
		w.Add(3)
		go func(_i int) {
			defer w.Done()
			go func(i int) {
				defer w.Done()
				m.Set(strconv.Itoa(i), i)
			}(_i)
			go func(i int) {
				defer w.Done()
				m.Get(strconv.Itoa(rand.Intn(i + 1)))
			}(_i)
		}(i)
	}
	w.Wait()
}

// cpu: AMD Ryzen 5 5600H with Radeon Graphics
// BenchmarkSyncMap-12      1000000              1582 ns/op
func BenchmarkSyncMap(b *testing.B) {
	var m sync.Map
	var w sync.WaitGroup
	for i := 0; i < b.N; i++ {
		w.Add(1)
		go func(_i int) {
			defer w.Done()
			m.Store(strconv.Itoa(_i), _i)
		}(i)
	}
	w.Wait()
}

// cpu: AMD Ryzen 5 5600H with Radeon Graphics
// BenchmarkSyncMapAsyncRW-12        908904              2968 ns/op
func BenchmarkSyncMapAsyncRW(b *testing.B) {
	var m sync.Map
	var w sync.WaitGroup
	for i := 0; i < b.N; i++ {
		w.Add(3)
		go func(_i int) {
			defer w.Done()
			go func(i int) {
				defer w.Done()
				m.Store(strconv.Itoa(i), i)
			}(_i)
			go func(i int) {
				defer w.Done()
				m.Load(strconv.Itoa(rand.Intn(i + 1)))
			}(_i)
		}(i)
	}
	w.Wait()
}
