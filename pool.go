package goroutine_pool

import (
	"fmt"
)

type Pool struct {
	Entorychannel chan *Task
	WorkerNum     int
	Jobchannel    chan *Task
}

func NewPool(cap int) *Pool {
	return &Pool{
		Entorychannel: make(chan *Task),
		WorkerNum:     cap,
		Jobchannel:    make(chan *Task),
	}
}

func (p *Pool) Worker(workerId int) {
	//从jobchannel中获取
	for v := range p.Jobchannel {
		v.Excute()
		fmt.Printf("workerId %d 执行了该任务", workerId)
	}
}

func (p *Pool) ReceiveTask(t *Task) {
	p.Entorychannel <- t
}

func (p *Pool) Run() {
	defer close(p.Entorychannel)
	defer close(p.Jobchannel)
	//开启工作池
	for i := 0; i < p.WorkerNum; i++ {
		go p.Worker(i)
	}
	//从入口chanel读取task，给jobchanel发送任务
	for v := range p.Entorychannel {
		p.Jobchannel <- v
	}
}
