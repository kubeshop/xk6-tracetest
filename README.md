# xk6-tracetest

This extension adds tracetest support to [k6](https://github.com/grafana/k6)!

That means that if you're testing an instrumented system, you can use this extension to trigger test runs.

Currently, it supports HTTP requests and the following propagation formats: `tracecontext`, `baggage`, `b3` `ot`, `jaeger` and `xray`.

It is implemented using the [xk6](https://github.com/grafana/xk6) extension system.

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download `xk6`:

```bash
$ go install go.k6.io/xk6/cmd/xk6@latest
```

2. Build the binary:

```bash
$ xk6 build --with github.com/kubeshop/xk6-tracetest@latest
```

## Available Variables

If you want to configure the tracetest k6 binary you can do it by using any of the following environment variables

- **XK6_TRACETEST_SERVER_URL:** Updates the tracetest server url for API interactions (can be overwritten by the script config)
- **XK6_TRACETEST_SERVER_PATH:** Updates the tracetest server path for API interactions (can be overwritten by the script config)
- **XK6_TRACETEST_SERVER_PATH:** Updates the tracetest server token that will be used to authenticate with the server (can be overwritten by the script config)

You can also set a default tracetest endpoint when running the k6 binary by using the following option:

`./k6 run examples/test-from-id.js -o xk6-tracetest=<server-url>`

## Example

To run a full example take a look at the fully flesh demo we have for you in the Tracetest main mono repo: [examples/tracetest-k6](https://github.com/kubeshop/tracetest/tree/main/examples/tracetest-k6)

## Running on https://app.tracetest.io

`./k6 run examples/test-from-id.js --env XK6_TRACETEST_API_TOKEN=<your token> -o xk6-tracetest=https://api.tracetest.io`
