import { Http, Tracetest } from "k6/x/tracetest";
import { sleep } from "k6";

export const options = {
  vus: 1,
  duration: "5s",
};

const http = new Http({ propagator: ["b3"] });
const testId = "J0d887oVR";
const tracetest = Tracetest({
  serverUrl: "http://localhost:3000",
});

export default function () {
  const url = "http://localhost:8081/pokemon?take=5";
  const response = http.get(url);
  tracetest.runTest(testId, response.trace_id, true, {
    id: "123",
    url,
    method: "GET",
  });

  sleep(1);
}

export function handleSummary() {
  return {
    stdout: tracetest.summary(),
  };
}
