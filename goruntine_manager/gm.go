package goruntine_manager

import (
	"errors"
)

type GoroutineManagerI interface {
	Start() error           //启动
	Exit()                  //终止goroutine manger 实例
	Reboot() error          //重启
	Run(func()) error       //运行一个Goroutine
	Dilatation(n int) error //扩容
	GoroutineNum() int      //获取Goroutine数量
}

type GoroutineManagerImpl struct {
	n       int //goroutine num
	handleC chan func()
}

func NewGoroutineManger(n int) GoroutineManagerI {
	g := new(GoroutineManagerImpl)
	g.handleC = make(chan func(), n)
	g.n = n
	return g
}

func (g *GoroutineManagerImpl) Start() error {
	if g.n <= 0 {
		return errors.New("goroutine num less 1 Refuse to start")
	}
	for i := 0; i < g.n; i++ {
		go func() {
			for f := range g.handleC {
				f()
			}
		}()
	}
	return nil
}

func (g *GoroutineManagerImpl) Exit() {
	close(g.handleC)
	g.handleC = nil
}

func (g *GoroutineManagerImpl) Run(f func()) error {
	if g.handleC == nil {
		return errors.New("GoroutineManager have already stop")
	}
	g.handleC <- f
	return nil
}

func (g *GoroutineManagerImpl) Reboot() error {
	g.Exit()
	return g.Start()
}

func (g *GoroutineManagerImpl) Dilatation(n int) error {
	if n <= 0 {
		return errors.New("goroutine num less 1 Refuse to start")
	}
	g.n += n
	for i := 0; i < n; i++ {
		go func() {
			for f := range g.handleC {
				f()
			}
		}()
	}
	return nil
}

func (g *GoroutineManagerImpl) GoroutineNum() int {
	return g.n
}
