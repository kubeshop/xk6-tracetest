package models

import (
	"fmt"

	"github.com/kubeshop/tracetest/cli/openapi"
)

type Test struct {
	ID   string
	Name string
}

type TracetestRun struct {
	TestId  string
	TestRun *openapi.TestRun
}

func (tr *TracetestRun) Summary(baseUrl string) string {
	runUrl := fmt.Sprintf("%s/test/%s/run/%d", baseUrl, tr.TestId, *tr.TestRun.Id)

	failingSpecs := true
	if tr.TestRun != nil && tr.TestRun.Result != nil && tr.TestRun.Result.AllPassed != nil {
		failingSpecs = !*tr.TestRun.Result.AllPassed
	}

	lastError := ""
	if tr.TestRun != nil && tr.TestRun.LastErrorState != nil {
		lastError = *tr.TestRun.LastErrorState
	}

	summary := fmt.Sprintf("RunState=%s FailingSpecs=%t, TracetestURL= %s", *tr.TestRun.State, failingSpecs, runUrl)

	if lastError != "" {
		summary += fmt.Sprintf(", LastError=%s", lastError)
	}

	return summary
}

func (tr *TracetestRun) IsSuccessful() bool {
	if tr.TestRun != nil && tr.TestRun.Result != nil && tr.TestRun.Result.AllPassed != nil {
		return *tr.TestRun.Result.AllPassed
	}

	return false
}
