package observer

import (
	"fmt"
	"testing"
	"time"
)

func TestEventBus(t *testing.T) {
	eb := NewAsyncEventBus()
	eb.Subscribe("read", func(msg string) {
		fmt.Println("read1 !", msg)
	})
	eb.Subscribe("read", func(msg string) {
		fmt.Println("read2 !", msg)
	})
	fmt.Println(eb.handlers)
	eb.Publish("read", "hi")

	time.Sleep(1 * time.Second)

}

func handler1(msg string) {
	fmt.Println("handler1", msg)
}

func TestEventBusUnscribe(t *testing.T) {

	eb := NewAsyncEventBus()
	eb.Subscribe("read", func(msg string) {
		fmt.Println("read1 !", msg)
	})
	eb.Subscribe("read", handler1)

	fmt.Println(eb.handlers)
	eb.Publish("read", "hi")

	time.Sleep(3 * time.Microsecond)
	t.Log("remove handler1")
	eb.UnSubscribe("read", handler1)
	eb.Publish("read", "h2")
	time.Sleep(1 * time.Second)
}
