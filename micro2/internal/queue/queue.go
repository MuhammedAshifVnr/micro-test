package queue

type TaskQueue struct {
	tasks chan func()
}

func NewTaskQueue(bufferSize int) *TaskQueue {
	q := &TaskQueue{
		tasks: make(chan func(), bufferSize),
	}
	go q.start()
	return q
}

func (q *TaskQueue) start() {
	for task := range q.tasks {
		task()
	}
}

// AddTask adds a task to the queue for sequential execution.
func (q *TaskQueue) AddTask(task func()) {
	q.tasks <- task
}

// Stop gracefully shuts down the queue.
func (q *TaskQueue) Stop() {
	close(q.tasks)
}
