package main

import (
	"fmt"
	"sync"
)

type New func() interface{}

type Pool struct {
	mutex   sync.Mutex
	objects []interface{}
	New     New
}

func NewPool(size int, new New) *Pool {
	objects := make([]interface{}, size)
	for i := 0; i < size; i++ {
		objects[i] = new()
	}
	return &Pool{
		objects: objects,
		New:     new,
	}
}

func (p *Pool) Get() interface{} {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if len(p.objects) > 0 {
		obj := p.objects[0]
		p.objects = p.objects[1:]
		return obj
	} else {
		return p.New()
	}

}

func (p *Pool) put(obj interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.objects = append(p.objects, obj)

}
func main() {

	pool := NewPool(1,func() interface{} {
		fmt.Println("new pool")
		return 1
	})

	x := pool.Get()
	fmt.Println(x)

	pool.put(2)

	x = pool.Get()
	fmt.Println(x)

	x = pool.Get()
	fmt.Println(x)



	fmt.Println(600 * 12 * 5 )											//36000
	fmt.Println(1800 * 12 * 5 - 38000 -5000 )								//6500

}
