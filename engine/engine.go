package engine

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type EventLoop struct {
	messageQueue []Command
	isFinished   chan (bool)
	Await        bool
}

func (eventloop *EventLoop) Start() {
	eventloop.isFinished = make(chan bool)
	go func() {
		for !eventloop.Await {
		}
		for i := 0; i < len(eventloop.messageQueue); i++ {
			cmd := eventloop.messageQueue[i]
			cmd.Execute(eventloop)
		}
		eventloop.isFinished <- true
	}()
}

func (eventloop *EventLoop) Post(cmd Command) {
	eventloop.messageQueue = append(eventloop.messageQueue, cmd)
}

func (eventloop *EventLoop) AwaitFinish() {
	eventloop.Await = true
	<-eventloop.isFinished
}
