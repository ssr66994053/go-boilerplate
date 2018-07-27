package pool

// Job 执行接口
type Job interface {
	Payload()
}

// Pool 对象
type Pool struct {
	jobs    chan Job
	workers chan bool
}

// NewPool 生成Pool实例
func NewPool(maxJob, maxWorker int) *Pool {
	jobs := make(chan Job, maxJob)
	workers := make(chan bool, maxWorker)
	for i := 0; i < maxWorker; i++ {
		workers <- true
	}
	return &Pool{jobs, workers}
}

// Run 执行
func (p *Pool) Run() {
	for {
		// 取到任务
		job := <-p.jobs
		// 获取执行器
		<-p.workers
		// 执行
		go func(job Job) {
			defer func() {
				// 返回执行器
				p.workers <- true
			}()
			job.Payload()
		}(job)
	}
}

// Submit 提交任务
func (p *Pool) Submit(job Job) {
	p.jobs <- job
}

// Close 关闭
func (p *Pool) Close() {
	//FIXME 等待任务完成 关闭chan Run跳槽循环
}
