package thread_pool

import "log"

// 多协程执行多个任务
type ThreadPool struct {
	Size  int
	Total int
	Queue chan func() error

	result   chan error
	Callback func()
}

func NewThreadPool(threadSize, taskNum int) *ThreadPool {
	pool := new(ThreadPool)
	pool.Init(threadSize, taskNum)
	return pool
}

// 初始化
func (self *ThreadPool) Init(threadSize, taskNum int) {
	self.Size = threadSize
	self.Total = taskNum
	self.Queue = make(chan func() error, taskNum)
	self.result = make(chan error, taskNum)
}

func (self *ThreadPool) Start() {
	// 开启size个协程执行任务
	for i := 0; i < self.Size; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}

				err := task()
				self.result <- err
			}
		}()
	}

	// 等待执行结果
	for i := 0; i < self.Total; i++ {
		res, ok := <-self.result
		if !ok {
			break
		}
		if res != nil {
			log.Println(res)
		}
	}

	// 执行回调
	if self.Callback != nil {
		self.Callback()
	}
}

func (self *ThreadPool) Stop() {
	close(self.Queue)
	close(self.result)
}

func (self *ThreadPool) AddTask(task func() error) {
	self.Queue <- task
}

func (self *ThreadPool) SetCallBack(callBack func()) {
	self.Callback = callBack
}
