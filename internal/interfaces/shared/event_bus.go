package shared

import (
	"fmt"
	"reflect"
	"sync"
)

// EventBus 结构体
type EventBus struct {
	handlers map[reflect.Type][]reflect.Value // 事件类型 -> 处理器列表
	mu       sync.RWMutex                    // 并发安全
}

// NewEventBus 创建一个新的 EventBus
func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[reflect.Type][]reflect.Value),
	}
}

// Register 注册事件处理器（根据函数参数类型推断事件类型）
func (bus *EventBus) Register(handler any)  {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	// 检查 handler 是否是函数
	handlerValue := reflect.ValueOf(handler)
	if handlerValue.Kind() != reflect.Func {
		panic("handler must be a function")
	}

	// 检查函数的参数是否只有一个
	handlerType := handlerValue.Type()
	if handlerType.NumIn() != 1 {
		panic("handler must have exactly one parameter")
	}

	// 获取事件类型（函数的第一个参数类型）
	eventType := handlerType.In(0)

	// 添加 handler
	bus.handlers[eventType] = append(bus.handlers[eventType], handlerValue)
}

// Dispatch 分发事件
func (bus *EventBus) Dispatch(event any) error {
	bus.mu.RLock()
	defer bus.mu.RUnlock()

	// 获取事件类型
	eventType := reflect.TypeOf(event)

	// 找到所有处理器
	handlers, ok := bus.handlers[eventType]

	if !ok {
		return fmt.Errorf("no handler for event %v", eventType)
	}

	// 调用所有处理器
	for _, handler := range handlers {
		if err := handler.Call([]reflect.Value{reflect.ValueOf(event)})[0].Interface(); err != nil {
			return err.(error)
		}
	}

	println("event dispatched")

	return nil
}
