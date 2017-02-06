package threadpool

type TPWork interface {
	Execute()
}

type ThreadPool struct {
	workers  int
	workList chan TPWork
	term     bool
}

func CreateThreadPool(w int) ThreadPool {
	tp := ThreadPool{
		workers:  w,
		workList: make(chan TPWork, w),
		term:     false}

	for i := 0; i < tp.workers; i++ {
		go tp.workerThread()
	}
	return tp
}

func (tp *ThreadPool) SubmitWork(w TPWork) {
	tp.workList <- w
}

func (tp *ThreadPool) workerThread() {
	for !tp.term {
		work := <-tp.workList
		work.Execute()
	}
}
