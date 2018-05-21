package event

type EventBlock struct {
	ID           uint64
	EventContent Event
}

type Queue []EventBlock

// Init Queue
func NewQueue() *Queue {
	var queue *Queue = new(Queue)
	return queue
}

// Queuing data to queue and return Queue
func Enqueue(eventBlock EventBlock, queue *Queue) {
	*queue = append(*queue, eventBlock)
}

// Dequeue data
func Dequeue(queue *Queue) EventBlock {
	var head EventBlock = (*queue)[0]
	*queue = (*queue)[1:]
	return head
}
