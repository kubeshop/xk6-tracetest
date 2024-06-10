package httpClient

import (
	"context"
	"fmt"
	"time"

	"github.com/grafana/sobek"
	"github.com/kubeshop/xk6-tracetest/models"
	"github.com/kubeshop/xk6-tracetest/utils"
	k6HTTP "go.k6.io/k6/js/modules/k6/http"
	"go.k6.io/k6/lib"
	"go.k6.io/k6/metrics"
)

const (
	TraceID      = "trace_id"
	TestID       = "test_id"
	ShouldWait   = "should_wait"
	VariableName = "variable_name"
)

func (c *HttpClient) WithTrace(fn HttpFunc, url sobek.Value, args ...sobek.Value) (*HTTPResponse, error) {
	state := c.vu.State()
	if state == nil {
		return nil, fmt.Errorf("HTTP requests can only be made in the VU (virtual user) context")
	}

	traceID, _, err := (&models.TraceID{
		Prefix:            models.K6Prefix,
		Code:              models.K6Code_Cloud,
		UnixTimestampNano: uint64(time.Now().UnixNano()) / uint64(time.Millisecond),
	}).Encode()
	if err != nil {
		return nil, err
	}

	tracingHeaders := c.options.Propagator.GenerateHeaders(traceID)

	rt := c.vu.Runtime()
	var params *sobek.Object
	if len(args) < 2 {
		params = rt.NewObject()
		if len(args) == 0 {
			args = []sobek.Value{sobek.Null(), params}
		} else {
			args = append(args, params)
		}
	} else {
		jsParams := args[1]
		if utils.IsNilly(jsParams) {
			params = rt.NewObject()
			args[1] = params
		} else {
			params = jsParams.ToObject(rt)
		}
	}

	var headers *sobek.Object
	if jsHeaders := params.Get("headers"); utils.IsNilly(jsHeaders) {
		headers = rt.NewObject()
		params.Set("headers", headers)
	} else {
		headers = jsHeaders.ToObject(rt)
	}
	for key, val := range tracingHeaders {
		headers.Set(key, val)
	}

	c.setTags(rt, state, traceID, params)
	defer c.deleteTags(state)

	res, err := fn(c.vu.Context(), url, args...)
	return &HTTPResponse{Response: res, TraceID: traceID}, err
}

func (c *HttpClient) setTags(rt *sobek.Runtime, state *lib.State, traceID string, params *sobek.Object) {
	tracetestOptions := models.NewTracetestOptions(rt, params)
	state.Tags.Modify(func(tagsAndMeta *metrics.TagsAndMeta) {
		tagsAndMeta.SetMetadata(TraceID, traceID)

		if tracetestOptions.TestID != "" {
			tagsAndMeta.SetMetadata(TestID, tracetestOptions.TestID)
		} else if c.options.Tracetest.TestID != "" {
			tagsAndMeta.SetMetadata(TestID, c.options.Tracetest.TestID)
		}

		if tracetestOptions.ShouldWait || c.options.Tracetest.ShouldWait {
			tagsAndMeta.SetMetadata(ShouldWait, "true")
		}

		tagsAndMeta.SetMetadata(VariableName, tracetestOptions.VariableName)
	})
}

func (c *HttpClient) deleteTags(state *lib.State) {
	state.Tags.Modify(func(tagsAndMeta *metrics.TagsAndMeta) {
		tagsAndMeta.DeleteMetadata(TraceID)
		tagsAndMeta.DeleteMetadata(TestID)
		tagsAndMeta.DeleteMetadata(ShouldWait)
		tagsAndMeta.DeleteMetadata(VariableName)
	})
}

func requestToHttpFunc(method string, request HttpRequestFunc) HttpFunc {
	return func(ctx context.Context, url sobek.Value, args ...sobek.Value) (*k6HTTP.Response, error) {
		return request(method, url, args...)
	}
}
