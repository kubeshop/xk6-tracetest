package tracetest

import (
	"sync"

	"github.com/kubeshop/xk6-tracetest/models"
)

func (t *Tracetest) processQueue() {
	t.bufferLock.Lock()
	bufferedJobs := t.buffer
	t.buffer = make([]models.Job, 0, len(bufferedJobs)) // flushing queue
	t.bufferLock.Unlock()

	t.parallelJobProcessor(bufferedJobs)
}

func (t *Tracetest) queueJob(job models.Job) {
	t.bufferLock.Lock()
	defer t.bufferLock.Unlock()

	t.buffer = append(t.buffer, job)
}

func (t *Tracetest) parallelJobProcessor(jobs []models.Job) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(jobs))

	defer waitGroup.Wait()

	for _, job := range jobs {
		go func(job models.Job) {
			defer waitGroup.Done()

			run, err := t.runTest(&job)
			job = job.HandleRunResponse(run, err)

			t.processedBuffer.Store(job.ID, job)
		}(job)
	}
}
