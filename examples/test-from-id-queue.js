import { Http, Tracetest } from "k6/x/tracetest";
import { sleep } from "k6";

export const options = {
  vus: 1,
  duration: "5s",
};

const http = new Http();
const testId = "J0d887oVR";
const tracetest = Tracetest();

tracetest.updateFromConfig({
  server_url: "http://localhost:11633",
});

export default function () {
  const url = "http://localhost:8081/pokemon?take=5";
  const response = http.get(url);
  tracetest.runTest(
    response.trace_id,
    {
      test_id: testId,
      should_wait: true,
      variable_name: "TRACE_ID",
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
