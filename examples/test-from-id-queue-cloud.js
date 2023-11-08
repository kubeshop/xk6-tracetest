import { Http, Tracetest } from "k6/x/tracetest";
import { sleep } from "k6";

export const options = {
  vus: 1,
  duration: "5s",
};

const http = new Http();
const testId = "c80sJ_4SR";
const tracetest = Tracetest();
tracetest.updateFromConfig({
  api_token: "your-api-token",
});

export default function () {
  const url = "http://localhost:8081/pokemon?take=5";
  const response = http.get(url);

  tracetest.runTest(
    response.trace_id,
    {
      test_id: testId,
      variable_name: "TRACE_ID",
      should_wait: true,
    },
    {
      id: "123",
      url,
      method: "GET",
    }
  );

  sleep(1);
}

export function handleSummary() {
  return {
    stdout: tracetest.summary(),
  };
}
