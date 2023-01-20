package models

import (
	"fmt"

	"github.com/kubeshop/xk6-tracetest/openapi"
)

type JobType string

const (
	RunTestFromId JobType = "runTestFromId"
)

type JobStatus string

const (
	Pending JobStatus = "pending"
	Running JobStatus = "running"
	Failed  JobStatus = "failed"
	Success JobStatus = "success"
)

type Job struct {
	TraceID          string
	TestID           string
	VariableName     string
	JobType          JobType
	Request          Request
	Run              *TracetestRun
	JobStatus        JobStatus
	TracetestOptions TracetestOptions
	Error            string
}

func NewJob(traceId string, options TracetestOptions, request Request) Job {
	return Job{
		JobType:   RunTestFromId,
		Request:   request,
		JobStatus: Pending,

		TraceID:          traceId,
		TestID:           options.TestID,
		TracetestOptions: options,
	}
}

func (job Job) HandleRunResponse(run *openapi.TestRun, err error) Job {
	if run == nil {
		job.JobStatus = Failed
		job.Error = err.Error()
	} else {
		job.JobStatus = Success
		job.Run = &TracetestRun{
			TestRun: run,
			TestId:  job.TestID,
		}
	}

	return job
}

func (job Job) Summary(baseUrl string) string {
	runSummary := fmt.Sprintf("JobStatus=%s, Error=%s", string(job.JobStatus), job.Error)
	if job.Run != nil {
		runSummary = job.Run.Summary(baseUrl)
	}

	return fmt.Sprintf("Request=%s - %s, TraceID=%s, %s", job.Request.Method, job.Request.URL, job.TraceID, runSummary)
}

func (job Job) IsSuccessful() bool {
	isJobStatusSuccessful := job.JobStatus == Success
	runExists := job.Run != nil

	return isJobStatusSuccessful && runExists && job.Run.IsSuccessful()
}
