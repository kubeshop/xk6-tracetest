package tracetest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/kubeshop/xk6-tracetest/models"
	"github.com/kubeshop/xk6-tracetest/openapi"
)

func NewAPIClient(options models.ApiOptions) *openapi.APIClient {
	url, err := url.Parse(options.ServerUrl)

	if err != nil {
		panic(err)
	}

	config := openapi.NewConfiguration()
	config.Host = url.Host
	config.Scheme = url.Scheme

	if options.APIToken != "" {
		jwt, err := getJWTFromToken(url, options.APIToken)
		if err != nil {
			panic(err)
		}
		config.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", jwt))
	}

	if options.ServerPath != "" {
		config.Servers = []openapi.ServerConfiguration{
			{
				URL: options.ServerPath,
			},
		}
	}

	return openapi.NewAPIClient(config)
}

func getJWTFromToken(url *url.URL, token string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s://%s/tokens/%s/exchange", url.Scheme, url.Host, token), nil)
	if err != nil {
		return "", fmt.Errorf("could not create request for token exchange: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not send request for token exchange: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		// Probably an OS version of tracetest
		return "", fmt.Errorf("tracetest server doesn't support API Tokens")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read body from response: %w", err)
	}

	respJson := struct {
		JWT string `json:"jwt"`
	}{}

	err = json.Unmarshal(respBody, &respJson)
	if err != nil {
		return "", fmt.Errorf("could not unmarshal response body: %w", err)
	}

	return respJson.JWT, nil
}

func (t *Tracetest) runTest(job models.Job) (*openapi.TestRun, error) {
	request := t.client.ApiApi.RunTest(context.Background(), job.TestID)
	request = request.RunInformation(openapi.RunInformation{
		Variables: []openapi.EnvironmentValue{{
			Key:   &job.TracetestOptions.VariableName,
			Value: &job.TraceID,
		}},
		Metadata: job.Request.Metadata,
	})

	run, _, err := t.client.ApiApi.RunTestExecute(request)
	return run, err
}

func (t *Tracetest) waitForTestResult(testID, testRunID string) (openapi.TestRun, error) {
	var (
		testRun   openapi.TestRun
		lastError error
		wg        sync.WaitGroup
	)

	wg.Add(1)
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				readyTestRun, err := t.getIsTestReady(context.Background(), testID, testRunID)
				if err != nil {
					lastError = err
					wg.Done()
					return
				}

				if readyTestRun != nil {
					testRun = *readyTestRun
					wg.Done()
					return
				}
			}
		}
	}()
	wg.Wait()

	if lastError != nil {
		return openapi.TestRun{}, lastError
	}

	return testRun, nil
}

func (t *Tracetest) getIsTestReady(ctx context.Context, testID, testRunId string) (*openapi.TestRun, error) {
	req := t.client.ApiApi.GetTestRun(ctx, testID, testRunId)
	run, _, err := t.client.ApiApi.GetTestRunExecute(req)

	if err != nil {
		return &openapi.TestRun{}, fmt.Errorf("could not execute GetTestRun request: %w", err)
	}

	if *run.State == "FAILED" || *run.State == "FINISHED" {
		return run, nil
	}

	return nil, nil
}

func (t *Tracetest) jobSummary() (successfulJobs, failedJobs []models.Job) {
	t.processedBuffer.Range(func(_, value interface{}) bool {
		if job, ok := value.(models.Job); ok {
			if job.IsSuccessful() {
				successfulJobs = append(successfulJobs, job)
			} else {
				failedJobs = append(failedJobs, job)
			}
		}

		return true
	})

	return
}

func (t *Tracetest) stringSummary() string {
	successfulJobs, failedJobs := t.jobSummary()
	failedSummary := "[FAILED] \n"
	successfulSummary := "[SUCCESSFUL] \n"
	totalRuns := len(successfulJobs) + len(failedJobs)
	failedRuns := len(failedJobs)
	successfulRuns := len(successfulJobs)

	for _, job := range failedJobs {
		failedSummary += fmt.Sprintf("[%s] \n", job.Summary(t.apiOptions.ServerUrl))
	}

	for _, job := range successfulJobs {
		successfulSummary += fmt.Sprintf("[%s] \n", job.Summary(t.apiOptions.ServerUrl))
	}

	totalResults := fmt.Sprintf("[TotalRuns=%d, SuccessfulRus=%d, FailedRuns=%d] \n", totalRuns, successfulRuns, failedRuns)

	if failedRuns == 0 {
		failedSummary = ""
	}

	if successfulRuns == 0 {
		successfulSummary = ""
	}

	return totalResults + failedSummary + successfulSummary
}

type JsonResult struct {
	TotalRuns      int
	SuccessfulRuns int
	FailedRuns     int
	Failed         []models.Job
	Successful     []models.Job
}

func (t *Tracetest) jsonSummary() JsonResult {
	JsonResult := JsonResult{
		TotalRuns:      0,
		SuccessfulRuns: 0,
		FailedRuns:     0,
		Failed:         []models.Job{},
		Successful:     []models.Job{},
	}

	t.processedBuffer.Range(func(_, value interface{}) bool {
		if job, ok := value.(models.Job); ok {
			JsonResult.TotalRuns += 1
			if job.IsSuccessful() {
				JsonResult.Successful = append(JsonResult.Successful, job)
				JsonResult.SuccessfulRuns += 1
			} else {
				JsonResult.Failed = append(JsonResult.Failed, job)
				JsonResult.FailedRuns += 1
			}
		}

		return true
	})

	return JsonResult
}
