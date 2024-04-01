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

	"github.com/golang-jwt/jwt"
	"github.com/kubeshop/tracetest/cli/openapi"
	"github.com/kubeshop/xk6-tracetest/models"
)

var (
	traceID      = "TRACE_ID"
	resourceType = "Test"
	testName     = "K6 Test"
	testID       = "k6-test"
	testTrigger  = "k6"
	defaultTest  = openapi.TestResource{
		Type: &resourceType,
		Spec: &openapi.Test{
			Id:   &testID,
			Name: &testName,
			Trigger: &openapi.Trigger{
				Type: &testTrigger,
			},
		},
	}
)

func NewAPIClient(options models.ApiOptions) (*openapi.APIClient, string) {
	url, err := url.Parse(options.ServerUrl)

	if err != nil {
		panic(err)
	}

	config := openapi.NewConfiguration()
	config.Host = url.Host
	config.Scheme = url.Scheme

	jwt := ""
	if options.APIToken != "" {
		version, err := getVersion(url)
		if err != nil {
			panic(err)
		}

		url, err = url.Parse(*version.ApiEndpoint)
		if err != nil {
			panic(err)
		}

		jwt, err = getJWTFromToken(url, options.APIToken)
		if err != nil {
			panic(err)
		}

		config.Host = url.Host
		config.Scheme = url.Scheme
		options.ServerPath = url.Path
		config.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", jwt))
	}

	if options.ServerPath != "" {
		config.Servers = []openapi.ServerConfiguration{
			{
				URL: options.ServerPath,
			},
		}
	}

	return openapi.NewAPIClient(config), jwt
}

func getVersion(url *url.URL) (*openapi.Version, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s://%s/version.json", url.Scheme, url.Host), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request for version: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request for version: %w", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body from response: %w", err)
	}
	defer resp.Body.Close()

	var version openapi.Version

	err = json.Unmarshal(respBody, &version)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response body: %w", err)
	}

	return &version, nil
}

func getJWTFromToken(url *url.URL, token string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s://%s%s/tokens/%s/exchange", url.Scheme, url.Host, url.Path, token), nil)
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
	defer resp.Body.Close()
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

func getTokenClaims(tokenString string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (t *Tracetest) upsertTest(ctx context.Context) (*openapi.TestResource, error) {
	req := t.client.ResourceApiApi.UpsertTest(ctx)
	req = req.TestResource(defaultTest)

	test, _, err := t.client.ResourceApiApi.UpsertTestExecute(req)
	return test, err
}

func (t *Tracetest) runTest(job *models.Job) (*openapi.TestRun, error) {
	if job.TestID == "" {
		_, err := t.upsertTest(context.Background())
		if err != nil {
			return nil, fmt.Errorf("could not create test: %w", err)
		}

		job.TestID = testID
	}

	request := t.client.ApiApi.RunTest(context.Background(), job.TestID)
	request = request.RunInformation(openapi.RunInformation{
		RunGroupId: &job.RunGroupId,
		Variables: []openapi.VariableSetValue{{
			Key:   &traceID,
			Value: &job.TraceID,
		}},
		Metadata: job.Request.Metadata,
	})

	run, resp, err := t.client.ApiApi.RunTestExecute(request)
	respBody, readErr := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if readErr != nil {
		return nil, fmt.Errorf("could not read body from response: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("could not execute test run: %s %w,", respBody, err)
	}

	return run, nil
}

func (t *Tracetest) waitForRunGroup(runGroupId string) (openapi.RunGroup, error) {
	var (
		runGroup  openapi.RunGroup
		lastError error
		wg        sync.WaitGroup
	)

	wg.Add(1)
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				readyGroup, err := t.getIsRunGroupReady(context.Background(), runGroupId)
				if err != nil {
					lastError = err
					wg.Done()
					return
				}

				if readyGroup != nil {
					runGroup = *readyGroup
					wg.Done()
					return
				}
			}
		}
	}()
	wg.Wait()

	if lastError != nil {
		return openapi.RunGroup{}, lastError
	}

	return runGroup, nil
}

func (t *Tracetest) getIsRunGroupReady(ctx context.Context, runGroupId string) (*openapi.RunGroup, error) {
	req := t.client.ApiApi.GetRunGroup(ctx, runGroupId)
	runGroup, _, err := t.client.ApiApi.GetRunGroupExecute(req)

	if err != nil {
		return &openapi.RunGroup{}, fmt.Errorf("could not execute GetTestRun request: %w", err)
	}

	if isRunGroupDone(*runGroup.Status) {
		return runGroup, nil
	}

	return nil, nil
}

func (t *Tracetest) jobSummary() (successfulJobs, failedJobs []models.Job) {
	t.processedBuffer.Range(func(_, value interface{}) bool {
		if job, ok := value.(models.Job); ok {
			req := t.client.ApiApi.GetTestRun(context.Background(), job.TestID, job.Run.TestRun.GetId())
			run, _, err := t.client.ApiApi.GetTestRunExecute(req)
			if err != nil {
				t.logger.Errorf("could not get test run: %s", err)
				return false
			}

			job.Run.TestRun = run

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

func (t *Tracetest) getBaseUrl() string {
	base := t.apiOptions.ServerUrl

	if t.jwt != "" {
		claims, _ := getTokenClaims(t.jwt)
		organizationId := claims["organization_id"].(string)
		environmentId := claims["environment_id"].(string)

		return fmt.Sprintf("%s/organizations/%s/environments/%s", base, organizationId, environmentId)
	}

	return base
}

func (t *Tracetest) stringSummary() string {
	successfulJobs, failedJobs := t.jobSummary()
	failedSummary := "[FAILED] \n"
	successfulSummary := "[SUCCESSFUL] \n"
	totalRuns := len(successfulJobs) + len(failedJobs)
	failedRuns := len(failedJobs)
	successfulRuns := len(successfulJobs)

	baseUrl := t.getBaseUrl()

	for _, job := range failedJobs {
		failedSummary += fmt.Sprintf("[%s] \n", job.Summary(baseUrl))
	}

	for _, job := range successfulJobs {
		successfulSummary += fmt.Sprintf("[%s] \n", job.Summary(baseUrl))
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
			req := t.client.ApiApi.GetTestRun(context.Background(), job.TestID, job.Run.TestRun.GetId())
			run, _, err := t.client.ApiApi.GetTestRunExecute(req)
			if err != nil {
				t.logger.Errorf("could not get test run: %s", err)
				return false
			}

			job.Run.TestRun = run

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

func isRunGroupDone(state string) bool {
	return state == "failed" || state == "succeed"
}
