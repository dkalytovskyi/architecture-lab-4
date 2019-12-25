package engine

import "sync"

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

var isFinished = make(chan bool)

type EventLoop struct {
	sync.Mutex
	messageQueue []Command
	Await        bool
}

func (eventloop *EventLoop) Start() {
	go func() {
		for !eventloop.Await {
		}
		for i := 0; i < len(eventloop.messageQueue); i++ {
			eventloop.Lock()
			cmd := eventloop.messageQueue[i]
			eventloop.Unlock()
			cmd.Execute(eventloop)
		}
		isFinished <- true
	}()
}

func (eventloop *EventLoop) Post(cmd Command) {
	eventloop.Lock()
	eventloop.messageQueue = append(eventloop.messageQueue, cmd)
	eventloop.Unlock()

}

func (eventloop *EventLoop) AwaitFinish() {
	eventloop.Await = true
	<- isFinished
}
