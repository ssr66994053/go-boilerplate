package pool

import "sync"

// Job 执行接口
type Job interface {
	Payload()
}

// Pool 对象
type Pool struct {
	jobs     chan Job
	workers  chan bool
	stopSign chan bool
	wait     *sync.WaitGroup
}

// NewPool 生成Pool实例
func NewPool(maxJob, maxWorker int) *Pool {
	jobs := make(chan Job, maxJob)
	workers := make(chan bool, maxWorker)
	for i := 0; i < maxWorker; i++ {
		workers <- true
	}
	return &Pool{jobs, workers, make(chan bool, 1), &sync.WaitGroup{}}
}

// Run 执行
func (p *Pool) Run() {
	for {
		select {
		case job := <-p.jobs: // 取到任务
			p.wait.Add(1)
			// 获取执行器
			<-p.workers
			// 执行
			go func(job Job) {
				defer func() {
					// 返回执行器
					p.workers <- true
					p.wait.Done()
				}()
				job.Payload()
			}(job)
		case <-p.stopSign: // 关闭
			break
		}
	}
}

// Submit 提交任务
func (p *Pool) Submit(job Job) {
	p.jobs <- job
}

// Close 关闭
func (p *Pool) Close() {
	defer func() {
		// 关闭chan
		close(p.stopSign)
		close(p.workers)
		close(p.jobs)
	}()
	// 发出关闭信号 Run跳槽循环
	p.stopSign <- true
	//等待任务完成
	p.wait.Wait()
}
