package observer

import (
	"fmt"
	"reflect"
	"sync"
)

type IEventBus interface {
	// On(eventName string, handler interface{})
	// Off(evnetName string, handler interface{})
	// 订阅事件
	Subscribe(eventName string, handler interface{}) error
	// 取消订阅事件
	UnSubscribe(eventName string, handler interface{}) error
	// 发布事件
	Publish(eventName string, args ...interface{})
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: make(map[string][]reflect.Value),
		lock:     sync.Mutex{},
	}
}

func (b *AsyncEventBus) Subscribe(eventName string, handler interface{}) error {

	b.lock.Lock()
	defer b.lock.Unlock()

	v := reflect.ValueOf(handler)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler must be Func but is %s", v.Type().Kind())
	}

	// 往事件注册函数, 不存在初始化数组
	handlers, ok := b.handlers[eventName]
	if !ok {
		handlers = []reflect.Value{}
	}
	handlers = append(handlers, v)
	b.handlers[eventName] = handlers
	return nil
}

func (b *AsyncEventBus) UnSubscribe(eventName string, handler interface{}) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	handlers, ok := b.handlers[eventName]
	if !ok {
		return fmt.Errorf("the eventName dosent's exist")
	}
	handlerFunc := reflect.ValueOf(handler)
	for i, v := range handlers {
		if v == handlerFunc {
			handlers = append(handlers[:i], handlers[i+1:]...)
			b.handlers[eventName] = handlers
		}
	}
	return nil
}

// 发布事件
func (b *AsyncEventBus) Publish(eventName string, args ...interface{}) {
	handlers, ok := b.handlers[eventName]
	if !ok {
		return
	}

	params := make([]reflect.Value, len(args))
	for i, v := range args {
		params[i] = reflect.ValueOf(v)
	}

	for _, v := range handlers {
		go v.Call(params)
	}

	// for i := range handlers {
	// 	go handlers[i].Call(params)
	// }
}
