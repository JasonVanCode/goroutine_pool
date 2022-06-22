package goroutine_pool

type Task struct {
	f func()
}

func NewTask(f func()) *Task {
	return &Task{
		f: f,
	}
}

func (t *Task) Excute() {
	t.f()
}
