package goroutinepool


/*
协程池恶汉和懒汉模式 https://mp.weixin.qq.com/s/obxJbpKMzzO0HyBEdP9opQ

 */


type StarvationPool struct {
	totalnum int
	tasks chan func()
}

func (p *StarvationPool) ListenWorkers() {
	for i:=0;i<p.totalnum;i++{
		go func() {
			for {
				select {
				case f := <-p.tasks:
					f()
				}
			}
		}()
	}
}

const Maxtasks 	int = 10

func NewStarvationPool(n int) *StarvationPool {
	pool:=&StarvationPool{
		totalnum: n,
		tasks:    make(chan func(),Maxtasks),
	}
	pool.ListenWorkers()
	return pool
}

func (p *StarvationPool) PushTask2Pool(f func())  {
	p.tasks<-f
}

